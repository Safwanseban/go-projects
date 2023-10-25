package handlers

import (
	"net/http"

	interfaces "github.com/Safwanseban/go-book-gateway/internal/client/interfaces"
	"github.com/Safwanseban/go-book-gateway/internal/types"
	"github.com/Safwanseban/go-book-gateway/utils"
	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	AuthClient interfaces.AuthClient
}

func NewAuthHandler(client interfaces.AuthClient) *AuthHandler {
	return &AuthHandler{
		AuthClient: client,
	}
}

func (s *AuthHandler) Register(ctx *gin.Context) {

	var user types.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{

			"message": "bad request",
		})

		return
	}
	err := validator.New().Struct(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": utils.ParseError(err),
		})
		return

	}
	result, err := s.AuthClient.SignUpUser(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return

	}

	ctx.JSON(http.StatusCreated, result)

}

func (a *AuthHandler) Login(ctx *gin.Context) {
	var user types.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{

			"message": "bad request",
		})

		return
	}
	resp, err := a.AuthClient.LoginService(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{

			"message": "authserver error",
		})

		return
	}
	if resp.Status != http.StatusOK {
		ctx.JSON(http.StatusInternalServerError, gin.H{

			"message": "internal server error",
		})

		return
	}
	ctx.JSON(http.StatusOK, resp)

}
