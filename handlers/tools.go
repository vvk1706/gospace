package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user/gospace/models"
)

// ListTools displays all AI tools
func (h *Handler) ListTools(c *gin.Context) {
	var tools []models.Tool
	
	if err := h.DB.Order("created_at desc").Find(&tools).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "tools.html", gin.H{
			"title": "AI Tools Repository",
			"error": "Failed to fetch tools",
		})
		return
	}

	c.HTML(http.StatusOK, "tools.html", gin.H{
		"title": "AI Tools Repository",
		"tools": tools,
	})
}

// AddToolForm displays the form to add a new tool
func (h *Handler) AddToolForm(c *gin.Context) {
	c.HTML(http.StatusOK, "add_tool.html", gin.H{
		"title": "Add AI Tool",
	})
}

// CreateTool handles the creation of a new tool
func (h *Handler) CreateTool(c *gin.Context) {
	var tool models.Tool

	if err := c.ShouldBind(&tool); err != nil {
		c.HTML(http.StatusBadRequest, "add_tool.html", gin.H{
			"title": "Add AI Tool",
			"error": "Invalid form data: " + err.Error(),
		})
		return
	}

	if err := h.DB.Create(&tool).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "add_tool.html", gin.H{
			"title": "Add AI Tool",
			"error": "Failed to create tool: " + err.Error(),
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/tools")
}

// DeleteTool handles the deletion of a tool
func (h *Handler) DeleteTool(c *gin.Context) {
	id := c.Param("id")
	toolID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tool ID"})
		return
	}

	if err := h.DB.Delete(&models.Tool{}, toolID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tool"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/tools")
}

// Made with Bob
