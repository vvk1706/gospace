package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/user/gospace/handlers"
	"github.com/user/gospace/models"
)

func TestListCalculatorHistory_Success(t *testing.T) {
	db := setupTestDB(t)
	h := handlers.NewHandler(db)

	// Create test data
	history1 := models.NewCalculatorHistory(5, 3, "add", 8)
	history2 := models.NewCalculatorHistory(10, 2, "divide", 5)
	db.Create(history1)
	db.Create(history2)

	// Setup router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.LoadHTMLGlob("../templates/*")
	router.GET("/calculator/history", h.ListCalculatorHistory)

	// Make request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/calculator/history", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Calculator History")
}

func TestListCalculatorHistory_Empty(t *testing.T) {
	db := setupTestDB(t)
	h := handlers.NewHandler(db)

	// Setup router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.LoadHTMLGlob("../templates/*")
	router.GET("/calculator/history", h.ListCalculatorHistory)

	// Make request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/calculator/history", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "No calculation history found")
}

func TestDeleteCalculatorHistory_Success(t *testing.T) {
	db := setupTestDB(t)
	h := handlers.NewHandler(db)

	// Create test data
	history := models.NewCalculatorHistory(5, 3, "add", 8)
	db.Create(history)

	// Setup router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.LoadHTMLGlob("../templates/*")
	router.POST("/calculator/history/:id/delete", h.DeleteCalculatorHistory)

	// Make request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/calculator/history/1/delete", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusSeeOther, w.Code)
	assert.Equal(t, "/calculator/history", w.Header().Get("Location"))

	// Verify deletion
	var count int64
	db.Model(&models.CalculatorHistory{}).Count(&count)
	assert.Equal(t, int64(0), count)
}

func TestDeleteCalculatorHistory_InvalidID(t *testing.T) {
	db := setupTestDB(t)
	h := handlers.NewHandler(db)

	// Setup router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.LoadHTMLGlob("../templates/*")
	router.POST("/calculator/history/:id/delete", h.DeleteCalculatorHistory)

	// Make request with invalid ID
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/calculator/history/invalid/delete", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid ID provided")
}

func TestDeleteCalculatorHistory_NonExistent(t *testing.T) {
	db := setupTestDB(t)
	h := handlers.NewHandler(db)

	// Setup router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.LoadHTMLGlob("../templates/*")
	router.POST("/calculator/history/:id/delete", h.DeleteCalculatorHistory)

	// Make request with non-existent ID
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/calculator/history/999/delete", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Record not found")
}

// Made with Bob
