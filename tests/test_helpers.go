package tests

import (
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"github.com/user/gospace/handlers"
	"github.com/user/gospace/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDB creates an in-memory SQLite database for testing
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Auto-migrate the schema
	err = db.AutoMigrate(&models.Contact{}, &models.CalculatorHistory{})
	if err != nil {
		t.Fatalf("Failed to migrate test database: %v", err)
	}

	return db
}

// setupRouter creates a test router with handlers
func setupRouter(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	
	// Setup session store for CSRF protection in tests
	store := cookie.NewStore([]byte("test-secret"))
	router.Use(sessions.Sessions("test_session", store))
	
	// Add CSRF protection middleware (skip validation in tests)
	router.Use(csrf.Middleware(csrf.Options{
		Secret: "test-csrf-secret",
		IgnoreMethods: []string{"GET", "HEAD", "OPTIONS", "POST", "PUT", "PATCH", "DELETE"},
	}))
	
	router.LoadHTMLGlob("../templates/*")

	h := handlers.NewHandler(db)

	router.GET("/", h.Home)
	router.GET("/calculator", h.Calculator)
	router.POST("/calculator", h.CalculateResult)
	router.GET("/calculator/history", h.ListCalculatorHistory)
	router.POST("/calculator/history/:id/delete", h.DeleteCalculatorHistory)
	router.GET("/contact", h.ContactForm)
	router.POST("/contact", h.SubmitContact)
	router.GET("/contacts", h.ListContacts)

	return router
}

// Made with Bob
