package routes

import (
	"github.com/gofiber/fiber/v2"
	
	"github.com/satya-sudo/Art_wall_be.git/handlers"
)

func RegisterAuthRoutes(app *fiber.App) {
	app.Post("/auth", handlers.Login)
}
