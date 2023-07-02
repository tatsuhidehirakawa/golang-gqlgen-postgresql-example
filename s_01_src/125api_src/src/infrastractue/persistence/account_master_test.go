package persistence

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	sqlc "github.com/gqlgensamples/golang-gqlgen-postgresql-example/db"
)

func TestAccountMasterRepo_AccountMaster(t *testing.T) {
	tests := []struct {
		name    string
		userID  string
		want    *sqlc.AccountMaster
		wantErr error
	}{
		{
			name:   "ok",
			userID: "101",
			want: &sqlc.AccountMaster{
				AccountID:   "101",
				AccountName: "account1",
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := accountMasterRepo.AccountMaster(context.Background(), tt.userID)
			if diff := cmp.Diff(tt.wantErr, err); len(diff) != 0 {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.want, got); len(diff) != 0 {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}

		})
	}

}
