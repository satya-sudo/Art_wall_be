package models

import "gorm.io/gorm"

type ArtistInfo struct {
	gorm.Model
	Name            string            `json:"name"`
	Bio             string            `json:"bio"` // markdown content
	Email           string            `json:"email"`
	SocialLinks     map[string]string `gorm:"type:jsonb" json:"social_links"` // e.g., {"instagram": "...", "twitter": "..."}
	ProfileImageURL string            `json:"profile_image_url"`
}
