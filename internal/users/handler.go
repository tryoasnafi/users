package users

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "need user id parameter")
	}
	resp, err := h.service.GetUserById(uint(id))
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, UserToResponse(resp))
	return nil
}

func (h UserHandler) CreateUser(c echo.Context) error {
	userReq := CreateUserRequest{}
	if err := c.Bind(&userReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}
	
	if err := c.Validate(userReq); err != nil {
		return err
	}

	user, err := h.service.CreateUser(userReq)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, UserToResponse(user))
	return nil
}

func (h UserHandler) UpdateUser(c echo.Context) error {
	userReq := UpdateUserRequest{}
	if err := c.Bind(&userReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}
	if err := c.Validate(userReq); err != nil {
		return err
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "need user id parameter")
	}
	user, err := h.service.UpdateUser(uint(id), userReq)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, UserToResponse(user))
	return nil
}

func (h UserHandler) DeleteUser(c echo.Context) error {
	return common.ErrNotImplemented
}
