package main

import (
	"events_example/domain/event"
	"events_example/graph"
	"events_example/graph/generated"
	"events_example/graph/model"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/google/uuid"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	uuidForFirst := uuid.New()
	events := map[uuid.UUID]model.Event{
		uuidForFirst: {ID: uuidForFirst, Name: "firstEvent"},
	}

	tr := event.NewTestRepository(&events)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		TR: tr,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
