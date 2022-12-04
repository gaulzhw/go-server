package mysql

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/gaulzhw/go-server/internal/store"
	"github.com/gaulzhw/go-server/internal/store/mysql/user"
)

type mysqlStoreFactory struct {
	db *gorm.DB

	user store.User
}

var _ store.Factory = (*mysqlStoreFactory)(nil)

func (ds *mysqlStoreFactory) Users() store.User {
	return ds.user
}

func (ds *mysqlStoreFactory) Close() error {
	db, err := ds.db.DB()
	if err != nil {
		return errors.Wrap(err, "get gorm db instance failed")
	}
	return db.Close()
}

// NewStoreFactory create mysql store with the given config.
func NewStoreFactory(db *gorm.DB) (store.Factory, error) {
	store := &mysqlStoreFactory{
		db:   db,
		user: user.NewStore(db),
	}
	return store, nil
}
