package user

import (
	"context"

	"github.com/gaulzhw/go-server/pkg/apis/v1"
	metav1 "github.com/gaulzhw/go-server/pkg/meta/v1"
)

func (s *Service) Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error {
	if err := s.store.Users().Create(ctx, user, opts); err != nil {
		return err
	}
	return nil
}
