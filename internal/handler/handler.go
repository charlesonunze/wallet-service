package handler

import (
	"context"

	repo "github.com/charlesonunze/wallet-service/internal/db/repo"
	"github.com/charlesonunze/wallet-service/internal/service"
	walletpb "github.com/charlesonunze/wallet-service/pb/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

func (s *server) CreditUser(ctx context.Context, req *walletpb.CreditUserRequest) (*emptypb.Empty, error) {
	var res emptypb.Empty
	err := req.Validate()
	if err != nil {
		return &res, err
	}

	svc := s.GetService()
	err = svc.CreditUser(ctx, req.UserId, req.Amount)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

func (s *server) DebitUser(ctx context.Context, req *walletpb.DebitUserRequest) (*emptypb.Empty, error) {
	var res emptypb.Empty
	err := req.Validate()
	if err != nil {
		return &res, err
	}

	svc := s.GetService()
	err = svc.DebitUser(ctx, req.UserId, req.Amount)
	if err != nil {
		return &res, err
	}

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
		UserId:  b.UserID,
		Balance: b.Balance,
	}, nil
}
