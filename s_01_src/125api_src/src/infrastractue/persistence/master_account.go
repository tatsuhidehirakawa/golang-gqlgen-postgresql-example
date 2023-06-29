package persistence

import (
	"context"
	"database/sql"

	sqlc "github.com/gqlgensamples/golang-gqlgen-postgresql-example/db"
	"github.com/gqlgensamples/golang-gqlgen-postgresql-example/domain/repository"
)

type MasterAccountRepository struct {
	sqlc *sqlc.Queries
}

var _ repository.IMasterAccountRepository = (*MasterAccountRepository)(nil)

func NewMasterAccountRepository(sqlc *sqlc.Queries) *MasterAccountRepository {
	return &MasterAccountRepository{
		sqlc: sqlc,
	}
}

func (r *MasterAccountRepository) GetMasterAccount(ctx context.Context, userID string) (*sqlc.AccountMaster, error) {
	accountMaster, err := r.sqlc.GetAccountMaster(ctx, userID)
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
