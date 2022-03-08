package handler

import (
	"context"

	repo "github.com/charlesonunze/wallet-service/internal/db/repo"
	"github.com/charlesonunze/wallet-service/internal/service"
	walletpb "github.com/charlesonunze/wallet-service/pb/v1"
)

type server struct {
	repo *repo.Queries
}

// New - returns an instance of the WalletServiceServer
func New(repo *repo.Queries) walletpb.WalletServiceServer {
	return &server{
		repo,
	}
}

func (s *server) GetService() service.WalletService {
	return service.New(s.repo)
}

func (s *server) CreditUser(ctx context.Context, req *walletpb.CreditUserRequest) (*walletpb.CreditUserResponse, error) {
	var res walletpb.CreditUserResponse
	err := req.Validate()
	if err != nil {
		return &res, err
	}

	svc := s.GetService()
	balance, err := svc.CreditUser(ctx, req.UserId, req.Amount)
	if err != nil {
		return &res, err
	}

	res.Balance = balance
	return &res, nil
}

func (s *server) DebitUser(ctx context.Context, req *walletpb.DebitUserRequest) (*walletpb.DebitUserResponse, error) {
	var res walletpb.DebitUserResponse
	err := req.Validate()
	if err != nil {
		return &res, err
	}

	svc := s.GetService()
	balance, err := svc.DebitUser(ctx, req.UserId, req.Amount)
	if err != nil {
		return &res, err
	}

	res.Balance = balance
	return &res, nil
}

func (s *server) GetUserBalance(ctx context.Context, req *walletpb.GetUserBalanceRequest) (*walletpb.GetUserBalanceResponse, error) {
	var result walletpb.GetUserBalanceResponse
	err := req.Validate()
	if err != nil {
		return &result, err
	}

	svc := s.GetService()
	b, err := svc.GetUserBalance(ctx, req.UserId)
	if err != nil {
		return &result, err
	}

	return &walletpb.GetUserBalanceResponse{
		Balance: b.Balance,
	}, nil
}
