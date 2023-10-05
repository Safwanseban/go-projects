package api

import (
	handlers "github.com/Safwanseban/go-book-gateway/internal/api/handlers/user"
	"github.com/gin-gonic/gin"
)

type AppServer struct {
	engine *gin.Engine
}

func NewAppServer(userHandler *handlers.AuthHandler) *AppServer {
	srv := gin.Default()
	userAuth := srv.Group("/auth")
	{
		userAuth.GET("/register", userHandler.Register)
	}
	return &AppServer{
		engine: srv,
	}

}

func (app *AppServer) Start() {

	if err := app.engine.Run(":3000"); err != nil {
		panic("error initializing server")
	}
}
