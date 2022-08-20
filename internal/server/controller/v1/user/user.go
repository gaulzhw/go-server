package user

import (
	srvv1 "github.com/gaulzhw/go-server/internal/server/service/v1"
	"github.com/gaulzhw/go-server/internal/server/store"
)

// UserController create a user handler used to handle request for user resource.
type UserController struct {
	srv srvv1.Service
}

// NewUserController creates a user handler.
func NewUserController(store store.Factory) *UserController {
	return &UserController{
		srv: srvv1.NewService(store),
	}
}
