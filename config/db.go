package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"

	"github.com/satya-sudo/Art_wall_be.git/models"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	err = db.AutoMigrate(
		&models.User{},
		&models.Tag{},
		&models.ArtPost{},
		&models.ArtistInfo{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
	DB = db
	log.Println("Connected to database")
}
