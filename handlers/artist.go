package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/satya-sudo/Art_wall_be.git/config"
	"github.com/satya-sudo/Art_wall_be.git/models"
)

func GetArtistInfo(c *fiber.Ctx) error {
	var artist models.ArtistInfo
	result := config.DB.First(&artist)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "artist info not found",
		})
	}
	return c.JSON(artist)
}

func UpdateArtistInfo(c *fiber.Ctx) error {
	var body, artistInfo models.ArtistInfo
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	config.DB.FirstOrCreate(&artistInfo)
	artistInfo.Name = body.Name
	artistInfo.Bio = body.Bio
	artistInfo.Email = body.Email
	artistInfo.SocialLinks = body.SocialLinks
	artistInfo.ProfileImageURL = body.ProfileImageURL

	config.DB.Save(&artistInfo)
	return c.JSON(artistInfo)
}
