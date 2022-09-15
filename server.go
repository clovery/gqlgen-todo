package main

import (
	"gqlgen-todo/graph"
	"gqlgen-todo/graph/generated"
	"log"
	"net/http"
	"os"

	model "gqlgen-todo/internal/model"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "3102"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	model.Init()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
