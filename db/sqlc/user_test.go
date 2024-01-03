package db

import (
	"context"
	"testing"

	"github.com/eliasmanj/budgets-api/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       utils.RandomUser(),
		Email:          utils.RandomEmail(),
		HashedPassword: utils.RandomString(10),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	return user
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}
	arg := ListUsersParams{
		Limit:  5,
		Offset: 0,
	}
	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, users)

	for _, user := range users {
		require.NotEmpty(t, user)
	}

}

func TestDeleteUser(t *testing.T) {
	user := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user.Username)
	require.NoError(t, err)
}

func TestGetUser(t *testing.T) {
	arg := createRandomUser(t)
	user, err := testQueries.GetUser(context.Background(), arg.Username)
	require.NoError(t, err)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestCreateDuplicateEmailUser(t *testing.T) {
	email := utils.RandomEmail()
	arg1 := CreateUserParams{
		Username:       utils.RandomUser(),
		Email:          email,
		HashedPassword: utils.RandomString(10),
	}
	_, err := testQueries.CreateUser(context.Background(), arg1)
	require.NoError(t, err)
	arg2 := CreateUserParams{
		Username:       utils.RandomUser(),
		Email:          email,
		HashedPassword: utils.RandomString(10),
	}
	_, err = testQueries.CreateUser(context.Background(), arg2)
	require.Error(t, err)
}
