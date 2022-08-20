package v1

//go:generate mockgen -self_package=github.com/gaulzhw/go-server/internal/server/service/v1 -destination mock_service.go -package v1 github.com/gaulzhw/go-server/internal/server/service/v1 Service,UserSrv,PostSrv

import "github.com/gaulzhw/go-server/internal/server/store"

// Service defines functions used to return resource interface.
type Service interface {
	Users() UserSrv
	Posts() PostSrv
}

type service struct {
	store store.Factory
}

// NewService returns Service interface.
func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}

func (s *service) Users() UserSrv {
	return newUsers(s)
}

func (s *service) Posts() PostSrv {
	return newPosts(s)
}
