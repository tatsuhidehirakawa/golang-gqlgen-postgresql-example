package main

//go:generate go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
//go:generate sqlc generate

import (
	"context"
	"log"
	"net/http"

	"github.com/gqlgensamples/golang-gqlgen-postgresql-example/config"
	"github.com/gqlgensamples/golang-gqlgen-postgresql-example/infrastractue/persistence"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gqlgensamples/golang-gqlgen-postgresql-example/graph"
	"github.com/gqlgensamples/golang-gqlgen-postgresql-example/io"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db, err := io.NewSQLdatabase(cfg)
	if err != nil {
		log.Fatal(err)
	}

	repository, err := persistence.NewRepositories(db)
	if err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{
					Repo: repository,
				},
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/queries"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
