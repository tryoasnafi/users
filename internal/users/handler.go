package users

import (
	"github.com/labstack/echo/v4"
	"github.com/tryoasnafi/users/common"
	"gorm.io/gorm"
)

type UserHandler struct {
	service Service
}

func NewUserHandler(db *gorm.DB) UserHandler {
	repo := NewUserRepository(db)
	service := NewUserService(repo)
	return UserHandler{service: service}
}

func (h UserHandler) GetAllUsers(c echo.Context) error {
	return common.ErrNotImplemented
}

func (h UserHandler) GetUserById(c echo.Context) error {
	return common.ErrNotImplemented
}

func (h UserHandler) CreateUser(c echo.Context) error {
	return common.ErrNotImplemented
}

func (h UserHandler) UpdateUser(c echo.Context) error {
	return common.ErrNotImplemented
}

func (h UserHandler) DeleteUser(c echo.Context) error {
	return common.ErrNotImplemented
}
