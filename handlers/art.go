package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/satya-sudo/Art_wall_be.git/config"
	"github.com/satya-sudo/Art_wall_be.git/models"
	"strings"
)

type InputStruct struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ImageURL    string   `json:"image_url"`
	Tags        []string `json:"tags"`
}

func GetAllArtPosts(c *fiber.Ctx) error {
	var posts []models.ArtPost
	err := config.DB.Preload("Tags").Find(&posts).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to get all posts",
		})
	}
	return c.JSON(posts)
}

func GetArtPost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.ArtPost
	err := config.DB.Preload("Tags").First(&post, id).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{})
	}
	return c.JSON(post)
}

func GetArtPostsByTag(c *fiber.Ctx) error {
	tagNames := strings.Split(c.Query("tags"), ",")

	var posts []models.ArtPost
	err := config.DB.
		Joins("JOIN art_post_tags ON art_post_tags.art_post_id = art_posts.id").
		Joins("JOIN tags ON tags.id = art_post_tags.tag_id").
		Where("tags.name IN ?", tagNames).
		Preload("Tags").
		Group("art_posts.id").
		Find(&posts).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch posts",
		})
	}

	return c.JSON(posts)
}

func CreateArtPost(c *fiber.Ctx) error {
	var input InputStruct
	if err := c.BodyParser(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to parse body",
		})
	}
	var tags []models.Tag
	for _, name := range input.Tags {
		var tag models.Tag
		config.DB.FirstOrCreate(&tag, models.Tag{Name: name})
		tags = append(tags, tag)
	}
	post := models.ArtPost{
		Title:       input.Title,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		Tags:        tags,
	}
	if err := config.DB.Create(&post).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create post",
		})
	}
	return c.JSON(post)
}

func UpdateArtPost(c *fiber.Ctx) error {
	id := c.Params("id")

	var post models.ArtPost
	if err := config.DB.First(&post, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Not found",
		})
	}

	var input InputStruct
	if err := c.BodyParser(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to parse body",
		})
	}
	var tags []models.Tag
	for _, name := range input.Tags {
		var tag models.Tag
		config.DB.FirstOrCreate(&tag, models.Tag{Name: name})
		tags = append(tags, tag)
	}
	post.Title = input.Title
	post.Description = input.Description
	post.ImageURL = input.ImageURL
	post.Tags = tags
	config.DB.Model(&post).Association("Tags").Replace(tags)

	if err := config.DB.Save(&post).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update post",
		})
	}
	return c.JSON(post)
}

func DeleteArtPost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.ArtPost
	if err := config.DB.First(&post, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to get post",
		})
	}
	if err := config.DB.Delete(&post).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete post",
		})
	}
	return c.JSON(fiber.Map{
		"success": "ArtPost deleted",
	})
}
