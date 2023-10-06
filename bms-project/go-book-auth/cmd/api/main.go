package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Safwanseban/go-book-autha/internal/api/handler"
	"github.com/Safwanseban/go-book-autha/internal/user/pb"
	"google.golang.org/grpc"
)

func main() {
	var authSrv handler.AuthServer
	server := grpc.NewServer()
	pb.RegisterAuthServiceServer(server, authSrv)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("error listneing to server")
	}

	fmt.Println("starting the grpc server")
	if err := server.Serve(lis); err != nil {
		panic("error initializing auth server")
	}

}
