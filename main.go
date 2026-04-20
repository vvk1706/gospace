package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/user/gospace/config"
	"github.com/user/gospace/handlers"
	"github.com/user/gospace/models"
)

func main() {
	// Load configuration from environment variables
	cfg := config.LoadConfig()
	
	// Initialize PostgreSQL database
	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to PostgreSQL database")
	
	// Auto-migrate database schema
	if err := db.AutoMigrate(&models.Contact{}, &models.CalculatorHistory{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed")

	// Initialize Gin router
	router := gin.Default()

	// Load HTML templates
	router.LoadHTMLGlob("templates/*")

	// Serve static files
	router.Static("/static/css", "./static/css")

	// Initialize handlers
	h := handlers.NewHandler(db)

	// Routes
	router.GET("/", h.Home)
	router.GET("/calculator", h.Calculator)
	router.POST("/calculator", h.CalculateResult)
	router.GET("/calculator/history", h.ListCalculatorHistory)
	router.POST("/calculator/history/:id/delete", h.DeleteCalculatorHistory)
	router.GET("/contact", h.ContactForm)
	router.POST("/contact", h.SubmitContact)
	router.GET("/contacts", h.ListContacts)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s...", port)
	log.Println("Using PostgreSQL database for persistent storage")
	log.Printf("Access the application at http://localhost:%s", port)
	
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
