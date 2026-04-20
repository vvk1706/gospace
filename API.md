# API Documentation

This document describes all available API endpoints in the Gin Web Application.

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
- **Response**: HTML page with calculation result
- **Error Responses**:
  - `400 Bad Request`: Invalid numbers or operation
  - `400 Bad Request`: Division by zero (when operation is `divide` and num2 is 0)

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
- **Description**: Displays all contacts from the in-memory database
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

Example JSON representation:
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

The application uses an in-memory database implementation ([`config/mock_database.go`](config/mock_database.go)) by default:
- Thread-safe operations with mutex protection
- Automatic ID generation
- Email uniqueness validation
- No external database required

To use PostgreSQL instead, modify [`main.go`](main.go:14) to use `config.InitDB()` instead of `config.NewMockDB()`.

## Static Assets

- **CSS**: [`static/css/style.css`](static/css/style.css)
- **JavaScript**:
  - [`static/js/main.js`](static/js/main.js) - Common functionality
  - [`static/js/calculator.js`](static/js/calculator.js) - Calculator-specific
  - [`static/js/contact.js`](static/js/contact.js) - Contact form-specific