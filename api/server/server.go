package server

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/younghan-jo/gqlgen-todos/api/env"
	"github.com/younghan-jo/gqlgen-todos/api/handler"
)

const (
	serverTimeout time.Duration = 5 * time.Second
)

type server struct {
	srv *http.Server
	db  *sql.DB
}

func NewServer() *server {

	debug := flag.Bool("debug", true, "is Debug Mode")
	flag.Parse()

	// Swagger
	swagger, err := handler.GetSwagger()
	if err != nil {
		panic(err)
	}
	swagger.Servers = nil

	// Log
	newLog(*debug)

	// DB
	db, err := newDB(fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		env.Get().DB.Host,
		env.Get().DB.Port,
		env.Get().DB.User,
		env.Get().DB.Password,
		env.Get().DB.Database,
	))
	if err != nil {
		panic(err)
	}

	// init boiler
	boil.SetDB(db)
	boil.DebugMode = *debug
	boil.DebugWriter = log.Logger

	// echo server
	e := echo.New()
	e.Use(
		// CORS Policy
		echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
			AllowOrigins: []string{
				env.Get().AllowOrigin,
			},
			AllowMethods: []string{
				http.MethodGet,
				http.MethodHead,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
				http.MethodConnect,
				http.MethodOptions,
				http.MethodTrace,
			},
			AllowHeaders: []string{"*"},
		}),
		// Recovery
		echomiddleware.Recover(),
		// Request Logger
		echomiddleware.RequestLoggerWithConfig(echomiddleware.RequestLoggerConfig{
			LogURI:    true,
			LogStatus: true,
			LogMethod: true,
			LogValuesFunc: func(c echo.Context, v echomiddleware.RequestLoggerValues) error {
				log.Info().
					Str("method", v.Method).
					Str("URI", v.URI).
					Int("status", v.Status).
					Msg("request")
				return nil
			},
		}),
		// OpenAPI Validator for swagger
		middleware.OapiRequestValidator(swagger),
	)

	if h, err := handler.NewHandler(); err != nil {
		panic(err)
	} else {
		handler.RegisterHandlers(e, h)
	}

	return &server{
		srv: &http.Server{
			Addr:    fmt.Sprintf(":%d", env.Get().Server.Port),
			Handler: e.Server.Handler,
		},
		db: db,
	}
}

func (s *server) Run() error {
	log.Info().Msgf("server listening on %s\n", s.srv.Addr)

	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *server) Shutdown() {
	log.Info().Msg("Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), serverTimeout)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		panic(err)
	}

	if err := s.db.Close(); err != nil {
		panic(err)
	}

	log.Info().Msg("Server terminated")
}
