package db

import (
	"context"
	"testing"

	db "github.com/eliasmanj/budgets-api/db/sqlc"
	"github.com/eliasmanj/budgets-api/utils"
	"github.com/stretchr/testify/require"
)

func createRandomBudget(t *testing.T, user db.User) db.Budget {
	arg := db.CreateBudgetParams{
		BudgetName:  utils.RandomString(10),
		BudgetOwner: user.Username,
		Amount:      utils.RandomMoney(),
	}
	budget, err := testQueries.CreateBudget(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, budget)
	require.Equal(t, arg.Amount, budget.Amount)
	require.Equal(t, arg.BudgetName, budget.BudgetName)
	require.Equal(t, arg.BudgetOwner, budget.BudgetOwner)
	return budget
}

func TestCreateBudget(t *testing.T) {
	usr := createRandomUser(t)
	createRandomBudget(t, usr)
}

func TestListBudgets(t *testing.T) {
	usr := createRandomUser(t)
	for i := 0; i < 10; i++ {
		createRandomBudget(t, usr)
	}
	arg := db.ListBudgetsParams{
		BudgetOwner: usr.Username,
		Limit:       5,
		Offset:      0,
	}
	budgets, err := testQueries.ListBudgets(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, budgets)
	for _, budget := range budgets {
		require.NotEmpty(t, budget)
		require.Equal(t, arg.BudgetOwner, budget.BudgetOwner)
	}
}

func TestDeleteBudgets(t *testing.T) {
	usr := createRandomUser(t)
	budget := createRandomBudget(t, usr)
	err := testQueries.DeleteBudget(context.Background(), budget.ID)
	require.NoError(t, err)
}
