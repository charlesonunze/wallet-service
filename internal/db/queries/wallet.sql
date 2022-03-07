-- name: CreateWallet :one
INSERT INTO wallets (user_id, balance)
VALUES ($1, $2) 
RETURNING *;

-- name: GetWallet :one
SELECT * FROM wallets
WHERE user_id = $1 LIMIT 1;

-- name: UpdateWallet :one
UPDATE wallets
SET balance = $2
WHERE user_id = $1
RETURNING *;
