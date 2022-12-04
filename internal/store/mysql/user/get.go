package user

import (
	"context"

	"github.com/gaulzhw/go-server/pkg/apis/v1"
	metav1 "github.com/gaulzhw/go-server/pkg/meta/v1"
)

// Create creates a new user account.
func (s *Store) Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error) {
	user := &v1.User{}
	err := s.db.Where("name = ? and status = 1", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
