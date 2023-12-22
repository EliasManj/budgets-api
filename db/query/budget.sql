-- name: GetBudget :one
SELECT * FROM budgets
WHERE id = $1 LIMIT 1;

-- name: ListBudgets :many
SELECT * FROM budgets
WHERE budget_owner = $1
ORDER BY budget_name
LIMIT $2
OFFSET $3;

-- name: CreateBudget :one
INSERT INTO budgets (
  budget_name, budget_owner, amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteBudget :exec
DELETE FROM budgets WHERE id = $1;