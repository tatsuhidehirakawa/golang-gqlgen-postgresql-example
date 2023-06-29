-- name: GetAccountMaster :one
SELECT * FROM account_master
WHERE account_id = $1 LIMIT 1;

-- name: GetOfferList :one
SELECT * FROM offer_list
WHERE offer_id = $1 LIMIT 1;