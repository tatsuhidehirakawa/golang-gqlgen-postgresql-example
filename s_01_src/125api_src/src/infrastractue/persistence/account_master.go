package persistence

import (
	"context"
	"database/sql"

	sqlc "github.com/gqlgensamples/golang-gqlgen-postgresql-example/db"
	"github.com/gqlgensamples/golang-gqlgen-postgresql-example/domain/repository"
)

type AccountMasterRepository struct {
	sqlc *sqlc.Queries
}

var _ repository.IMasterAccountRepository = (*AccountMasterRepository)(nil)

func NewAccountMaterRepository(sqlc *sqlc.Queries) *AccountMasterRepository {
	return &AccountMasterRepository{
		sqlc: sqlc,
	}
}

func (r *AccountMasterRepository) GetAccountMaster(ctx context.Context, userID string) (*sqlc.AccountMaster, error) {
	accountMaster, err := r.sqlc.AccountMaster(ctx, userID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, nil
		default:
			return nil, err
		}
	}
	return &accountMaster, nil
}
