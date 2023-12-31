// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: budget.sql

package db

import (
	"context"
)

const createBudget = `-- name: CreateBudget :one
INSERT INTO budgets (
  budget_name, budget_owner, amount
) VALUES (
  $1, $2, $3
)
RETURNING id, budget_name, amount, budget_owner, created_at
`

type CreateBudgetParams struct {
	BudgetName  string
	BudgetOwner string
	Amount      int64
}

func (q *Queries) CreateBudget(ctx context.Context, arg CreateBudgetParams) (Budget, error) {
	row := q.db.QueryRowContext(ctx, createBudget, arg.BudgetName, arg.BudgetOwner, arg.Amount)
	var i Budget
	err := row.Scan(
		&i.ID,
		&i.BudgetName,
		&i.Amount,
		&i.BudgetOwner,
		&i.CreatedAt,
	)
	return i, err
}

const deleteBudget = `-- name: DeleteBudget :exec
DELETE FROM budgets WHERE id = $1
`

func (q *Queries) DeleteBudget(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteBudget, id)
	return err
}

const getBudget = `-- name: GetBudget :one
SELECT id, budget_name, amount, budget_owner, created_at FROM budgets
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetBudget(ctx context.Context, id int32) (Budget, error) {
	row := q.db.QueryRowContext(ctx, getBudget, id)
	var i Budget
	err := row.Scan(
		&i.ID,
		&i.BudgetName,
		&i.Amount,
		&i.BudgetOwner,
		&i.CreatedAt,
	)
	return i, err
}

const listBudgets = `-- name: ListBudgets :many
SELECT id, budget_name, amount, budget_owner, created_at FROM budgets
WHERE budget_owner = $1
ORDER BY budget_name
LIMIT $2
OFFSET $3
`

type ListBudgetsParams struct {
	BudgetOwner string
	Limit       int32
	Offset      int32
}

func (q *Queries) ListBudgets(ctx context.Context, arg ListBudgetsParams) ([]Budget, error) {
	rows, err := q.db.QueryContext(ctx, listBudgets, arg.BudgetOwner, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Budget{}
	for rows.Next() {
		var i Budget
		if err := rows.Scan(
			&i.ID,
			&i.BudgetName,
			&i.Amount,
			&i.BudgetOwner,
			&i.CreatedAt,
		); err != nil {
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
