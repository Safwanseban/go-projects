package api

import (
	"fmt"
	"log"
	"net"

	"github.com/Safwanseban/go-book-autha/internal/user/pb"
	"google.golang.org/grpc"
)

type AuthGrpcServer struct {
	AuthServer *grpc.Server
	listner    net.Listener
}

func NewGrpcServer(authService pb.AuthServiceServer) *AuthGrpcServer {

	srv := grpc.NewServer()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("error listning to server")
	}
	pb.RegisterAuthServiceServer(srv, authService)
	return &AuthGrpcServer{
		AuthServer: srv,
		listner:    lis,
	}

}
func (s *AuthGrpcServer) Start() {
	fmt.Println("starting auth server on 50051 port")
	if err := s.AuthServer.Serve(s.listner); err != nil {
		panic("error initializing grpc server")
	}
}
