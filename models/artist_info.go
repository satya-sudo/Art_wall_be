package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

type JSONMap map[string]string

func (j JSONMap) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONMap) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Failed to unmarshal JSONMap value: %v", value)
	}
	return json.Unmarshal(bytes, j)
}

type ArtistInfo struct {
	gorm.Model
	Name            string  `json:"name"`
	Bio             string  `json:"bio"` // markdown content
	Email           string  `json:"email"`
	SocialLinks     JSONMap `gorm:"type:jsonb" json:"social_links"`
	ProfileImageURL string  `json:"profile_image_url"`
}
