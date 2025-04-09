package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/satya-sudo/Art_wall_be.git/handlers"
	"github.com/satya-sudo/Art_wall_be.git/middleware"
)

func RegisterUploadRoutes(app *fiber.App) {
	app.Post("/upload", middleware.RequiresAuth, handlers.UploadImage)
}
