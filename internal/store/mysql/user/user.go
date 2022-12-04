package user

import (
	"gorm.io/gorm"

	interfacestore "github.com/gaulzhw/go-server/internal/store"
)

type store struct {
	db *gorm.DB
}

var _ interfacestore.User = (*store)(nil)

func NewStore(db *gorm.DB) interfacestore.User {
	return &store{
		db: db,
	}
}
