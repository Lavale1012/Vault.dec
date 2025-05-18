package config

import (
	"fmt"
	"os"
	"vault-dev/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Connect to the database
	// This is a placeholder function. You should implement the actual database connection logic here.
	// For example, you can use GORM or any other ORM library to connect to your database.
	err := godotenv.Load()
	if err != nil {
		panic("Error loading DBConnection")
	}
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	// Use the connection string to connect to the database
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	fmt.Println("Connected to the database")
	// Migrate the schema
	db.AutoMigrate(&models.SnippetModel{}, &models.UserModel{})
	DB = db

}
