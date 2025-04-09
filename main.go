package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/satya-sudo/Art_wall_be.git/config"
	"github.com/satya-sudo/Art_wall_be.git/routes"
	"github.com/satya-sudo/Art_wall_be.git/seed"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	config.ConnectDB()
	seed.SeedMasterUser()

	routes.RegisterRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Art Gallery Backend is running! 🎨")
	})

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
}
