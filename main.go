package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/user/gospace/config"
	"github.com/user/gospace/handlers"
)

func main() {
	// Initialize mock database (no real PostgreSQL needed)
	db := config.NewMockDB()
	log.Println("Using in-memory mock database (no PostgreSQL required)")

	// Initialize Gin router
	router := gin.Default()

	// Load HTML templates
	router.LoadHTMLGlob("templates/*")

	// Serve static files
	router.Static("/static/css", "./static/css")
	router.Static("/static/js", "./static/js")

	// Initialize handlers
	h := handlers.NewHandler(db)

	// Routes
	router.GET("/", h.Home)
	router.GET("/calculator", h.Calculator)
	router.POST("/calculator", h.CalculateResult)
	router.GET("/contact", h.ContactForm)
	router.POST("/contact", h.SubmitContact)
	router.GET("/contacts", h.ListContacts)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s...", port)
	log.Println("No database setup required - using in-memory storage")
	log.Printf("Access the application at http://localhost:%s", port)
	
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
