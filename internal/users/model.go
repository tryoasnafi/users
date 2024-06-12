package users

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	ID          uint           `gorm:"column:id;primarykey"`
	FirstName   string         `gorm:"column:first_name"`
	LastName    string         `gorm:"column:last_name"`
	Address     string         `gorm:"column:address"`
	DOB         datatypes.Date `gorm:"column:dob"`
	Email       string         `gorm:"column:email;unique"`
	PhoneNumber string         `gorm:"column:phone_number"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index"`
}
