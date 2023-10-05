package handlers

import (
	"net/http"

	interfaces "github.com/Safwanseban/go-book-gateway/internal/client/interface"
	"github.com/Safwanseban/go-book-gateway/internal/types"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthClient interfaces.AuthClient
}

func (s AuthHandler) Register(ctx *gin.Context) {

	var user types.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{

			"message": "bad request",
		})

		return
	}
	resp, err := s.AuthClient.SignUpUser(user)
	if err != nil {
		ctx.JSON(int(resp.Status), gin.H{
			"message": "internal server error",
		})
		return

	}
	ctx.JSON(http.StatusCreated, gin.H{
		"id":      resp.Id,
		"message": "success",
	})

}
