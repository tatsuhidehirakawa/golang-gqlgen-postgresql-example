package repository

import (
	"context"

	sqlc "github.com/gqlgensamples/golang-gqlgen-postgresql-example/db"
)

type IOfferListRepository interface {
	OfferLists(ctx context.Context) ([]sqlc.OfferList, error)
}
