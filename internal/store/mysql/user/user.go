package user

import (
	"gorm.io/gorm"

	globalstore "github.com/gaulzhw/go-server/internal/store"
)

type Store struct {
	db *gorm.DB
}

var _ globalstore.User = (*Store)(nil)

func NewStore(db *gorm.DB) globalstore.User {
	return &Store{
		db: db,
	}
}
