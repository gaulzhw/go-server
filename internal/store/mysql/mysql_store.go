package mysql

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
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
func NewStore() (store.Factory, error) {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		opts.Username,
		opts.Password,
		opts.Host,
		opts.Database,
		true,
		"Local",
	)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)

	store := &mysqlStore{
		db:   db,
		user: user.NewStore(db),
	}

	if store == nil || err != nil {
		return nil, errors.WithMessagef(err, "failed to get mysql store, mysqlStore: %+v", store)
	}

	return store, nil
}
