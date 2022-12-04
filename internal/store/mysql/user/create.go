package user

import (
	"context"

	"github.com/gaulzhw/go-server/pkg/apis/v1"
	metav1 "github.com/gaulzhw/go-server/pkg/meta/v1"
)

// Create creates a new user account.
func (s *Store) Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error {
	return s.db.Create(&user).Error
}
