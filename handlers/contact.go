package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user/gospace/models"
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

	// Validate required fields
	if name == "" || surname == "" || email == "" {
		c.HTML(http.StatusBadRequest, "contact.html", gin.H{
			"title": "Contact Form",
			"error": "Please fill in all fields correctly",
		})
		return
	}

	contact := models.NewContact(name, surname, email)

	// Save contact to database
	if err := h.DB.CreateContact(contact); err != nil {
		c.HTML(http.StatusInternalServerError, "contact.html", gin.H{
			"title": "Contact Form",
			"error": "Failed to save contact. Email might already exist.",
		})
		return
	}

	c.HTML(http.StatusOK, "contact.html", gin.H{
		"title":   "Contact Form",
		"success": "Contact saved successfully!",
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
