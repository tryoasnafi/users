package users

import (
	"gorm.io/gorm"
)

type Repository interface {
	All() ([]User, error)
	FindById(id uint) (User, error)
	Add(user *User) error
	Update(user *User) error
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
	user := User{}
	result := repo.DB.First(&user, id)
	return user, result.Error
}

func (repo UserRepository) Add(user *User) error {
	return repo.DB.Create(user).Error
}

func (repo UserRepository) Update(user *User) error {
	return repo.DB.Updates(user).Error
}
func (repo UserRepository) Delete(id uint) error {
	result := repo.DB.Delete(&User{}, id)
	if result.RowsAffected == 0 {
		return ErrUserNotFound
	}
	return nil
}
