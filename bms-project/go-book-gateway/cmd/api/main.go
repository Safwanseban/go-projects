package main

import (
	"github.com/Safwanseban/go-book-gateway/internal/configs"
	"github.com/gin-gonic/gin"
)

func init() {

}
func main() {
	config := configs.NewConfig()
	srv := gin.Default()
	if err := srv.Run(config.String("Port")); err != nil {
		panic("error initializing server")
	}

}
