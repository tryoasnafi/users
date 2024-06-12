package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/tryoasnafi/users/database"
	"github.com/tryoasnafi/users/database/migration"
)

func init() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file", err)
	}
}

func main()  {
	db, err := database.GetDB()
	if err != nil {
		log.Fatal("failed connect to database", err)
	}
	if err := migration.Migrate(db); err != nil {
		log.Fatal("failed to migrate database", err)
	}
	log.Println("database migrated successfully")
}