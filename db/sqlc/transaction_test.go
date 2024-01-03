package db

import (
	"context"
	"testing"
	"time"

	"github.com/eliasmanj/budgets-api/utils"
	"github.com/stretchr/testify/require"
)

func createRandomTransaction(t *testing.T, usr User, acc Account, budget Budget) Transaction {
	arg := CreateTransactionParams{
		Description: utils.RandomString(15),
		BudgetID:    budget.ID,
		AccountID:   acc.ID,
		Amount:      utils.RandomMoney(),
	}
	tx, err := testQueries.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, tx)
	require.NotZero(t, tx.ID)
	require.Equal(t, arg.Description, tx.Description)
	require.Equal(t, arg.BudgetID, tx.BudgetID)
	require.Equal(t, arg.AccountID, tx.AccountID)
	require.Equal(t, arg.Amount, tx.Amount)
	return tx
}

func TestCreateTransaction(t *testing.T) {
	usr := createRandomUser(t)
	acc := createRandomAccount(t, usr)
	budget := createRandomBudget(t, usr)
	createRandomTransaction(t, usr, acc, budget)
}

func TestGetTransaction(t *testing.T) {
	usr := createRandomUser(t)
	acc := createRandomAccount(t, usr)
	budget := createRandomBudget(t, usr)
	arg := createRandomTransaction(t, usr, acc, budget)
	tx, err := testQueries.GetTransaction(context.Background(), arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, arg.ID, tx.ID)
	require.NotEmpty(t, arg.AccountID, tx.AccountID)
	require.NotEmpty(t, arg.BudgetID, tx.BudgetID)
	require.NotEmpty(t, arg.Description, tx.Description)
	require.WithinDuration(t, arg.CreatedAt, tx.CreatedAt, time.Second)
}

func TestListTransactionsByBudget(t *testing.T) {
	usr := createRandomUser(t)
	acc := createRandomAccount(t, usr)
	budget := createRandomBudget(t, usr)
	for i := 0; i < 10; i++ {
		createRandomTransaction(t, usr, acc, budget)
	}
	args := ListTransactionsByBudgetParams{
		BudgetID: budget.ID,
		Limit:    5,
		Offset:   0,
	}
	txs, err := testQueries.ListTransactionsByBudget(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, txs)
	for _, tx := range txs {
		require.NotEmpty(t, tx)
		require.NotZero(t, tx.ID)
		require.Equal(t, acc.ID, tx.AccountID)
		require.Equal(t, budget.ID, tx.BudgetID)
	}
}

func TestListTransactionsByAccount(t *testing.T) {
	usr := createRandomUser(t)
	acc := createRandomAccount(t, usr)
	budget := createRandomBudget(t, usr)
	for i := 0; i < 10; i++ {
		createRandomTransaction(t, usr, acc, budget)
	}
	args := ListTransactionsByAccountParams{
		AccountID: acc.ID,
		Limit:     5,
		Offset:    0,
	}
	txs, err := testQueries.ListTransactionsByAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, txs)
	for _, tx := range txs {
		require.NotEmpty(t, tx)
		require.NotZero(t, tx.ID)
		require.Equal(t, acc.ID, tx.AccountID)
		require.Equal(t, budget.ID, tx.BudgetID)
	}
}
