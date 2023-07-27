package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/danny-yamamoto/go-graphql-federation-example/users/graph"
	"github.com/danny-yamamoto/go-graphql-federation-example/users/internal"
)

const defaultPort = "4000"

func main() {
	port := os.Getenv("USERS_PORT")
	if port == "" {
		port = defaultPort
	}

	es := internal.NewExecutableSchema(internal.Config{Resolvers: &graph.Resolver{}})
	srv := handler.NewDefaultServer(internal.NewExecutableSchema(internal.Config{Resolvers: &graph.Resolver{Schema: es.Schema()}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
