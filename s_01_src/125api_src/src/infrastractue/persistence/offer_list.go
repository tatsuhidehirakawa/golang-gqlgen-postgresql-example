package persistence

import (
	"context"
	"database/sql"

	sqlc "github.com/gqlgensamples/golang-gqlgen-postgresql-example/db"
	"github.com/gqlgensamples/golang-gqlgen-postgresql-example/domain/repository"
)

type OfferListRepository struct {
	sqlc *sqlc.Queries
}

var _ repository.IOfferListRepository = (*OfferListRepository)(nil)

func NewOfferListRepository(sqlc *sqlc.Queries) *OfferListRepository {
	return &OfferListRepository{
		sqlc: sqlc,
	}
}

func (r *OfferListRepository) OfferLists(ctx context.Context) ([]sqlc.OfferList, error) {
	ol, err := r.sqlc.OfferLists(ctx)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, nil
		default:
			return nil, err
		}
	}
	return ol, nil
}
