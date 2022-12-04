package v1

import (
	"github.com/gaulzhw/go-server/internal/service/v1/user"
	"github.com/gaulzhw/go-server/internal/store"
)

type Service struct {
	users *user.Service
}

// NewService returns Service interface.
func NewService(store store.Factory) *Service {
	return &Service{
		users: user.NewService(store),
	}
}

func (s *Service) Users() *user.Service {
	return s.users
}
