package users

import (
	"net/http"

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

func (h UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {

}

func (h UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {

}

func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

}

func (h UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (h UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
