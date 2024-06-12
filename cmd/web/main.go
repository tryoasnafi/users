package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/tryoasnafi/users/common"
	"github.com/tryoasnafi/users/database"
	"github.com/tryoasnafi/users/internal/users"

	httpSwagger "github.com/swaggo/http-swagger/v2"
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

	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Use(common.ApiVersionCtx("v1"))
		r.Mount("/users", users.Router(db))
	})
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(
		fmt.Sprintf("http://localhost%s/docs/doc.json", addr),
	)))

	// starting the server
	log.Println("Starting user service on port", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Failed to start account service: %v", err)
	}
}
