package migration

import (
	"github.com/tryoasnafi/users/internal/users"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		users.User{},
	)
}