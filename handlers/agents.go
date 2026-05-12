package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user/gospace/models"
)

// ListAgents displays all AI agents
func (h *Handler) ListAgents(c *gin.Context) {
	var agents []models.Agent
	
	if err := h.DB.Order("created_at desc").Find(&agents).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "agents.html", gin.H{
			"title": "AI Agents Repository",
			"error": "Failed to fetch agents",
		})
		return
	}

	c.HTML(http.StatusOK, "agents.html", gin.H{
		"title":  "AI Agents Repository",
		"agents": agents,
	})
}

// AddAgentForm displays the form to add a new agent
func (h *Handler) AddAgentForm(c *gin.Context) {
	c.HTML(http.StatusOK, "add_agent.html", gin.H{
		"title": "Add AI Agent",
	})
}

// CreateAgent handles the creation of a new agent
func (h *Handler) CreateAgent(c *gin.Context) {
	var agent models.Agent

	if err := c.ShouldBind(&agent); err != nil {
		c.HTML(http.StatusBadRequest, "add_agent.html", gin.H{
			"title": "Add AI Agent",
			"error": "Invalid form data: " + err.Error(),
		})
		return
	}

	if err := h.DB.Create(&agent).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "add_agent.html", gin.H{
			"title": "Add AI Agent",
			"error": "Failed to create agent: " + err.Error(),
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/agents")
}

// DeleteAgent handles the deletion of an agent
func (h *Handler) DeleteAgent(c *gin.Context) {
	id := c.Param("id")
	agentID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid agent ID"})
		return
	}

	if err := h.DB.Delete(&models.Agent{}, agentID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete agent"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/agents")
}

// Made with Bob
