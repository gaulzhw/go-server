package v1

import (
	"github.com/gaulzhw/go-server/internal/controller/v1/user"
	"github.com/gaulzhw/go-server/internal/store"
)

// Controller create a user handler used to handle request for user resource.
type Controller struct {
	users *user.Controller
}

// NewController creates a user handler.
func NewController(store store.Factory) *Controller {
	return &Controller{
		users: user.NewController(store),
	}
}

func (c *Controller) Users() *user.Controller {
	return c.users
}
