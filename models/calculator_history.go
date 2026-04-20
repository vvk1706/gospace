package models

import (
	"time"
)

// CalculatorHistory represents a calculator operation record
type CalculatorHistory struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Num1      float64   `json:"num1"`
	Num2      float64   `json:"num2"`
	Operation string    `json:"operation"`
	Result    float64   `json:"result"`
	Version   int       `gorm:"default:0" json:"version"` // For optimistic locking
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewCalculatorHistory creates a new calculator history record
func NewCalculatorHistory(num1, num2 float64, operation string, result float64) *CalculatorHistory {
	return &CalculatorHistory{
		Num1:      num1,
		Num2:      num2,
		Operation: operation,
		Result:    result,
	}
}

// Made with Bob
