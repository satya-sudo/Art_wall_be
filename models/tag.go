package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name     string    `gorm:"unique" json:"name"`
	ArtPosts []ArtPost `gorm:"many2many:art_post_tags;"`
}
