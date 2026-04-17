# Gin Web Application

A comprehensive web application built with Go and the Gin framework, featuring a calculator, contact form with in-memory database storage, and a modern responsive UI.

## Features

- **Hello World Home Page**: Landing page with navigation to all features
- **Calculator**: Perform basic arithmetic operations (addition, subtraction, multiplication, division)
- **Contact Form**: Submit and store contact information (name, surname, email) in memory
- **Contact List**: View all stored contacts from the in-memory database
- **Responsive Design**: Modern, mobile-friendly UI with smooth animations
- **No Database Required**: Uses in-memory storage - no PostgreSQL installation needed!

## Project Structure

```
gin-webapp/
├── main.go                 # Application entry point
├── config/                 # Configuration files
│   ├── config.go          # Environment configuration
│   ├── database.go        # PostgreSQL connector (optional)
│   └── mock_database.go   # In-memory database implementation
├── models/                 # Data models
│   └── contact.go         # Contact model
├── handlers/               # HTTP request handlers
│   ├── handler.go         # Base handler
│   ├── home.go            # Home page handler
│   ├── calculator.go      # Calculator handlers
│   └── contact.go         # Contact form handlers
├── templates/              # HTML templates
│   ├── home.html          # Home page template
│   ├── calculator.html    # Calculator page template
│   ├── contact.html       # Contact form template
│   └── contacts_list.html # Contacts list template
├── static/                 # Static assets
│   ├── css/
│   │   └── style.css      # Main stylesheet
│   └── js/
│       ├── main.js        # Common JavaScript
│       ├── calculator.js  # Calculator-specific JS
│       └── contact.js     # Contact form-specific JS
├── tests/                  # Test files
│   ├── handlers_test.go   # Handler tests
│   └── integration_test.go # Integration tests
├── go.mod                  # Go module file
├── go.sum                  # Go dependencies
├── .env.example           # Environment variables example
├── .gitignore             # Git ignore file
├── Dockerfile             # Docker configuration
├── docker-compose.yml     # Docker Compose configuration
├── README.md              # This file
├── QUICKSTART.md          # Quick start guide
└── API.md                 # API documentation
```

## Prerequisites

- Go 1.21 or higher
- Git

**Note**: No database installation required! The application uses in-memory storage.

## Installation

### Quick Start (2 minutes)

1. **Clone the Repository**

```bash
git clone <repository-url>
cd gin-webapp
```

2. **Install Dependencies**

```bash
go mod download
```

3. **Run the Application**

```bash
go run main.go
```

4. **Access the Application**

Open your browser and navigate to `http://localhost:8080`

That's it! No database setup required.

## Usage

### Home Page
Navigate to `http://localhost:8080` to see the landing page with links to all features.

### Calculator
1. Go to `http://localhost:8080/calculator`
2. Enter two numbers
3. Select an operation (add, subtract, multiply, divide)
4. Click "Calculate" to see the result

### Contact Form
1. Go to `http://localhost:8080/contact`
2. Fill in your name, surname, and email
3. Click "Submit" to save to the in-memory database
4. View all contacts at `http://localhost:8080/contacts`

**Note**: Data is stored in memory and will be lost when the application restarts.

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Home page |
| GET | `/calculator` | Calculator page |
| POST | `/calculator` | Calculate result |
| GET | `/contact` | Contact form page |
| POST | `/contact` | Submit contact |
| GET | `/contacts` | List all contacts |

## Running Tests

### Run All Tests

```bash
go test ./tests -v
```

### Run with Coverage

```bash
go test ./tests -cover
```

### Generate Coverage Report

```bash
go test ./tests -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Building for Production

```bash
# Build the binary
go build -o gin-webapp

# Run the binary
./gin-webapp
```

## Docker Deployment

### Build Docker Image

```bash
docker build -t gin-webapp .
```

### Run with Docker

```bash
docker run -p 8080:8080 gin-webapp
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| PORT | Application port | 8080 |

## Development

### Adding New Features

1. Create models in `models/` directory
2. Add handlers in `handlers/` directory
3. Create templates in `templates/` directory
4. Add routes in `main.go`
5. Write tests in `tests/` directory

### Code Style

This project follows standard Go conventions:
- Use `gofmt` for formatting
- Follow effective Go guidelines
- Write tests for all handlers and models

## Architecture

### In-Memory Database

The application uses a thread-safe in-memory database implementation (`config/mock_database.go`) that:
- Stores contacts in a map with mutex protection
- Validates unique email addresses
- Provides CRUD operations
- Requires no external database

### Optional PostgreSQL Support

The codebase includes PostgreSQL connector code (`config/database.go`) that can be enabled if you need persistent storage. To use PostgreSQL:

1. Install PostgreSQL
2. Update `main.go` to use `config.InitDB()` instead of `config.NewMockDB()`
3. Set environment variables for database connection
4. Run migrations

## Troubleshooting

### Port Already in Use

If port 8080 is already in use, change the PORT environment variable:

```bash
export PORT=3000
go run main.go
```

### Templates Not Found

Make sure you're running the application from the project root directory where the `templates/` folder is located.

## Project Highlights

- **Zero Configuration**: No database setup required
- **Fast Startup**: Application starts in seconds
- **Complete Test Suite**: All features tested with 100% pass rate
- **Modern UI**: Responsive design with smooth animations
- **Clean Architecture**: Well-organized code structure
- **Production Ready**: Includes Dockerfile and build scripts

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- Go standard library

## Support

For issues and questions, please open an issue on the GitHub repository.