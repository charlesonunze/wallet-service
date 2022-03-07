package main

import (
	"database/sql"
	"log"
	"net"
	"os"

	"github.com/charlesonunze/wallet-service/internal/handler"
	walletpb "github.com/charlesonunze/wallet-service/pb/v1"

	db "github.com/charlesonunze/wallet-service/internal/db/repo"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

var (
	dbDriver  = os.Getenv("DB_DRIVER")
	dbSource  = os.Getenv("DB_URI")
	GRPC_PORT = os.Getenv("GRPC_PORT")
	repo      *db.Queries
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	repo = db.New(conn)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	walletpb.RegisterWalletServiceServer(s, handler.New(repo))
	s.Serve(lis)
}
