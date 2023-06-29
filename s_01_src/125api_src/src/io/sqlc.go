// データベースとの接続
package io

import (
	"fmt"
	"log"

	"github.com/gqlgensamples/golang-gqlgen-postgresql-example/config"
	sqlc "github.com/gqlgensamples/golang-gqlgen-postgresql-example/db"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	Sqlc *sqlc.Queries
}

func NewDB(cfg *config.Config) (*DB, error) {
	dsn := fmt.Sprintf(
		"user=%s dbname=%s password=%s  host=%s port=5432 sslmode=disable",
		cfg.PostgresUser, cfg.PostgresDB, cfg.PostgresPassword, cfg.PostgresHost)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	d := sqlc.New(db)

	return &DB{Sqlc: d}, nil
}
