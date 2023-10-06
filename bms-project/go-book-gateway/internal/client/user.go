package client

import (
	"context"
	"fmt"
	"log"

	interfaces "github.com/Safwanseban/go-book-gateway/internal/client/interfaces"
	"github.com/Safwanseban/go-book-gateway/internal/types"
	"github.com/Safwanseban/go-book-gateway/internal/user/pb"
	"github.com/knadh/koanf/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthServiceClient struct {
	AuthClient pb.AuthServiceClient
}

// SystemAvailableCheck implements interfaces.AuthClient.
func (auth *AuthServiceClient) SystemAvailableCheck() (*pb.SystemResponse, error) {
	res, err := auth.AuthClient.CheckSystem(context.Background(), &pb.EmptyRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SignUpUser implements interfaces.AuthClient.
func (auth *AuthServiceClient) SignUpUser(user types.User) (*pb.RegisterResponse, error) {
	fmt.Println("here")
	resp, err := auth.AuthClient.Register(context.Background(), &pb.RegisterRequest{
		// Username:    user.UserName,
		Email:    user.Email,
		Password: user.Password,
		// Phonenumber: user.Password,
	})
	fmt.Println(resp, err)
	if err != nil {
		return resp, err

	}
	return resp, nil

}

// LoginService implements interfaces.AuthClient.
func (auth *AuthServiceClient) LoginService(user types.User) (uint, string) {

	auth.AuthClient.Login(context.Background(), &pb.LoginRequest{
		Email:    user.Email,
		Password: user.Password,
	})
	return 0, ""
}

func NewUserClient(config *koanf.Koanf) interfaces.AuthClient {

	connection, err := grpc.Dial(config.String("userAuthUrl"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error connecting to authservice")

	}

	grpcClient := pb.NewAuthServiceClient(connection)
	return &AuthServiceClient{
		AuthClient: grpcClient,
	}
}
