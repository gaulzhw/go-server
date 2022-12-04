package user

import (
	"context"

	apisv1 "github.com/gaulzhw/go-server/pkg/apis/v1"
	metav1 "github.com/gaulzhw/go-server/pkg/meta/v1"
)

func (s *store) Get(ctx context.Context, username string, opts metav1.GetOptions) (*apisv1.User, error) {
	user := &apisv1.User{}
	err := s.db.Where("name = ? and status = 1", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
