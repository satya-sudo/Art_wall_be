package main

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/satya-sudo/Art_wall_be.git/config"
	"github.com/satya-sudo/Art_wall_be.git/routes"
	"github.com/satya-sudo/Art_wall_be.git/seed"
)

func main() {
	// Load .env file only in development
	_ = godotenv.Load() // Don't fatal — safe fallback
	port := os.Getenv("PORT")
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	config.ConnectDB()
	seed.SeedMasterUser()

	routes.RegisterRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Art Gallery Backend is running! 🎨")
	})

	log.Fatal(app.Listen(":" + port))
}
