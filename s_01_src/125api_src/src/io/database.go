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

type SQLdatabase struct {
	SQLX *sqlx.DB
	SQLC *sqlc.Queries
}

func NewSQLdatabase(cfg *config.Config) (*SQLdatabase, error) {
	dsn := fmt.Sprintf(
		"user=%s dbname=%s password=%s  host=%s port=5432 sslmode=disable",
		cfg.PostgresUser, cfg.PostgresDB, cfg.PostgresPassword, cfg.PostgresHost)
	sqlxDB, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	err = sqlxDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	//sqlc
	sqlcDB, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	err = sqlcDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	cdb := sqlc.New(sqlcDB)

	return &SQLdatabase{SQLX: sqlxDB, SQLC: cdb}, nil
}
