package main

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/tryoasnafi/users/common"
	"github.com/tryoasnafi/users/database"
	"github.com/tryoasnafi/users/internal/middlewares"
	"github.com/tryoasnafi/users/internal/users"
	"github.com/tryoasnafi/users/internal/validation"

	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/tryoasnafi/users/docs"
)

func init() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file", err)
	}
}

//	@title			User Details Service API
//	@version		1.0
//	@description	This is a user details service API.
//	@host			localhost:9090
//	@BasePath		/api
func main() {
	db, err := database.GetDB()
	if err != nil {
		log.Fatal("failed connect to database", err)
	}

	addr := fmt.Sprintf(":%s", common.Getenv("APP_PORT", "8080"))

	e := echo.New()
	e.Validator = &validation.CustomValidator{Validator: validator.New()}

	apiRoute := e.Group("/api")
	apiRoute.GET("/docs/*", echoSwagger.WrapHandler)
	
	v1 := apiRoute.Group("/v1")
	v1.Use(echo.WrapMiddleware(middlewares.ApiVersionCtx("1")))
	users.SetRoutes(v1, db)

	// starting the server
	log.Println("Starting user service on port", addr)
	e.Logger.Fatal(e.Start(addr))
}
