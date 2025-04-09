package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"uniqueIndex;size:255" json:"username"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Email        string    `gorm:"uniqueIndex;size:255" json:"email"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}
