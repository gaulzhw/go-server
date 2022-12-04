package store

import (
	"context"

	"github.com/gaulzhw/go-server/pkg/apis/v1"
	metav1 "github.com/gaulzhw/go-server/pkg/meta/v1"
)

// User defines the user storage interface.
type User interface {
	Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error)
	Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error
}
