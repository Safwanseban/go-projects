package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	db "github.com/Safwanseban/go-book-autha/internal/repo"
	"github.com/Safwanseban/go-book-autha/internal/types"
	"github.com/Safwanseban/go-book-autha/internal/user/pb"
	"google.golang.org/protobuf/types/known/anypb"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	query *db.Queries
}

func NewAuthService(queries *db.Queries) pb.AuthServiceServer {

	return &AuthServer{
		query: queries,
	}
}

// CheckSystem implements pb.AuthServiceServer.
func (*AuthServer) CheckSystem(context.Context, *pb.EmptyRequest) (*pb.SystemResponse, error) {
	return &pb.SystemResponse{
		Message: "system is ready for connection",
		Status:  http.StatusOK,
	}, nil
}

// mustEmbedUnimplementedAuthServiceServer implements pb.AuthServiceServer.
func (*AuthServer) mustEmbedUnimplementedAuthServiceServer() {
	panic("unimplemented")
}

// Login implements pb.AuthServiceServer.
func (a *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	user, err := a.query.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
		}, err
	}
	fmt.Println(user)
	return &pb.LoginResponse{
		Status: http.StatusOK,
	}, nil
}

// Register implements pb.AuthServiceServer.
func (a *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user types.User
	user.Password = req.Password
	if err := user.HashPassword(); err != nil {
		return nil, errors.New("error hashing password")

	}
	id, err := a.query.CreateUser(ctx, db.CreateUserParams{
		Username:    req.Username,
		Email:       req.Email,
		Password:    user.Password,
		PhoneNumber: sql.NullString{String: req.Phonenumber},
	})

	if err != nil {
		return nil, errors.New("error creating user")
	}

	return &pb.RegisterResponse{
		Result: map[string]*anypb.Any{"id": {Value: []byte(strconv.Itoa(int(id)))},
			"status": {Value: []byte("success")},
		},
	}, nil
}
