# GoSpace

A comprehensive web application built with Go and the Gin framework, featuring a calculator with history tracking, contact form with PostgreSQL database storage, and a modern responsive UI.

## Features

- **Hello World Home Page**: Landing page with navigation to all features
- **Calculator**: Perform basic arithmetic operations (addition, subtraction, multiplication, division)
- **Calculator History**: View and manage calculation history with delete functionality
- **Contact Form**: Submit and store contact information (name, surname, email) in PostgreSQL
- **Contact List**: View all stored contacts from the database
- **Responsive Design**: Modern, mobile-friendly UI
- **No JavaScript**: Pure HTML/CSS with server-side rendering for maximum compatibility and security
- **PostgreSQL Database**: Persistent storage with GORM ORM
- **CSRF Protection**: All POST endpoints protected against cross-site request forgery attacks
- **Session Management**: Secure session handling for CSRF tokens

## Project Structure

```
gospace/
├── main.go                 # Application entry point
├── gospace                 # Compiled binary
├── config/                 # Configuration files
│   ├── config.go          # Environment configuration
│   ├── database.go        # PostgreSQL connector
│   └── mock_database.go   # Legacy mock database (deprecated)
├── models/                 # Data models
│   ├── contact.go         # Contact model
│   └── calculator_history.go # Calculator history model
├── handlers/               # HTTP request handlers
│   ├── handler.go         # Base handler
│   ├── home.go            # Home page handler
│   ├── calculator.go      # Calculator handlers
│   └── contact.go         # Contact form handlers
├── templates/              # HTML templates
│   ├── home.html          # Home page template
│   ├── calculator.html    # Calculator page template
│   ├── calculator_history.html # Calculator history template
│   ├── contact.html       # Contact form template
│   └── contacts_list.html # Contacts list template
├── static/                 # Static assets
│   └── css/
│       └── style.css      # Main stylesheet
├── tests/                  # Test files
│   ├── test_helpers.go    # Shared test utilities
│   ├── handlers_test.go   # Handler tests
│   ├── calculator_history_test.go # Calculator history tests
│   └── integration_test.go # Integration tests
├── go.mod                  # Go module file
├── go.sum                  # Go dependencies
├── .env.example           # Environment variables example
├── .gitignore             # Git ignore file
├── Dockerfile             # Docker configuration
├── docker-compose.yml     # Docker Compose configuration
├── k8s-deployment.yaml    # Kubernetes app deployment
├── k8s-postgres.yaml      # Kubernetes PostgreSQL deployment
├── deploy-docker.sh       # Docker deployment script
├── deploy-k8s.sh          # Kubernetes deployment script
├── README.md              # This file
├── QUICKSTART.md          # Quick start guide
├── DEPLOYMENT.md          # Comprehensive deployment guide
├── KUBERNETES.md          # Kubernetes deployment guide
└── API.md                 # API documentation
```

## Prerequisites

- Go 1.21 or higher
- Git
- PostgreSQL 15+ (or use Docker/Kubernetes deployment)

## Installation

### Option 1: Local Development

1. **Clone the Repository**

```bash
git clone <repository-url>
cd gospace
```

2. **Set up PostgreSQL**

Create a database and update `.env` file:

```bash
cp .env.example .env
# Edit .env with your PostgreSQL credentials
```

3. **Install Dependencies**

```bash
go mod download
```

4. **Run the Application**

```bash
go run main.go
```

5. **Access the Application**

Open your browser and navigate to `http://localhost:8080`

### Option 2: Docker Compose (Recommended)

The easiest way to run the application with all dependencies:

```bash
./deploy-docker.sh
```

Or manually:

```bash
docker build -t gospace:latest .
docker-compose up -d
```

Access at `http://localhost:8080`

### Option 3: Kubernetes

Deploy to Kubernetes cluster:

```bash
./deploy-k8s.sh
```

Or manually:

```bash
# Build image
docker build -t gospace:latest .

# Deploy PostgreSQL
kubectl apply -f k8s-postgres.yaml

# Deploy application
kubectl apply -f k8s-deployment.yaml
```

Access via NodePort at `http://localhost:30080`

## Usage

### Home Page
Navigate to `http://localhost:8080` to see the landing page with links to all features.

### Calculator
1. Go to `http://localhost:8080/calculator`
2. Enter two numbers
3. Select an operation (add, subtract, multiply, divide)
4. Click "Calculate" to see the result
5. View calculation history at `http://localhost:8080/calculator/history`
6. Delete past calculations from the history page

### Contact Form
1. Go to `http://localhost:8080/contact`
2. Fill in your name, surname, and email
3. Click "Submit" to save to the PostgreSQL database
4. View all contacts at `http://localhost:8080/contacts`

**Note**: Data is stored in PostgreSQL and persists across application restarts.

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Home page |
| GET | `/calculator` | Calculator page |
| POST | `/calculator` | Calculate result |
| GET | `/calculator/history` | View calculation history |
| POST | `/calculator/history/:id/delete` | Delete calculation |
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
go build -o gospace main.go

# Run the binary
./gospace
```

## Docker Deployment

### Using Docker Compose (with PostgreSQL)

The project includes a [`docker-compose.yml`](docker-compose.yml) that sets up both the application and PostgreSQL database:

```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

### Build Docker Image Only

```bash
docker build -t gospace .
```

### Run with Docker (Standalone)

```bash
# Requires external PostgreSQL
docker run -p 8080:8080 \
  -e DB_HOST=host.docker.internal \
  -e DB_PORT=5432 \
  -e DB_USER=postgres \
  -e DB_PASSWORD=postgres \
  -e DB_NAME=gospace \
  gospace
```

**Note**: The application requires PostgreSQL. Use Docker Compose for a complete setup with database included.

## Environment Variables

The application supports the following environment variables (see [`.env.example`](.env.example)):

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| PORT | Application port | 8080 | No |
| DB_HOST | PostgreSQL host | localhost | Only if using PostgreSQL |
| DB_PORT | PostgreSQL port | 5432 | Only if using PostgreSQL |
| DB_USER | PostgreSQL user | postgres | Only if using PostgreSQL |
| DB_PASSWORD | PostgreSQL password | postgres | Only if using PostgreSQL |
| DB_NAME | PostgreSQL database name | gospace | Only if using PostgreSQL |
| DB_SSLMODE | PostgreSQL SSL mode | disable | Only if using PostgreSQL |

**Note**: Database environment variables are only needed if you switch from in-memory storage to PostgreSQL. By default, the application uses in-memory storage and only requires the PORT variable (which defaults to 8080).

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