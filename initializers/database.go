package initializers

import (
	"log"
	"os"

	"github.com/abdulbasit108/go-postgresql-basic-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
}

func ApplyMigrations() {
	DB.AutoMigrate(&models.Movies{})
	DB.AutoMigrate(&models.User{})
}
