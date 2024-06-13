package users

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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

//	GetAllUsers
//	@Summary	return all users in database
//	@Schemes
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	[]UserResponse
//	@Router		/v1/users [get]
func (h UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ListUserToResponse(users))
}

//	GetUserByID
//	@Summary	get user by given id
//	@Schemes
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	UserResponse
//	@Failure	400	{object}	MessageResponse
//	@Failure	404	{object}	MessageResponse
//	@Router		/v1/users/{id} [get]
//	@Param		id	path	int	true	"User ID"
func (h UserHandler) GetUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, ErrNeedUserID)
	}
	resp, err := h.service.GetUserById(uint(id))
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, MessageResponse{Message: err.Error()})
		}
		return err
	}
	return c.JSON(http.StatusOK, UserToResponse(resp))
}

//	Create new user
//	@Summary	create new user in the service
//	@Schemes
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Success	201			{object}	UserResponse
//	@Failure	400			{object}	MessageResponse
//	@Router		/v1/users 				[post]
//	@Param		request		body		CreateUserRequest	true	"User"
func (h UserHandler) CreateUser(c echo.Context) error {
	userReq := CreateUserRequest{}
	if err := c.Bind(&userReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrBadRequest)
	}

	if err := c.Validate(userReq); err != nil {
		return err
	}

	user, err := h.service.CreateUser(userReq)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, UserToResponse(user))
}

//	UpdateUserByID
//	@Summary	update user details with given id
//	@Schemes
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	UserResponse
//	@Failure	400	{object}	MessageResponse
//	@Failure	404	{object}	MessageResponse
//	@Router		/v1/users/{id} [put]
//	@Param		id		path	int					true	"User ID"
//	@Param		request	body	UpdateUserRequest	true	"New User Details"
func (h UserHandler) UpdateUser(c echo.Context) error {
	userReq := UpdateUserRequest{}
	if err := c.Bind(&userReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrBadRequest)
	}
	if err := c.Validate(userReq); err != nil {
		return err
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, MessageResponse{Message: ErrNeedUserID.Error()})
	}
	user, err := h.service.UpdateUser(uint(id), userReq)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, MessageResponse{Message: err.Error()})
		}
		return err
	}

	return c.JSON(http.StatusOK, UserToResponse(user))
}

//	DeleteUserByID
//	@Summary	delete a user with given id
//	@Schemes
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	UserResponse
//	@Failure	400	{object}	MessageResponse
//	@Failure	404	{object}	MessageResponse
//	@Router		/v1/users/{id}  [delete]
//	@Param		id	path	int	true	"User ID"
func (h UserHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrNeedUserID)
	}
	if err := h.service.DeleteUser(uint(id)); err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, MessageResponse{Message: err.Error()})
		}
		return err
	}
	return c.NoContent(http.StatusOK)
}
