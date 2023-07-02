package persistence

import (
	sqlc "github.com/gqlgensamples/golang-gqlgen-postgresql-example/db"
	"github.com/gqlgensamples/golang-gqlgen-postgresql-example/domain/repository"
	"github.com/gqlgensamples/golang-gqlgen-postgresql-example/io"
)

type Repositories struct {
	db           *io.SQLdatabase
	sqlc         *sqlc.Queries
	User         repository.IUserRepository
	AccountMater repository.IMasterAccountRepository
}

func NewRepositories(db *io.SQLdatabase) (*Repositories, error) {
	return &Repositories{
		db:           db,
		sqlc:         db.SQLC,
		User:         NewUserRepository(db),
		AccountMater: NewAccountMaterRepository(db.SQLC),
	}, nil
}
