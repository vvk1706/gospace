package validation

import (
	"regexp"
	"strings"
	"unicode"
)

// IsValidEmail validates email format using regex
func IsValidEmail(email string) bool {
	if email == "" {
		return false
	}
	// RFC 5322 compliant email regex (simplified)
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// CapitalizeName capitalizes the first letter of each word in a name
func CapitalizeName(name string) string {
	if name == "" {
		return ""
	}

	words := strings.Fields(name)
	for i, word := range words {
		if len(word) > 0 {
			runes := []rune(word)
			runes[0] = unicode.ToUpper(runes[0])
			for j := 1; j < len(runes); j++ {
				runes[j] = unicode.ToLower(runes[j])
			}
			words[i] = string(runes)
		}
	}

	return strings.Join(words, " ")
}

// IsValidNumber checks if a string is a valid number (integer or float)
func IsValidNumber(s string) bool {
	if s == "" {
		return false
	}

	// Allow optional minus sign, digits, and optional decimal point with digits
	numberRegex := regexp.MustCompile(`^-?\d+(\.\d+)?$`)
	return numberRegex.MatchString(s)
}

// SanitizeNumericInput removes non-numeric characters except minus and decimal point
func SanitizeNumericInput(s string) string {
	if s == "" {
		return ""
	}

	// Keep only digits, minus sign (at start), and decimal point
	var result strings.Builder
	hasDecimal := false

	for i, char := range s {
		if unicode.IsDigit(char) {
			result.WriteRune(char)
		} else if char == '-' && i == 0 {
			result.WriteRune(char)
		} else if char == '.' && !hasDecimal {
			result.WriteRune(char)
			hasDecimal = true
		}
	}

	return result.String()
}

// TruncateString truncates a string to the specified maximum length
func TruncateString(s string, maxLength int) string {
	if len(s) <= maxLength {
		return s
	}
	return s[:maxLength]
}

// Made with Bob
