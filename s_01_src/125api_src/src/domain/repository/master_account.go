package repository

import (
	"context"

	sqlc "github.com/gqlgensamples/golang-gqlgen-postgresql-example/db"
)

type IMasterAccountRepository interface {
	GetMasterAccount(ctx context.Context, ID string) (*sqlc.AccountMaster, error)
}
