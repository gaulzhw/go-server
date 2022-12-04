package user

import (
	svcv1 "github.com/gaulzhw/go-server/internal/service/v1"
	"github.com/gaulzhw/go-server/internal/store"
)

// Controller create a user handler used to handle request for user resource.
type Controller struct {
	svc svcv1.Service
}

// NewController creates a user handler.
func NewController(store store.Factory) *Controller {
	return &Controller{
		svc: svcv1.NewService(store),
	}
}
