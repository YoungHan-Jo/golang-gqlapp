package handler

import (
	"fmt"

	gql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
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

// (GET /companies)
func (h *Handler) GetCompanies(ctx echo.Context) error {
	panic("unimplemented")
}

// (POST /companies)
func (h *Handler) AddCompany(ctx echo.Context) error {
	panic("unimplemented")
}

func (h *Handler) PostGraphql(c echo.Context) error {
	h.gql.ServeHTTP(c.Response(), c.Request())
	return nil
}

func (h *Handler) GetPlayground(c echo.Context) error {
	fmt.Println("playground")
	playground.Handler("GraphQL playground", "/graphql").ServeHTTP(c.Response(), c.Request())
	return nil
}
