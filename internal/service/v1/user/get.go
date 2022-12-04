package user

import (
	"context"

	apisv1 "github.com/gaulzhw/go-server/pkg/apis/v1"
	metav1 "github.com/gaulzhw/go-server/pkg/meta/v1"
)

func (s *service) Get(ctx context.Context, username string, opts metav1.GetOptions) (*apisv1.User, error) {
	user, err := s.store.Users().Get(ctx, username, opts)
	if err != nil {
		return nil, err
	}

	return user, nil
}
