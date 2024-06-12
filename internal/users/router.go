package users

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)


func SetRoutes(e *echo.Group, db *gorm.DB) {
	handler := NewUserHandler(db)
	e.GET("/", handler.GetAllUsers)
	e.GET("/:id", handler.GetUserById)
	e.PUT("/:id", handler.UpdateUser)
	e.DELETE("/:id", handler.DeleteUser)
}
