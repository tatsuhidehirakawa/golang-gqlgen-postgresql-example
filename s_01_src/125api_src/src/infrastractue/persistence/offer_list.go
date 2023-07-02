package persistence

import (
	"context"

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
		return nil, err
	}
	if len(ol) == 0 {
		return []sqlc.OfferList{}, nil
	}

	return ol, nil
}
