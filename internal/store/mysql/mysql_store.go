package mysql

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/gaulzhw/go-server/internal/store"
	"github.com/gaulzhw/go-server/internal/store/mysql/user"
)

type mysqlStore struct {
	db *gorm.DB

	user store.User
}

var _ store.Factory = (*mysqlStore)(nil)

func (ds *mysqlStore) Users() store.User {
	return ds.user
}

func (ds *mysqlStore) Close() error {
	db, err := ds.db.DB()
	if err != nil {
		return errors.Wrap(err, "get gorm db instance failed")
	}
	return db.Close()
}

// NewStore create mysql store with the given config.
func NewStore(db *gorm.DB) (store.Factory, error) {
	store := &mysqlStore{
		db:   db,
		user: user.NewStore(db),
	}
	return store, nil
}
