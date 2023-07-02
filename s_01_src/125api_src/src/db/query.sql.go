// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package db

import (
	"context"
)

const accountMaster = `-- name: AccountMaster :one
SELECT account_id, account_name FROM account_master
WHERE account_id = $1 LIMIT 1
`

func (q *Queries) AccountMaster(ctx context.Context, accountID string) (AccountMaster, error) {
	row := q.db.QueryRowContext(ctx, accountMaster, accountID)
	var i AccountMaster
	err := row.Scan(&i.AccountID, &i.AccountName)
	return i, err
}

const offerLists = `-- name: OfferLists :many
SELECT offer_id, account_id FROM offer_list
`

func (q *Queries) OfferLists(ctx context.Context) ([]OfferList, error) {
	rows, err := q.db.QueryContext(ctx, offerLists)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []OfferList
	for rows.Next() {
		var i OfferList
		if err := rows.Scan(&i.OfferID, &i.AccountID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
