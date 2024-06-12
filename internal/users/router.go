package users

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) http.Handler {
	handler := NewUserHandler(db)
	r := chi.NewRouter()
	r.Get("/", handler.GetAllUsers)
	r.Get("/{userID}", handler.GetUserById)
	r.Put("/{userID}", handler.UpdateUser)
	r.Delete("/{userID}", handler.DeleteUser)
	return r
}
