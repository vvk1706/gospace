package models

import (
	"time"
)

// Contact represents a contact record
type Contact struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewContact creates a new contact with timestamps
func NewContact(name, surname, email string) *Contact {
	now := time.Now()
	return &Contact{
		Name:      name,
		Surname:   surname,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// TableName specifies the table name for Contact model
func (Contact) TableName() string {
	return "contacts"
}

// Made with Bob
