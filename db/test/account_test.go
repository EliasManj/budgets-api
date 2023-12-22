package db

import (
	"context"
	"testing"

	db "github.com/eliasmanj/budgets-api/db/sqlc"
	"github.com/eliasmanj/budgets-api/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T, usr db.User) db.Account {
	arg := db.CreateAccountParams{
		AccountName:  utils.RandomString(7),
		AccountOwner: usr.Username,
		AccountType:  utils.RandomAccountType(),
		Balance:      utils.RandomMoney(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.AccountName, account.AccountName)
	require.Equal(t, arg.AccountOwner, account.AccountOwner)
	require.Equal(t, arg.AccountType, account.AccountType)
	require.Equal(t, arg.Balance, account.Balance)
	return account
}

func TestCreateAccount(t *testing.T) {
	usr := createRandomUser(t)
	createRandomAccount(t, usr)
}

func TestListAccounts(t *testing.T) {
	usr := createRandomUser(t)
	for i := 0; i < 10; i++ {
		createRandomAccount(t, usr)
	}
	arg := db.ListAccountsParams{
		AccountOwner: usr.Username,
		Limit:        5,
		Offset:       0,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, arg.AccountOwner, account.AccountOwner)
	}
}

func TestDeleteAccount(t *testing.T) {
	usr := createRandomUser(t)
	account := createRandomAccount(t, usr)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)
}
