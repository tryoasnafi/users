package users

import (
	"net/http"

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
	resp, err := h.service.GetAllUsers()
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, resp)
	return nil
}

func (h UserHandler) GetUserById(c echo.Context) error {
	return common.ErrNotImplemented
}

func (h UserHandler) CreateUser(c echo.Context) error {
	userReq := CreateUserRequest{}
	if err := c.Bind(&userReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}
	user, err := h.service.CreateUser(userReq)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, UserToResponse(user))
	return nil
}

func (h UserHandler) UpdateUser(c echo.Context) error {
	return common.ErrNotImplemented
}

func (h UserHandler) DeleteUser(c echo.Context) error {
	return common.ErrNotImplemented
}
