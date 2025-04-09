package seed

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/satya-sudo/Art_wall_be.git/config"
	"github.com/satya-sudo/Art_wall_be.git/models"
)

func SeedMasterUser() {
	password := os.Getenv("ADMIN_PASSWORD")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password")
	}

	user := models.User{
		Username:     "admin",
		PasswordHash: string(hashedPassword),
		Email:        "admin@example.com",
	}

	result := config.DB.FirstOrCreate(&user, models.User{Username: "admin"})
	if result.Error != nil {
		log.Fatal("Failed to seed user:", result.Error)
	}

	fmt.Println("Master user seeded or already exists")
}
