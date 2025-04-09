package handlers

import (
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
	"github.com/satya-sudo/Art_wall_be.git/utils"
	"time"
)

func UploadImage(c *fiber.Ctx) error {
	fileHander, err := c.FormFile("image")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}
	file, err := fileHander.Open()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}
	defer file.Close()

	cld := utils.InitCloudinary()
	uploadParam := fmt.Sprintf(
		"art_gallary/%d_%s", time.Now().Unix(), fileHander.Filename)
	resp, err := cld.Upload.Upload(
		context.Background(), file, uploader.UploadParams{
			PublicID: uploadParam,
			Folder:   "art_gallary",
		})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"url": resp.SecureURL,
	})

}
