package api

import (
	"net/http"

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
		userAuth.POST("/register", userHandler.Register)
		srv.GET("/check-system", func(ctx *gin.Context) {
			res, err := userHandler.AuthClient.SystemAvailableCheck()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "system is not available for connection",
				})
				return
			}
			ctx.JSON(http.StatusOK, res)

		})
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
