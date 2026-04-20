package handlers

import (
	"gorm.io/gorm"
)

// Handler holds dependencies for HTTP handlers
type Handler struct {
	DB *gorm.DB
}

// NewHandler creates a new Handler instance
func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		DB: db,
	}
}
