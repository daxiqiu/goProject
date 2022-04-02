package db

import (
	"context"
	"database/sql"
	"goProject/until"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Blance:   until.RandomMoney(),
		Currency: until.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Blance, account.Blance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	a := createRandomAccount(t)
	account, err := testQueries.GetAccount(context.Background(), a.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, a.ID, account.ID)
	require.Equal(t, a.Owner, account.Owner)
	require.Equal(t, a.Blance, account.Blance)
	require.Equal(t, a.Currency, account.Currency)

	require.WithinDuration(t, a.CreatedAt, account.CreatedAt, time.Second)

}

func TestUpdateAccount(t *testing.T) {
	a := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:     a.ID,
		Blance: until.RandomMoney(),
	}

	_, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)

}

func TestDeleteAccount(t *testing.T) {
	a := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), a.ID)
	require.NoError(t, err)

	a2, err := testQueries.GetAccount(context.Background(), a.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, a2)

}

func TestListAccount(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}
