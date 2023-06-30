// O/Rマッパー
package persistence

import (
	"context"
	"database/sql"

	"github.com/gqlgensamples/golang-gqlgen-postgresql-example/domain/entity"
	"github.com/gqlgensamples/golang-gqlgen-postgresql-example/domain/repository"
	"github.com/gqlgensamples/golang-gqlgen-postgresql-example/io"
)

type UserRepository struct {
	database *io.SQLdatabase
}

var _ repository.IUserRepository = (*UserRepository)(nil)

func NewUserRepository(db *io.SQLdatabase) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (r *UserRepository) GetUser(ctx context.Context, userID string) (*entity.User, error) {
	var user entity.User
	// err := r.database.SQLX.Get(&user, "SELECT id, name FROM users WHERE id=$1", userID)
	err := r.database.SQLX.Get(&user, "SELECT offer_id, account_id FROM offer_list WHERE offer_id=$1", userID)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, nil
		default:
			return nil, err
		}
	}
	return &user, nil
}
