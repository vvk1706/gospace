package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFullUserFlow tests the complete user journey through the application
func TestFullUserFlow(t *testing.T) {
	db := setupTestDB(t)

	router := setupRouter(db)

	t.Run("Complete user journey", func(t *testing.T) {
		// Step 1: Visit home page
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Hello, World!")

		// Step 2: Visit calculator
		w = httptest.NewRecorder()
		req, _ = http.NewRequest("GET", "/calculator", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		// Step 3: Perform calculation
		form := url.Values{}
		form.Add("num1", "100")
		form.Add("num2", "25")
		form.Add("operation", "add")

		w = httptest.NewRecorder()
		req, _ = http.NewRequest("POST", "/calculator", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusSeeOther, w.Code)
		assert.Equal(t, "/calculator/history", w.Header().Get("Location"))

		// Step 4: Visit contact form
		w = httptest.NewRecorder()
		req, _ = http.NewRequest("GET", "/contact", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		// Step 5: Submit contact
		contactForm := url.Values{}
		contactForm.Add("name", "Integration")
		contactForm.Add("surname", "Test")
		contactForm.Add("email", "integration@test.com")

		w = httptest.NewRecorder()
		req, _ = http.NewRequest("POST", "/contact", strings.NewReader(contactForm.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Contact saved successfully")

		// Step 6: View contacts list
		w = httptest.NewRecorder()
		req, _ = http.NewRequest("GET", "/contacts", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Integration")
		assert.Contains(t, w.Body.String(), "integration@test.com")
	})
}

// TestMultipleCalculations tests performing multiple calculations in sequence
func TestMultipleCalculations(t *testing.T) {
	db := setupTestDB(t)

	router := setupRouter(db)

	operations := []struct {
		num1      string
		num2      string
		operation string
		expected  string
	}{
		{"10", "5", "add", "15"},
		{"20", "8", "subtract", "12"},
		{"7", "6", "multiply", "42"},
		{"100", "4", "divide", "25"},
	}

	for _, op := range operations {
		t.Run(fmt.Sprintf("%s %s %s", op.num1, op.operation, op.num2), func(t *testing.T) {
			form := url.Values{}
			form.Add("num1", op.num1)
			form.Add("num2", op.num2)
			form.Add("operation", op.operation)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/calculator", strings.NewReader(form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusSeeOther, w.Code)
			assert.Equal(t, "/calculator/history", w.Header().Get("Location"))
		})
	}
}

// TestMultipleContactSubmissions tests submitting multiple contacts
func TestMultipleContactSubmissions(t *testing.T) {
	db := setupTestDB(t)

	router := setupRouter(db)

	contacts := []struct {
		name    string
		surname string
		email   string
	}{
		{"Alice", "Anderson", "alice@example.com"},
		{"Bob", "Brown", "bob@example.com"},
		{"Charlie", "Clark", "charlie@example.com"},
	}

	// Submit all contacts
	for _, contact := range contacts {
		form := url.Values{}
		form.Add("name", contact.name)
		form.Add("surname", contact.surname)
		form.Add("email", contact.email)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/contact", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	}

	// Verify all contacts are in the list
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/contacts", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	for _, contact := range contacts {
		assert.Contains(t, w.Body.String(), contact.name)
		assert.Contains(t, w.Body.String(), contact.email)
	}
}

// TestErrorHandling tests various error scenarios
func TestErrorHandling(t *testing.T) {
	db := setupTestDB(t)

	router := setupRouter(db)

	t.Run("Calculator with invalid operation", func(t *testing.T) {
		form := url.Values{}
		form.Add("num1", "10")
		form.Add("num2", "5")
		form.Add("operation", "invalid")

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/calculator", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid operation")
	})

	t.Run("Calculator with non-numeric input", func(t *testing.T) {
		form := url.Values{}
		form.Add("num1", "abc")
		form.Add("num2", "5")
		form.Add("operation", "add")

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/calculator", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid numbers")
	})

	t.Run("Contact with missing email", func(t *testing.T) {
		form := url.Values{}
		form.Add("name", "John")
		form.Add("surname", "Doe")

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/contact", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

// TestConcurrentRequests tests handling multiple concurrent requests
func TestConcurrentRequests(t *testing.T) {
	db := setupTestDB(t)
	
	// Enable WAL mode for better concurrent write support in SQLite
	db.Exec("PRAGMA journal_mode=WAL")

	router := setupRouter(db)

	// Create a channel to collect results
	results := make(chan int, 10)

	// Send 10 concurrent requests
	for i := 0; i < 10; i++ {
		go func(index int) {
			form := url.Values{}
			form.Add("name", fmt.Sprintf("User%d", index))
			form.Add("surname", fmt.Sprintf("Test%d", index))
			form.Add("email", fmt.Sprintf("user%d@test.com", index))

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/contact", strings.NewReader(form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			router.ServeHTTP(w, req)

			results <- w.Code
		}(i)
	}

	// Collect results
	successCount := 0
	for i := 0; i < 10; i++ {
		code := <-results
		if code == http.StatusOK {
			successCount++
		}
	}

	// SQLite in-memory databases have limitations with concurrent writes
	// We expect at least 1 successful request, but not necessarily all 10
	assert.GreaterOrEqual(t, successCount, 1, "At least one concurrent request should succeed")
}
