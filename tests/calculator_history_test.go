package tests

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
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

func TestEditCalculatorHistory_Success(t *testing.T) {
	db := setupTestDB(t)
	h := handlers.NewHandler(db)

	// Create test data
	history := models.NewCalculatorHistory(5, 3, "add", 8)
	db.Create(history)

	// Setup router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.LoadHTMLGlob("../templates/*")
	router.POST("/calculator/history/:id/edit", h.EditCalculatorHistory)

	// Prepare form data
	form := url.Values{}
	form.Add("num1", "10")
	form.Add("num2", "5")
	form.Add("operation", "multiply")

	// Make request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/calculator/history/1/edit", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusSeeOther, w.Code)
	assert.Equal(t, "/calculator/history", w.Header().Get("Location"))

	// Verify update
	var updated models.CalculatorHistory
	db.First(&updated, 1)
	assert.Equal(t, 10.0, updated.Num1)
	assert.Equal(t, 5.0, updated.Num2)
	assert.Equal(t, "multiply", updated.Operation)
	assert.Equal(t, 50.0, updated.Result)
}

func TestEditCalculatorHistory_InvalidID(t *testing.T) {
	db := setupTestDB(t)
	h := handlers.NewHandler(db)

	// Setup router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.LoadHTMLGlob("../templates/*")
	router.POST("/calculator/history/:id/edit", h.EditCalculatorHistory)

	// Prepare form data
	form := url.Values{}
	form.Add("num1", "10")
	form.Add("num2", "5")
	form.Add("operation", "add")

	// Make request with invalid ID
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/calculator/history/invalid/edit", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid ID provided")
}

func TestEditCalculatorHistory_InvalidNumbers(t *testing.T) {
	db := setupTestDB(t)
	h := handlers.NewHandler(db)

	// Create test data
	history := models.NewCalculatorHistory(5, 3, "add", 8)
	db.Create(history)

	// Setup router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.LoadHTMLGlob("../templates/*")
	router.POST("/calculator/history/:id/edit", h.EditCalculatorHistory)

	// Prepare form data with invalid numbers
	form := url.Values{}
	form.Add("num1", "invalid")
	form.Add("num2", "5")
	form.Add("operation", "add")

	// Make request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/calculator/history/1/edit", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid numbers provided")
}

func TestEditCalculatorHistory_InvalidOperation(t *testing.T) {
	db := setupTestDB(t)
	h := handlers.NewHandler(db)

	// Create test data
	history := models.NewCalculatorHistory(5, 3, "add", 8)
	db.Create(history)

	// Setup router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.LoadHTMLGlob("../templates/*")
	router.POST("/calculator/history/:id/edit", h.EditCalculatorHistory)

	// Prepare form data with invalid operation
	form := url.Values{}
	form.Add("num1", "10")
	form.Add("num2", "5")
	form.Add("operation", "invalid")

	// Make request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/calculator/history/1/edit", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid operation")
}

func TestEditCalculatorHistory_DivideByZero(t *testing.T) {
	db := setupTestDB(t)
	h := handlers.NewHandler(db)

	// Create test data
	history := models.NewCalculatorHistory(5, 3, "add", 8)
	db.Create(history)

	// Setup router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.LoadHTMLGlob("../templates/*")
	router.POST("/calculator/history/:id/edit", h.EditCalculatorHistory)

	// Prepare form data with divide by zero
	form := url.Values{}
	form.Add("num1", "10")
	form.Add("num2", "0")
	form.Add("operation", "divide")

	// Make request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/calculator/history/1/edit", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "cannot divide by zero")
}

func TestEditCalculatorHistory_NonExistent(t *testing.T) {
	db := setupTestDB(t)
	h := handlers.NewHandler(db)

	// Setup router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.LoadHTMLGlob("../templates/*")
	router.POST("/calculator/history/:id/edit", h.EditCalculatorHistory)

	// Prepare form data
	form := url.Values{}
	form.Add("num1", "10")
	form.Add("num2", "5")
	form.Add("operation", "add")

	// Make request with non-existent ID
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/calculator/history/999/edit", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Record not found")
}

func TestPerformCalculation_AllOperations(t *testing.T) {
	tests := []struct {
		name      string
		num1      float64
		num2      float64
		operation string
		expected  float64
		expectErr bool
	}{
		{"Add", 5, 3, "add", 8, false},
		{"Subtract", 10, 4, "subtract", 6, false},
		{"Multiply", 6, 7, "multiply", 42, false},
		{"Divide", 20, 4, "divide", 5, false},
		{"Divide by zero", 10, 0, "divide", 0, true},
		{"Invalid operation", 5, 3, "invalid", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := setupTestDB(t)
			h := handlers.NewHandler(db)

			// Setup router
			gin.SetMode(gin.TestMode)
			router := gin.New()
			router.LoadHTMLGlob("../templates/*")
			router.POST("/calculator", h.CalculateResult)

			// Prepare form data
			form := url.Values{}
			form.Add("num1", string(rune(int(tt.num1))))
			form.Add("num2", string(rune(int(tt.num2))))
			form.Add("operation", tt.operation)

			// Make request
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/calculator", strings.NewReader(form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			router.ServeHTTP(w, req)

			if tt.expectErr {
				assert.Equal(t, http.StatusBadRequest, w.Code)
			}
		})
	}
}

// Made with Bob
