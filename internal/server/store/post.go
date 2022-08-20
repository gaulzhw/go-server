package store

import (
	"context"

	"github.com/gaulzhw/go-server/internal/pkg/model/server/v1"
	metav1 "github.com/gaulzhw/go-server/pkg/meta/v1"
)

// PostStore defines the post storage interface.
type PostStore interface {
	Create(ctx context.Context, post *v1.Post, opts metav1.CreateOptions) error
	Update(ctx context.Context, post *v1.Post, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username string, postID string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, username string, postIDs []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username string, postID string, opts metav1.GetOptions) (*v1.Post, error)
	List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.PostList, error)
}
