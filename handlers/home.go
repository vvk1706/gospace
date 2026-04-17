package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home handles the home page request
func (h *Handler) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Hello, World!",
	})
}

// Made with Bob
