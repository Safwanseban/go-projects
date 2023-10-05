package interfaces

import (
	"github.com/Safwanseban/go-book-gateway/internal/types"
	"github.com/Safwanseban/go-book-gateway/internal/user/pb"
)

type AuthClient interface {
	SignUpUser(types.User) (*pb.RegisterResponse, error)
	LoginService(types.User) (uint, string)
}
