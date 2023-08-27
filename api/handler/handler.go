package handler

import (
	"net/http"

	gql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo/v4"
	"github.com/younghan-jo/gqlgen-todos/api/graph"
	gql_middleware "github.com/younghan-jo/gqlgen-todos/api/graph/middleware"
)

type Handler struct {
	gql *gql_handler.Server
}

var _ ServerInterface = (*Handler)(nil)

func NewHandler() (*Handler, error) {

	// graphql handler
	gqlHandler := gql_handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{},
			},
		),
	)
	gqlHandler.SetRecoverFunc(gql_middleware.GraphQLRecovery)

	return &Handler{
		gql: gqlHandler,
	}, nil
}

// AddComment implements ServerInterface.
func (*Handler) AddComment(ctx echo.Context) error {
	panic("unimplemented")
}

// GetComments implements ServerInterface.
func (*Handler) GetComments(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "OK")
	// panic("unimplemented")
}

func (h *Handler) PostGraphql(c echo.Context) error {
	h.gql.ServeHTTP(c.Response(), c.Request())
	return nil
}
