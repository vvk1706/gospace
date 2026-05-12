package models

import (
	"time"

	"gorm.io/gorm"
)

// Tool represents an AI tool in the repository
type Tool struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Name        string         `gorm:"not null;size:255" json:"name" binding:"required"`
	Description string         `gorm:"type:text" json:"description" binding:"required"`
	Category    string         `gorm:"size:100" json:"category" binding:"required"`
	Version     string         `gorm:"size:50" json:"version"`
	Author      string         `gorm:"size:255" json:"author"`
	Repository  string         `gorm:"size:500" json:"repository"`
	Tags        string         `gorm:"type:text" json:"tags"` // Comma-separated tags
	Language    string         `gorm:"size:50" json:"language"` // Programming language
	Status      string         `gorm:"size:50;default:'active'" json:"status"` // active, deprecated, experimental
}

// Made with Bob
