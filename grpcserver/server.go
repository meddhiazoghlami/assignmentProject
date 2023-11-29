package grpcserver

import (
	"context"
	"database/sql"
	"log"

	pb "github.com/meddhiazoghlami/assignmentProject/proto"
	"github.com/meddhiazoghlami/assignmentProject/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	pb.UnimplementedGetBalanceServer
	Db *sql.DB
}

// GetWalletBalance implements grpcservice.GetBalanceServer
func (s *GrpcServer) GetWalletBalance(ctx context.Context, in *pb.GetBalanceRequest) (*pb.GetBalanceReply, error) {
	log.Printf("Received: %v also: %v", in.GetUserId(), in.GetWalletId())

	wallet, err := services.GetBalance(s.Db, in.GetWalletId(), in.GetUserId())
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	b, _ := wallet.Balance.Float64()
	return &pb.GetBalanceReply{Balance: float32(b)}, nil
}
