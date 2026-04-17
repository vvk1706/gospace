package tests

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/user/gin-webapp/config"
	"github.com/user/gin-webapp/handlers"
	"github.com/user/gin-webapp/models"
)

// setupTestDB creates a mock database for testing
func setupTestDB() *config.MockDB {
	return config.NewMockDB()
}

// setupRouter creates a test router with handlers
func setupRouter(db *config.MockDB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.LoadHTMLGlob("../templates/*")

	h := handlers.NewHandler(db)

	router.GET("/", h.Home)
	router.GET("/calculator", h.Calculator)
	router.POST("/calculator", h.CalculateResult)
	router.GET("/contact", h.ContactForm)
	router.POST("/contact", h.SubmitContact)
	router.GET("/contacts", h.ListContacts)

	return router
}

func TestHomeHandler(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Hello, World!")
}

func TestCalculatorGetHandler(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/calculator", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Calculator")
}

func TestCalculatorPostHandler(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	tests := []struct {
		name           string
		num1           string
		num2           string
		operation      string
		expectedStatus int
		shouldContain  string
	}{
		{
			name:           "Addition",
			num1:           "10",
			num2:           "5",
			operation:      "add",
			expectedStatus: http.StatusOK,
			shouldContain:  "15",
		},
		{
			name:           "Subtraction",
			num1:           "10",
			num2:           "5",
			operation:      "subtract",
			expectedStatus: http.StatusOK,
			shouldContain:  "5",
		},
		{
			name:           "Multiplication",
			num1:           "10",
			num2:           "5",
			operation:      "multiply",
			expectedStatus: http.StatusOK,
			shouldContain:  "50",
		},
		{
			name:           "Division",
			num1:           "10",
			num2:           "5",
			operation:      "divide",
			expectedStatus: http.StatusOK,
			shouldContain:  "2",
		},
		{
			name:           "Division by zero",
			num1:           "10",
			num2:           "0",
			operation:      "divide",
			expectedStatus: http.StatusBadRequest,
			shouldContain:  "Cannot divide by zero",
		},
		{
			name:           "Invalid number",
			num1:           "abc",
			num2:           "5",
			operation:      "add",
			expectedStatus: http.StatusBadRequest,
			shouldContain:  "Invalid numbers",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("num1", tt.num1)
			form.Add("num2", tt.num2)
			form.Add("operation", tt.operation)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/calculator", strings.NewReader(form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.shouldContain)
		})
	}
}

func TestContactFormGetHandler(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/contact", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Contact Form")
}

func TestContactFormPostHandler(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	tests := []struct {
		name           string
		formData       map[string]string
		expectedStatus int
		shouldContain  string
	}{
		{
			name: "Valid contact",
			formData: map[string]string{
				"name":    "John",
				"surname": "Doe",
				"email":   "john.doe@example.com",
			},
			expectedStatus: http.StatusOK,
			shouldContain:  "Contact saved successfully",
		},
		{
			name: "Missing fields",
			formData: map[string]string{
				"name":  "John",
				"email": "john.doe@example.com",
			},
			expectedStatus: http.StatusBadRequest,
			shouldContain:  "fill in all fields",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{}
			for key, value := range tt.formData {
				form.Add(key, value)
			}

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/contact", strings.NewReader(form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.shouldContain)
		})
	}
}

func TestListContactsHandler(t *testing.T) {
	db := setupTestDB()

	// Add test contacts
	db.CreateContact(models.NewContact("John", "Doe", "john@example.com"))
	db.CreateContact(models.NewContact("Jane", "Smith", "jane@example.com"))

	router := setupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/contacts", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John")
	assert.Contains(t, w.Body.String(), "Jane")
	assert.Contains(t, w.Body.String(), "john@example.com")
	assert.Contains(t, w.Body.String(), "jane@example.com")
}

func TestDuplicateEmailContact(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	// First submission
	form := url.Values{}
	form.Add("name", "John")
	form.Add("surname", "Doe")
	form.Add("email", "john@example.com")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/contact", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Second submission with same email
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/contact", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Failed to save contact")
}

// Made with Bob
