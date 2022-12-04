package options

import (
	"fmt"
	"time"

	"github.com/spf13/pflag"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gaulzhw/go-server/internal/store"
	mysqlstore "github.com/gaulzhw/go-server/internal/store/mysql"
)

type MySQLOptions struct {
	Host                  string
	Username              string
	Password              string
	Database              string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
}

func NewMySQLOptions() *MySQLOptions {
	return &MySQLOptions{}
}

func (o *MySQLOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Host, "mysql-host", "127.0.0.1",
		"MySQL service host address. Default to 127.0.0.1.")
	fs.StringVar(&o.Username, "mysql-username", "root",
		"Username for access to mysql service. Default to root.")
	fs.StringVar(&o.Password, "mysql-password", "123456",
		"Password for access to mysql. Default to 123456.")
	fs.StringVar(&o.Database, "mysql-database", "",
		"Database name for the server to use. Default to empty.")

	fs.IntVar(&o.MaxIdleConnections, "mysql-max-idle-connections", 100,
		"Maximum idle connections allowed to connect to mysql. Default to 100.")
	fs.IntVar(&o.MaxOpenConnections, "mysql-max-open-connections", 100,
		"Maximum open connections allowed to connect to mysql. Default to 100.")
	fs.DurationVar(&o.MaxConnectionLifeTime, "mysql-max-connection-life-time", time.Duration(10)*time.Second,
		"Maximum connection life time allowed to connect to mysql. Default to 10s.")
}

func (o *MySQLOptions) Validate() []error {
	var errs []error
	return errs
}

// NewClient create mysql store with the given config.
func (o *MySQLOptions) NewClient() (store.Factory, error) {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		o.Username,
		o.Password,
		o.Host,
		o.Database,
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
	sqlDB.SetMaxOpenConns(o.MaxOpenConnections)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(o.MaxConnectionLifeTime)
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(o.MaxIdleConnections)

	return mysqlstore.NewStoreFactory(db)
}
