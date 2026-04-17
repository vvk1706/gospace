package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user/gin-webapp/validation"
)

// Calculator handles the calculator page request
func (h *Handler) Calculator(c *gin.Context) {
	c.HTML(http.StatusOK, "calculator.html", gin.H{
		"title": "Calculator",
	})
}

// CalculateResult handles the calculation request
func (h *Handler) CalculateResult(c *gin.Context) {
	num1Str := c.PostForm("num1")
	num2Str := c.PostForm("num2")
	operation := c.PostForm("operation")

	// Sanitize numeric inputs
	num1Str = validation.SanitizeNumericInput(num1Str)
	num2Str = validation.SanitizeNumericInput(num2Str)

	// Validate inputs are not empty
	if num1Str == "" || num2Str == "" {
		c.HTML(http.StatusBadRequest, "calculator.html", gin.H{
			"title":     "Calculator",
			"error":     "Both numbers are required. Please enter valid numeric values.",
			"num1":      num1Str,
			"num2":      num2Str,
			"operation": operation,
		})
		return
	}

	// Validate numeric format
	if !validation.IsValidNumber(num1Str) {
		c.HTML(http.StatusBadRequest, "calculator.html", gin.H{
			"title":     "Calculator",
			"error":     "First number is invalid. Please enter a valid number (e.g., 42 or 3.14).",
			"num1":      num1Str,
			"num2":      num2Str,
			"operation": operation,
		})
		return
	}

	if !validation.IsValidNumber(num2Str) {
		c.HTML(http.StatusBadRequest, "calculator.html", gin.H{
			"title":     "Calculator",
			"error":     "Second number is invalid. Please enter a valid number (e.g., 42 or 3.14).",
			"num1":      num1Str,
			"num2":      num2Str,
			"operation": operation,
		})
		return
	}

	// Parse numbers
	num1, err1 := strconv.ParseFloat(num1Str, 64)
	num2, err2 := strconv.ParseFloat(num2Str, 64)

	if err1 != nil || err2 != nil {
		c.HTML(http.StatusBadRequest, "calculator.html", gin.H{
			"title":     "Calculator",
			"error":     "Unable to parse numbers. Please ensure you entered valid numeric values.",
			"num1":      num1Str,
			"num2":      num2Str,
			"operation": operation,
		})
		return
	}

	// Validate operation
	validOperations := map[string]bool{
		"add":      true,
		"subtract": true,
		"multiply": true,
		"divide":   true,
	}

	if !validOperations[operation] {
		c.HTML(http.StatusBadRequest, "calculator.html", gin.H{
			"title":     "Calculator",
			"error":     "Invalid operation selected. Please choose Add, Subtract, Multiply, or Divide.",
			"num1":      num1Str,
			"num2":      num2Str,
			"operation": operation,
		})
		return
	}

	var result float64
	var resultStr string

	switch operation {
	case "add":
		result = num1 + num2
		resultStr = strconv.FormatFloat(result, 'f', -1, 64)
	case "subtract":
		result = num1 - num2
		resultStr = strconv.FormatFloat(result, 'f', -1, 64)
	case "multiply":
		result = num1 * num2
		resultStr = strconv.FormatFloat(result, 'f', -1, 64)
	case "divide":
		if num2 == 0 {
			c.HTML(http.StatusBadRequest, "calculator.html", gin.H{
				"title":     "Calculator",
				"error":     "Cannot divide by zero. Please enter a non-zero value for the second number.",
				"num1":      num1Str,
				"num2":      num2Str,
				"operation": operation,
			})
			return
		}
		result = num1 / num2
		resultStr = strconv.FormatFloat(result, 'f', -1, 64)
	}

	c.HTML(http.StatusOK, "calculator.html", gin.H{
		"title":     "Calculator",
		"result":    resultStr,
		"num1":      num1Str,
		"num2":      num2Str,
		"operation": operation,
	})
}

// Made with Bob
