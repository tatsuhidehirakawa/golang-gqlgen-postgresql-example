package persistence

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	sqlc "github.com/gqlgensamples/golang-gqlgen-postgresql-example/db"
)

func TestOfferListRepository_OfferLists(t *testing.T) {
	tests := []struct {
		name    string
		userID  string
		want    []sqlc.OfferList
		wantErr error
	}{
		{
			name:   "ok",
			userID: "101",
			want: []sqlc.OfferList{
				{OfferID: "1", AccountID: "101"},
				{OfferID: "2", AccountID: "102"},
				{OfferID: "3", AccountID: "103"},
				{OfferID: "4", AccountID: "104"},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := offerListRepo.OfferLists(context.Background())
			if diff := cmp.Diff(tt.wantErr, err); len(diff) != 0 {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.want, got); len(diff) != 0 {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
