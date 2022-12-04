package user

import (
	"context"

	apisv1 "github.com/gaulzhw/go-server/pkg/apis/v1"
	metav1 "github.com/gaulzhw/go-server/pkg/meta/v1"
)

func (s *store) Create(ctx context.Context, user *apisv1.User, opts metav1.CreateOptions) error {
	return s.db.Create(&user).Error
}
