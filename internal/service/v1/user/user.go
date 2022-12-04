package user

import (
	"github.com/gaulzhw/go-server/internal/store"
)

type Service struct {
	store store.Factory
}

func NewService(store store.Factory) *Service {
	return &Service{
		store: store,
	}
}
