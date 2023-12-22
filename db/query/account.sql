-- name: CreateAccount :one
INSERT INTO accounts (
    account_name, balance, account_owner, account_type
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
WHERE account_owner = $1
order by account_name
LIMIT $2
OFFSET $3;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;

-- name: DeleteAllUserAccounts :exec
DELETE FROM accounts WHERE account_owner = $1;

-- name: UpdateAccountBalance :one
UPDATE accounts SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;