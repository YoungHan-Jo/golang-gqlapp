package server

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/younghan-jo/gqlgen-todos/api/graph"
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

	debug := flag.Bool("debug", false, "is Debug Mode")
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
	newDB(fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	))

}

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
