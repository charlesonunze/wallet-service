package service

import (
	"context"
	"errors"

	repo "github.com/charlesonunze/wallet-service/internal/db/repo"
)

// WalletService - interface for the wallet service
type WalletService interface {
	CreditUser(ctx context.Context, userID, amount int64) (int64, error)
	DebitUser(ctx context.Context, userID, amount int64) (int64, error)
	GetUserBalance(ctx context.Context, userID int64) (repo.Wallet, error)
}

type walletService struct {
	repo *repo.Queries
}

// New - returns an instance of the WalletService
func New(repo *repo.Queries) WalletService {
	return &walletService{
		repo,
	}
}

func (w *walletService) CreditUser(ctx context.Context, userID, amount int64) (int64, error) {
	var balance int64
	wallet, err := w.repo.GetWallet(ctx, userID)
	if err != nil {
		return balance, err
	}

	wallet, err = w.repo.UpdateWallet(ctx, repo.UpdateWalletParams{
		UserID:  userID,
		Balance: wallet.Balance + amount,
	})
	if err != nil {
		return balance, err
	}

	return wallet.Balance, nil
}

func (w *walletService) DebitUser(ctx context.Context, userID, amount int64) (int64, error) {
	var balance int64
	wallet, err := w.repo.GetWallet(ctx, userID)
	if err != nil {
		return balance, err
	}

	if amount > wallet.Balance {
		return balance, errors.New("insufficient funds")
	}

	wallet, err = w.repo.UpdateWallet(ctx, repo.UpdateWalletParams{
		UserID:  userID,
		Balance: wallet.Balance - amount,
	})
	if err != nil {
		return balance, err
	}

	return wallet.Balance, nil
}

func (w *walletService) GetUserBalance(ctx context.Context, userID int64) (repo.Wallet, error) {
	wallet, err := w.repo.GetWallet(ctx, userID)
	if err != nil {
		return wallet, err
	}

	return wallet, nil
}
