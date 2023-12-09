package user

import (
	"github.com/gaulzhw/go-server/internal/store"
)

type service struct {
	store store.Factory
}

func NewService(store store.Factory) *service {
	return &service{
		store: store,
	}
}
