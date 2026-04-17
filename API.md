# API Documentation

## Endpoints

### Home Page
- **URL**: `/`
- **Method**: `GET`
- **Description**: Displays the home page with links to all features
- **Response**: HTML page

### Calculator

#### Get Calculator Page
- **URL**: `/calculator`
- **Method**: `GET`
- **Description**: Displays the calculator form
- **Response**: HTML page

#### Perform Calculation
- **URL**: `/calculator`
- **Method**: `POST`
- **Content-Type**: `application/x-www-form-urlencoded`
- **Parameters**:
  - `num1` (required): First number
  - `num2` (required): Second number
  - `operation` (required): Operation type (`add`, `subtract`, `multiply`, `divide`)
- **Response**: HTML page with calculation result
- **Error Responses**:
  - `400 Bad Request`: Invalid numbers or operation
  - `400 Bad Request`: Division by zero

### Contact Form

#### Get Contact Form
- **URL**: `/contact`
- **Method**: `GET`
- **Description**: Displays the contact form
- **Response**: HTML page

#### Submit Contact
- **URL**: `/contact`
- **Method**: `POST`
- **Content-Type**: `application/x-www-form-urlencoded`
- **Parameters**:
  - `name` (required): First name
  - `surname` (required): Last name
  - `email` (required): Email address (must be unique)
- **Response**: HTML page with success message
- **Error Responses**:
  - `400 Bad Request`: Missing required fields
  - `500 Internal Server Error`: Database error (e.g., duplicate email)

#### List All Contacts
- **URL**: `/contacts`
- **Method**: `GET`
- **Description**: Displays all contacts from the database
- **Response**: HTML page with contacts table
- **Error Responses**:
  - `500 Internal Server Error`: Database error

## Data Models

### Contact
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