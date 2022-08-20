package post

import (
	srvv1 "github.com/gaulzhw/go-server/internal/server/service/v1"
	"github.com/gaulzhw/go-server/internal/server/store"
)

// PostController create a post handler used to handle request for post resource.
type PostController struct {
	srv srvv1.Service
}

// NewPostController creates a post handler.
func NewPostController(store store.Factory) *PostController {
	return &PostController{
		srv: srvv1.NewService(store),
	}
}
