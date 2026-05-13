package models

import (
	"time"

	"gorm.io/gorm"
)

// Agent represents an AI agent in the repository
type Agent struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Name        string         `gorm:"not null;size:255" json:"name" form:"name" binding:"required"`
	Description string         `gorm:"type:text" json:"description" form:"description" binding:"required"`
	Category    string         `gorm:"size:100" json:"category" form:"category" binding:"required"`
	Version     string         `gorm:"size:50" json:"version" form:"version"`
	Author      string         `gorm:"size:255" json:"author" form:"author"`
	Repository  string         `gorm:"size:500" json:"repository" form:"repository"`
	Tags        string         `gorm:"type:text" json:"tags" form:"tags"`                    // Comma-separated tags
	Status      string         `gorm:"size:50;default:'active'" json:"status" form:"status"` // active, deprecated, experimental
}

// Made with Bob
