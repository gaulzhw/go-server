package v1

import (
	"context"

	apisv1 "github.com/gaulzhw/go-server/pkg/apis/v1"
	metav1 "github.com/gaulzhw/go-server/pkg/meta/v1"
)

// User defines the user service interface.
type User interface {
	Get(ctx context.Context, username string, opts metav1.GetOptions) (*apisv1.User, error)
	Create(ctx context.Context, user *apisv1.User, opts metav1.CreateOptions) error
}
