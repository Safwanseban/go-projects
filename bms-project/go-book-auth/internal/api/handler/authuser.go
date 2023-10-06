package handler

import (
	"context"
	"net/http"

	"github.com/Safwanseban/go-book-autha/internal/user/pb"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
}

// CheckSystem implements pb.AuthServiceServer.
func (AuthServer) CheckSystem(context.Context, *pb.EmptyRequest) (*pb.SystemResponse, error) {
	return &pb.SystemResponse{
		Message: "system is ready for connection",
		Status:  http.StatusOK,
	}, nil
}

// mustEmbedUnimplementedAuthServiceServer implements pb.AuthServiceServer.
func (AuthServer) mustEmbedUnimplementedAuthServiceServer() {
	panic("unimplemented")
}

// Login implements pb.AuthServiceServer.
func (AuthServer) Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error) {
	panic("unimplemented")
}

// Register implements pb.AuthServiceServer.
func (AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return &pb.RegisterResponse{
		Id:     1,
		Status: 200,
	}, nil
}

