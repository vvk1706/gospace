# API Documentation

This document describes all available API endpoints in GoSpace.

## Endpoints

### Home Page
- **URL**: `/`
- **Method**: `GET`
- **Description**: Displays the home page with links to all features
- **Handler**: [`handlers.Home`](handlers/home.go:10)
- **Template**: [`templates/home.html`](templates/home.html)
- **Response**: HTML page

### Calculator

#### Get Calculator Page
- **URL**: `/calculator`
- **Method**: `GET`
- **Description**: Displays the calculator form
- **Handler**: [`handlers.Calculator`](handlers/calculator.go:11)
- **Template**: [`templates/calculator.html`](templates/calculator.html)
- **Response**: HTML page

#### Perform Calculation
- **URL**: `/calculator`
- **Method**: `POST`
- **Content-Type**: `application/x-www-form-urlencoded`
- **Handler**: [`handlers.CalculateResult`](handlers/calculator.go:18)
- **Parameters**:
  - `num1` (required): First number (float64)
  - `num2` (required): Second number (float64)
  - `operation` (required): Operation type (`add`, `subtract`, `multiply`, `divide`)
- **Response**: Redirect to `/calculator/history` after saving calculation
- **Error Responses**:
  - `400 Bad Request`: Invalid numbers or operation
  - `400 Bad Request`: Division by zero (when operation is `divide` and num2 is 0)
  - `500 Internal Server Error`: Database error

#### Get Calculator History
- **URL**: `/calculator/history`
- **Method**: `GET`
- **Description**: Displays all calculation history from the database
- **Handler**: [`handlers.ListCalculatorHistory`](handlers/calculator.go:60)
- **Template**: [`templates/calculator_history.html`](templates/calculator_history.html)
- **Response**: HTML page with history table
- **Error Responses**:
  - `500 Internal Server Error`: Database error

#### Delete Calculation
- **URL**: `/calculator/history/:id/delete`
- **Method**: `POST`
- **Description**: Deletes a calculation from history
- **Handler**: [`handlers.DeleteCalculatorHistory`](handlers/calculator.go:80)
- **Parameters**:
  - `id` (required): Calculation ID (uint, URL parameter)
- **Response**: Redirect to `/calculator/history`
- **Error Responses**:
  - `400 Bad Request`: Invalid ID format
  - `404 Not Found`: Calculation not found
  - `500 Internal Server Error`: Database error

### Contact Form

#### Get Contact Form
- **URL**: `/contact`
- **Method**: `GET`
- **Description**: Displays the contact form
- **Handler**: [`handlers.ContactForm`](handlers/contact.go:11)
- **Template**: [`templates/contact.html`](templates/contact.html)
- **Response**: HTML page

#### Submit Contact
- **URL**: `/contact`
- **Method**: `POST`
- **Content-Type**: `application/x-www-form-urlencoded`
- **Handler**: [`handlers.SubmitContact`](handlers/contact.go:18)
- **Parameters**:
  - `name` (required): First name (string)
  - `surname` (required): Last name (string)
  - `email` (required): Email address (string, must be unique)
- **Response**: HTML page with success message
- **Error Responses**:
  - `400 Bad Request`: Missing required fields
  - `500 Internal Server Error`: Database error (e.g., duplicate email)

#### List All Contacts
- **URL**: `/contacts`
- **Method**: `GET`
- **Description**: Displays all contacts from the PostgreSQL database
- **Handler**: [`handlers.ListContacts`](handlers/contact.go:50)
- **Template**: [`templates/contacts_list.html`](templates/contacts_list.html)
- **Response**: HTML page with contacts table
- **Error Responses**:
  - `500 Internal Server Error`: Database error

## Data Models

### Contact
Defined in [`models/contact.go`](models/contact.go)

```go
type Contact struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Surname   string    `json:"surname"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

### Calculator History
Defined in [`models/calculator_history.go`](models/calculator_history.go)

```go
type CalculatorHistory struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Num1      float64   `json:"num1"`
    Num2      float64   `json:"num2"`
    Operation string    `json:"operation"`
    Result    float64   `json:"result"`
    Version   int       `json:"version"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

Example JSON representation:
```json
{
  "id": 1,
  "num1": 10.0,
  "num2": 5.0,
  "operation": "add",
  "result": 15.0,
  "version": 1,
  "created_at": "2026-04-20T14:00:00Z",
  "updated_at": "2026-04-20T14:00:00Z"
}
```

Example JSON representation for Contact:
```json
{
  "id": 1,
  "name": "John",
  "surname": "Doe",
  "email": "john.doe@example.com",
  "created_at": "2026-04-17T09:00:00Z",
  "updated_at": "2026-04-17T09:00:00Z"
}
```

## Error Handling

All endpoints return appropriate HTTP status codes:
- `200 OK`: Successful request
- `400 Bad Request`: Invalid input or validation error
- `500 Internal Server Error`: Server or database error

Error messages are displayed in the HTML response with appropriate styling.

## Database

The application uses PostgreSQL with GORM ORM ([`config/database.go`](config/database.go)):
- Persistent data storage
- Automatic schema migration
- GORM model relationships
- Optimistic locking with version field for calculator history
- Email uniqueness validation for contacts
- Thread-safe operations

Database configuration is loaded from environment variables (see [`.env.example`](.env.example)).

## Static Assets

- **CSS**: [`static/css/style.css`](static/css/style.css) - All styling
- **No JavaScript**: The application uses pure HTML/CSS with server-side rendering for maximum compatibility and security

## Security Features

- **Input Validation**: All user inputs are validated server-side
- **SQL Injection Prevention**: GORM parameterized queries protect against SQL injection
- **Operation Whitelisting**: Only valid operations (add, subtract, multiply, divide) are allowed
- **ID Validation**: Calculator history IDs are validated before database operations
- **Optimistic Locking**: Version field prevents race conditions in concurrent updates
- **No Client-Side JavaScript**: Eliminates XSS attack vectors