package db

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomWallet(t *testing.T, userID int64) Wallet {
	arg := CreateWalletParams{
		UserID:  userID,
		Balance: int64(gofakeit.Number(1000, 10000)),
	}

	wallet, err := testQueries.CreateWallet(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, wallet)
	require.Equal(t, wallet.UserID, arg.UserID)
	require.Equal(t, wallet.Balance, arg.Balance)

	return wallet
}

func TestCreateWallet(t *testing.T) {
	user := createRandomUser(t)
	createRandomWallet(t, user.ID)
}

func TestGetWallet(t *testing.T) {
	user := createRandomUser(t)
	wallet1 := createRandomWallet(t, user.ID)

	wallet2, err := testQueries.GetWallet(context.Background(), wallet1.UserID)

	require.NoError(t, err)
	require.NotEmpty(t, wallet2)

	require.Equal(t, wallet1.ID, wallet2.ID)
	require.Equal(t, wallet1.UserID, wallet2.UserID)
	require.Equal(t, wallet1.Balance, wallet2.Balance)
	require.WithinDuration(t, wallet1.CreatedAt, wallet2.CreatedAt, time.Second)
}

func TestUpdateWallet(t *testing.T) {
	user := createRandomUser(t)
	wallet := createRandomWallet(t, user.ID)

	arg := UpdateWalletParams{
		UserID:  user.ID,
		Balance: int64(gofakeit.Number(100, 900)),
	}

	updatedWallet, err := testQueries.UpdateWallet(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updatedWallet)

	require.Equal(t, wallet.ID, updatedWallet.ID)
	require.Equal(t, wallet.UserID, updatedWallet.UserID)
	require.Equal(t, arg.Balance, updatedWallet.Balance)
	require.WithinDuration(t, wallet.CreatedAt, updatedWallet.CreatedAt, time.Second)
}
