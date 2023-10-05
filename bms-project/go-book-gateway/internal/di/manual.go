package di

import (
	"github.com/Safwanseban/go-book-gateway/internal/api"
	handlers "github.com/Safwanseban/go-book-gateway/internal/api/handlers/user"
	"github.com/Safwanseban/go-book-gateway/internal/client"
	"github.com/Safwanseban/go-book-gateway/internal/configs"
)

func InitializeApi() *api.AppServer {

	config := configs.NewConfig()
	client := client.NewUserClient(config)
	userHandler := handlers.NewAuthHandler(client)
	return api.NewAppServer(userHandler)

}
