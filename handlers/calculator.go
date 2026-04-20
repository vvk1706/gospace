package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user/gospace/models"
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

	num1, err1 := strconv.ParseFloat(num1Str, 64)
	num2, err2 := strconv.ParseFloat(num2Str, 64)

	if err1 != nil || err2 != nil {
		c.HTML(http.StatusBadRequest, "calculator.html", gin.H{
			"title": "Calculator",
			"error": "Invalid numbers provided",
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
				"title": "Calculator",
				"error": "Cannot divide by zero",
			})
			return
		}
		result = num1 / num2
		resultStr = strconv.FormatFloat(result, 'f', -1, 64)
	default:
		c.HTML(http.StatusBadRequest, "calculator.html", gin.H{
			"title": "Calculator",
			"error": "Invalid operation",
		})
		return
	}

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
	id := c.Param("id")
	
	if err := h.DB.Delete(&models.CalculatorHistory{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete history"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/calculator/history")
}

// EditCalculatorHistory handles updating a calculator history record
func (h *Handler) EditCalculatorHistory(c *gin.Context) {
	id := c.Param("id")
	
	num1Str := c.PostForm("num1")
	num2Str := c.PostForm("num2")
	operation := c.PostForm("operation")

	num1, err1 := strconv.ParseFloat(num1Str, 64)
	num2, err2 := strconv.ParseFloat(num2Str, 64)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid numbers provided"})
		return
	}

	var result float64
	switch operation {
	case "add":
		result = num1 + num2
	case "subtract":
		result = num1 - num2
	case "multiply":
		result = num1 * num2
	case "divide":
		if num2 == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot divide by zero"})
			return
		}
		result = num1 / num2
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid operation"})
		return
	}

	// Update the history record
	if err := h.DB.Model(&models.CalculatorHistory{}).Where("id = ?", id).Updates(map[string]interface{}{
		"num1":      num1,
		"num2":      num2,
		"operation": operation,
		"result":    result,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update history"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/calculator/history")
}
