package v1

import (
	"github.com/gaulzhw/go-server/internal/service/v1/user"
	"github.com/gaulzhw/go-server/internal/store"
)

type Service interface {
	Users() User
}

type service struct {
	users User
}

var _ Service = (*service)(nil)

// NewService returns Service interface.
func NewService(store store.Factory) Service {
	return &service{
		users: user.NewService(store),
	}
}

func (s *service) Users() User {
	return s.users
}
