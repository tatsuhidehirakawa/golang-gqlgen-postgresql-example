package persistence

import (
	"context"
	"os"
	"testing"

	"github.com/gqlgensamples/golang-gqlgen-postgresql-example/config"
	"github.com/gqlgensamples/golang-gqlgen-postgresql-example/io"
)

var (
	userRepo          *UserRepository
	accountMasterRepo *MasterAccountRepository
)

func TestMain(m *testing.M) {
	cfg, _ := config.LoadConfig(context.Background())
	db, _ := io.NewSQLdatabase(cfg)
	userRepo = NewUserRepository(db)
	accountMasterRepo = NewMasterAccountRepository(db.SQLC)

	res := m.Run()

	os.Exit(res)

}
