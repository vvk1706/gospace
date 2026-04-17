package handlers

import (
	"github.com/user/gin-webapp/config"
)

// Handler holds dependencies for HTTP handlers
type Handler struct {
	DB *config.MockDB
}

// NewHandler creates a new Handler instance
func NewHandler(db *config.MockDB) *Handler {
	return &Handler{
		DB: db,
	}
}

// Made with Bob
