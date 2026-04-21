package tests

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/user/gospace/models"
)

func TestListCalculatorHistory_Success(t *testing.T) {
	db := setupTestDB(t)

	// Create test data
	history1 := models.NewCalculatorHistory(5, 3, "add", 8)
	history2 := models.NewCalculatorHistory(10, 2, "divide", 5)
	db.Create(history1)
	db.Create(history2)

	// Setup router
	router := setupRouter(db)

	// Make request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/calculator/history", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Calculator History")
}

func TestListCalculatorHistory_Empty(t *testing.T) {
	db := setupTestDB(t)

	// Setup router
	router := setupRouter(db)

	// Make request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/calculator/history", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "No calculation history found")
}

func TestDeleteCalculatorHistory_Success(t *testing.T) {
	db := setupTestDB(t)

	// Create test data
	history := models.NewCalculatorHistory(5, 3, "add", 8)
	db.Create(history)

	// Setup router
	router := setupRouter(db)

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

	// Setup router
	router := setupRouter(db)

	// Make request with invalid ID
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/calculator/history/invalid/delete", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid ID provided")
}

func TestDeleteCalculatorHistory_NonExistent(t *testing.T) {
	db := setupTestDB(t)

	// Setup router
	router := setupRouter(db)

	// Make request with non-existent ID
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/calculator/history/999/delete", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Record not found")
}

func TestCalculateResult_RedirectsToHistory(t *testing.T) {
	db := setupTestDB(t)

	// Setup router
	router := setupRouter(db)

	// Prepare form data
	form := url.Values{}
	form.Add("num1", "10")
	form.Add("num2", "5")
	form.Add("operation", "add")

	// Make request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/calculator", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	// Verify redirect
	assert.Equal(t, http.StatusSeeOther, w.Code)
	assert.Equal(t, "/calculator/history", w.Header().Get("Location"))

	// Verify calculation was saved
	var count int64
	db.Model(&models.CalculatorHistory{}).Count(&count)
	assert.Equal(t, int64(1), count)
}

func TestCalculateResult_SavesCorrectResult(t *testing.T) {
	db := setupTestDB(t)

	// Setup router
	router := setupRouter(db)

	// Test different operations
	tests := []struct {
		name      string
		num1      string
		num2      string
		operation string
		expected  float64
	}{
		{"Addition", "10", "5", "add", 15},
		{"Subtraction", "10", "5", "subtract", 5},
		{"Multiplication", "10", "5", "multiply", 50},
		{"Division", "10", "5", "divide", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clear database
			db.Exec("DELETE FROM calculator_histories")

			// Prepare form data
			form := url.Values{}
			form.Add("num1", tt.num1)
			form.Add("num2", tt.num2)
			form.Add("operation", tt.operation)

			// Make request
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/calculator", strings.NewReader(form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			router.ServeHTTP(w, req)

			// Verify redirect
			assert.Equal(t, http.StatusSeeOther, w.Code)

			// Verify saved result
			var history models.CalculatorHistory
			db.First(&history)
			assert.Equal(t, tt.expected, history.Result)
			assert.Equal(t, tt.operation, history.Operation)
		})
	}
}

// Made with Bob
