package users

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)


func SetRoutes(e *echo.Group, db *gorm.DB) {
	handler := NewUserHandler(db)
	userRoutes := e.Group("/users")
	userRoutes.GET("", handler.GetAllUsers)
	userRoutes.GET("/:id", handler.GetUserById)
	userRoutes.PUT("/:id", handler.UpdateUser)
	userRoutes.DELETE("/:id", handler.DeleteUser)
}
