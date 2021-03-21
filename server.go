package main

import (
	"log"
	"net/http"
	"os"
	"github.com/adesokanayo/hackernews/graph"
	 "github.com/adesokanayo/hackernews/graph/generated"
    "github.com/adesokanayo/hackernews/internal/pkg/db/migrations/mysql"
	"github.com/99designs/gqlgen/handler"
	"github.com/adesokanayo/hackernews/internal/auth"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	database.InitDB()
	database.Migrate()
	server := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
