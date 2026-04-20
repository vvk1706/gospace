package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"github.com/user/gospace/models"
)

// ContactForm handles the contact form page request
func (h *Handler) ContactForm(c *gin.Context) {
	c.HTML(http.StatusOK, "contact.html", gin.H{
		"title": "Contact Form",
		"csrf":  csrf.GetToken(c),
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
			"csrf":  csrf.GetToken(c),
		})
		return
	}

	contact := models.NewContact(name, surname, email)

	// Save contact to database using GORM
	if err := h.DB.Create(contact).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "contact.html", gin.H{
			"title": "Contact Form",
			"error": "Failed to save contact. Email might already exist.",
			"csrf":  csrf.GetToken(c),
		})
		return
	}

	c.HTML(http.StatusOK, "contact.html", gin.H{
		"title":   "Contact Form",
		"success": "Contact saved successfully!",
		"csrf":    csrf.GetToken(c),
	})
}

// ListContacts handles listing all contacts
func (h *Handler) ListContacts(c *gin.Context) {
	var contacts []models.Contact
	if err := h.DB.Find(&contacts).Error; err != nil {
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
