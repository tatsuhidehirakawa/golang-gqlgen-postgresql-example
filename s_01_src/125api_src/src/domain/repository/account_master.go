package repository

import (
	"context"

	sqlc "github.com/gqlgensamples/golang-gqlgen-postgresql-example/db"
)

type IMasterAccountRepository interface {
	AccountMaster(ctx context.Context, ID string) (*sqlc.AccountMaster, error)
}
