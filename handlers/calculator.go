package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

	c.HTML(http.StatusOK, "calculator.html", gin.H{
		"title":  "Calculator",
		"result": resultStr,
		"num1":   num1Str,
		"num2":   num2Str,
	})
}
