package routes

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	RegisterAuthRoutes(app)
	RegisterArtistRoutes(app)
	RegisterArtRoutes(app)
	RegisterUploadRoutes(app)
}
