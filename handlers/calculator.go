package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user/gospace/models"
)

// performCalculation performs a calculation based on the operation
func performCalculation(num1, num2 float64, operation string) (float64, error) {
	switch operation {
	case "add":
		return num1 + num2, nil
	case "subtract":
		return num1 - num2, nil
	case "multiply":
		return num1 * num2, nil
	case "divide":
		if num2 == 0 {
			return 0, errors.New("cannot divide by zero")
		}
		return num1 / num2, nil
	default:
		return 0, errors.New("invalid operation")
	}
}

// validateOperation checks if the operation is valid
func validateOperation(operation string) bool {
	validOps := map[string]bool{
		"add":      true,
		"subtract": true,
		"multiply": true,
		"divide":   true,
	}
	return validOps[operation]
}

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

	num1, err1 := strconv.ParseFloat(num1Str, 64)
	num2, err2 := strconv.ParseFloat(num2Str, 64)

	if err1 != nil || err2 != nil {
		c.HTML(http.StatusBadRequest, "calculator.html", gin.H{
			"title": "Calculator",
			"error": "Invalid numbers provided",
		})
		return
	}

	// Validate operation
	if !validateOperation(operation) {
		c.HTML(http.StatusBadRequest, "calculator.html", gin.H{
			"title": "Calculator",
			"error": "Invalid operation",
		})
		return
	}

	// Perform calculation
	result, err := performCalculation(num1, num2, operation)
	if err != nil {
		c.HTML(http.StatusBadRequest, "calculator.html", gin.H{
			"title": "Calculator",
			"error": err.Error(),
		})
		return
	}

	resultStr := strconv.FormatFloat(result, 'f', -1, 64)

	// Save calculation to history
	history := models.NewCalculatorHistory(num1, num2, operation, result)
	if err := h.DB.Create(history).Error; err != nil {
		// Log error but don't fail the request
		c.HTML(http.StatusOK, "calculator.html", gin.H{
			"title":   "Calculator",
			"result":  resultStr,
			"num1":    num1Str,
			"num2":    num2Str,
			"warning": "Calculation successful but failed to save history",
		})
		return
	}

	c.HTML(http.StatusOK, "calculator.html", gin.H{
		"title":  "Calculator",
		"result": resultStr,
		"num1":   num1Str,
		"num2":   num2Str,
	})
}

// ListCalculatorHistory handles listing all calculator history
func (h *Handler) ListCalculatorHistory(c *gin.Context) {
	var history []models.CalculatorHistory
	if err := h.DB.Order("created_at DESC").Find(&history).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "calculator_history.html", gin.H{
			"title": "Calculator History",
			"error": "Failed to retrieve history",
		})
		return
	}

	c.HTML(http.StatusOK, "calculator_history.html", gin.H{
		"title":   "Calculator History",
		"history": history,
	})
}

// DeleteCalculatorHistory handles deleting a calculator history record
func (h *Handler) DeleteCalculatorHistory(c *gin.Context) {
	idStr := c.Param("id")
	
	// Validate ID is a valid integer
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.HTML(http.StatusBadRequest, "calculator_history.html", gin.H{
			"title": "Calculator History",
			"error": "Invalid ID provided",
		})
		return
	}

	// Delete the record and check if it existed
	result := h.DB.Delete(&models.CalculatorHistory{}, id)
	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "calculator_history.html", gin.H{
			"title": "Calculator History",
			"error": "Failed to delete history",
		})
		return
	}

	// Check if any rows were affected
	if result.RowsAffected == 0 {
		c.HTML(http.StatusNotFound, "calculator_history.html", gin.H{
			"title": "Calculator History",
			"error": "Record not found",
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/calculator/history")
}

