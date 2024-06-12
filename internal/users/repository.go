package users

import (
	"github.com/tryoasnafi/users/common"
	"gorm.io/gorm"
)

type Repository interface {
	All() ([]User, error)
	FindById(id uint) (User, error)
	Add(user *User) error
	Update(user User) error
	Delete(id uint) error
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{DB: db}
}

func (repo UserRepository) All() ([]User, error) {
	users := []User{}
	result := repo.DB.Find(&users)
	return users, result.Error
}

func (repo UserRepository) FindById(id uint) (User, error) {
	return User{}, common.ErrNotImplemented
}

func (repo UserRepository) Add(user *User) error {
	return repo.DB.Create(user).Error
}

func (repo UserRepository) Update(user User) error {
	return common.ErrNotImplemented
}
func (repo UserRepository) Delete(id uint) error {
	return common.ErrNotImplemented
}
