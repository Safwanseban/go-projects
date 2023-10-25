package interfaces

import (
	"github.com/Safwanseban/go-book-gateway/internal/types"
	"github.com/Safwanseban/go-book-gateway/internal/user/pb"
)

type AuthClient interface {
	SignUpUser(types.User) (map[string]any, error)
	LoginService(types.User) (*pb.LoginResponse, error)
	SystemAvailableCheck() (*pb.SystemResponse, error)
}
