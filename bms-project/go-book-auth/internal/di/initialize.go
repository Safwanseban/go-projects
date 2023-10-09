package di

import (
	"github.com/Safwanseban/go-book-autha/internal/api"
	"github.com/Safwanseban/go-book-autha/internal/api/services"
)

func InitializeServer() *api.AuthGrpcServer {
	authService := services.NewAuthService()
	return api.NewGrpcServer(authService)
}
