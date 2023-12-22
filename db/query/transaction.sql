-- name: GetTransaction :one
SELECT * FROM transactions
WHERE id = $1 LIMIT 1;

-- name: ListTransactionsByBudget :many
SELECT * FROM transactions
WHERE budget_id = $1
ORDER BY created_at
LIMIT $2
OFFSET $3;

-- name: ListTransactionsByAccount :many
SELECT * FROM transactions
WHERE account_id = $1
ORDER BY created_at
LIMIT $2
OFFSET $3;

-- name: CreateTransaction :one
INSERT INTO transactions (
  description, budget_id, account_id, amount
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: DeleteTransaction :exec
DELETE FROM transactions WHERE id = $1;