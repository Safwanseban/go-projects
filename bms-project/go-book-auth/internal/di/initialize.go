package di

import (
	"fmt"

	"github.com/Safwanseban/go-book-autha/configs"
	"github.com/Safwanseban/go-book-autha/internal/api"
	"github.com/Safwanseban/go-book-autha/internal/api/services"
	repo "github.com/Safwanseban/go-book-autha/internal/repo"
	"github.com/Safwanseban/go-book-autha/pg/db"
)

func InitializeServer() *api.AuthGrpcServer {
	config := configs.NewConfig()
	fmt.Println(config.String("user"))
	dbConn := db.ConnectTodDb(config)
	queries := repo.New(dbConn)
	authService := services.NewAuthService(queries)
	return api.NewGrpcServer(authService)
}
