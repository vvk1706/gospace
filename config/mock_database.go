package config

import (
	"sync"

	"github.com/user/gin-webapp/models"
)

// MockDB is an in-memory database implementation
type MockDB struct {
	contacts map[uint]*models.Contact
	nextID   uint
	mu       sync.RWMutex
}

// NewMockDB creates a new mock database
func NewMockDB() *MockDB {
	return &MockDB{
		contacts: make(map[uint]*models.Contact),
		nextID:   1,
	}
}

// CreateContact adds a new contact to the mock database
func (db *MockDB) CreateContact(contact *models.Contact) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	// Check for duplicate email
	for _, c := range db.contacts {
		if c.Email == contact.Email {
			return &DuplicateEmailError{Email: contact.Email}
		}
	}

	contact.ID = db.nextID
	db.nextID++
	db.contacts[contact.ID] = contact
	return nil
}

// GetAllContacts returns all contacts from the mock database
func (db *MockDB) GetAllContacts() ([]*models.Contact, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	contacts := make([]*models.Contact, 0, len(db.contacts))
	for _, contact := range db.contacts {
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

// GetContactByID retrieves a contact by ID
func (db *MockDB) GetContactByID(id uint) (*models.Contact, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	contact, exists := db.contacts[id]
	if !exists {
		return nil, &NotFoundError{ID: id}
	}
	return contact, nil
}

// UpdateContact updates an existing contact
func (db *MockDB) UpdateContact(contact *models.Contact) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, exists := db.contacts[contact.ID]; !exists {
		return &NotFoundError{ID: contact.ID}
	}

	db.contacts[contact.ID] = contact
	return nil
}

// DeleteContact removes a contact from the mock database
func (db *MockDB) DeleteContact(id uint) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, exists := db.contacts[id]; !exists {
		return &NotFoundError{ID: id}
	}

	delete(db.contacts, id)
	return nil
}

// Custom error types
type DuplicateEmailError struct {
	Email string
}

func (e *DuplicateEmailError) Error() string {
	return "email already exists: " + e.Email
}

type NotFoundError struct {
	ID uint
}

func (e *NotFoundError) Error() string {
	return "contact not found"
}

// Made with Bob
