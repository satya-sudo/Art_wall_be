package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/satya-sudo/Art_wall_be.git/handlers"
	"github.com/satya-sudo/Art_wall_be.git/middleware"
)

func RegisterArtistRoutes(app *fiber.App) {
	app.Get("/artist", handlers.GetArtistInfo)
	app.Put("/artist", middleware.RequiresAuth, handlers.UpdateArtistInfo)
}
