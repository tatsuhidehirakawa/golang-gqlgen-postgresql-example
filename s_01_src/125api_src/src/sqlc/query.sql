-- name: AccountMaster :one
SELECT * FROM account_master
WHERE account_id = $1 LIMIT 1;

-- name: OfferLists :many
SELECT * FROM offer_list;