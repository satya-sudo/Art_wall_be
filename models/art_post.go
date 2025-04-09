package models

import "gorm.io/gorm"

type ArtPost struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Tags        []Tag  `gorm:"many2many:art_post_tags;" json:"tags"`
}
