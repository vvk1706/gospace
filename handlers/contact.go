package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/user/gin-webapp/models"
	"github.com/user/gin-webapp/validation"
)

// ContactForm handles the contact form page request
func (h *Handler) ContactForm(c *gin.Context) {
	c.HTML(http.StatusOK, "contact.html", gin.H{
		"title": "Contact Form",
	})
}

// SubmitContact handles the contact form submission
func (h *Handler) SubmitContact(c *gin.Context) {
	name := c.PostForm("name")
	surname := c.PostForm("surname")
	email := c.PostForm("email")

	// Trim whitespace
	name = strings.TrimSpace(name)
	surname = strings.TrimSpace(surname)
	email = strings.TrimSpace(email)

	// Truncate inputs to 100 characters
	name = validation.TruncateString(name, 100)
	surname = validation.TruncateString(surname, 100)
	email = validation.TruncateString(email, 100)

	// Validate required fields
	if name == "" || surname == "" || email == "" {
		c.HTML(http.StatusBadRequest, "contact.html", gin.H{
			"title":   "Contact Form",
			"error":   "All fields are required. Please fill in your name, surname, and email address.",
			"name":    name,
			"surname": surname,
			"email":   email,
		})
		return
	}

	// Validate email format
	if !validation.IsValidEmail(email) {
		c.HTML(http.StatusBadRequest, "contact.html", gin.H{
			"title":   "Contact Form",
			"error":   "Invalid email format. Please enter a valid email address (e.g., user@example.com).",
			"name":    name,
			"surname": surname,
			"email":   email,
		})
		return
	}

	// Auto-capitalize names
	name = validation.CapitalizeName(name)
	surname = validation.CapitalizeName(surname)

	contact := models.NewContact(name, surname, email)

	// Save contact to database
	if err := h.DB.CreateContact(contact); err != nil {
		c.HTML(http.StatusInternalServerError, "contact.html", gin.H{
			"title":   "Contact Form",
			"error":   "Failed to save contact. This email address may already be registered.",
			"name":    name,
			"surname": surname,
			"email":   email,
		})
		return
	}

	c.HTML(http.StatusOK, "contact.html", gin.H{
		"title":   "Contact Form",
		"success": "Contact saved successfully! Thank you for submitting your information.",
	})
}

// ListContacts handles listing all contacts
func (h *Handler) ListContacts(c *gin.Context) {
	contacts, err := h.DB.GetAllContacts()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "contacts_list.html", gin.H{
			"title": "Contacts List",
			"error": "Failed to retrieve contacts",
		})
		return
	}

	c.HTML(http.StatusOK, "contacts_list.html", gin.H{
		"title":    "Contacts List",
		"contacts": contacts,
	})
}

// Made with Bob
