package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/satya-sudo/Art_wall_be.git/handlers"
	"github.com/satya-sudo/Art_wall_be.git/middleware"
)

func RegisterArtRoutes(app *fiber.App) {
	posts := app.Group("/posts")

	posts.Get("/", handlers.GetAllArtPosts)
	posts.Get("/:id", handlers.GetArtPost)
	posts.Get("/filter/tags", handlers.GetArtPostsByTag)

	// protected routes
	posts.Post("/", middleware.RequiresAuth, handlers.CreateArtPost)
	posts.Delete("/:id", middleware.RequiresAuth, handlers.DeleteArtPost)
	posts.Put("/:id", middleware.RequiresAuth, handlers.UpdateArtPost)

}
