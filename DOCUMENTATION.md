# Documentation Index

This document provides an overview of all documentation available for the GoSpace project.

## Quick Links

- **[README.md](README.md)** - Main project documentation with features, installation, and usage
- **[QUICKSTART.md](QUICKSTART.md)** - Get started in under 2 minutes
- **[API.md](API.md)** - Complete API endpoint documentation
- **[KUBERNETES.md](KUBERNETES.md)** - Kubernetes deployment guide

## Project Overview

GoSpace is a comprehensive web application built with Go and the Gin framework. It features:
- Calculator functionality
- Contact form with in-memory database storage
- Modern responsive UI
- No database setup required (uses in-memory storage by default)

## Documentation Structure

### Getting Started

1. **[QUICKSTART.md](QUICKSTART.md)** - Start here for the fastest setup
   - Prerequisites
   - Installation steps
   - Testing features
   - Building for production

2. **[README.md](README.md)** - Comprehensive project documentation
   - Features overview
   - Project structure
   - Installation guide
   - Usage instructions
   - Testing guide
   - Docker deployment
   - Environment variables

### API Reference

3. **[API.md](API.md)** - Complete API documentation
   - All endpoints with handlers
   - Request/response formats
   - Data models
   - Error handling
   - Static assets

### Deployment

4. **[KUBERNETES.md](KUBERNETES.md)** - Kubernetes deployment
   - Quick deployment guide
   - Resource specifications
   - Useful commands
   - Troubleshooting
   - Production considerations

## Project Structure

```
gospace/
├── main.go                 # Application entry point
├── gospace                 # Compiled binary (gitignored)
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
├── k8s-deployment.yaml    # Kubernetes deployment manifest
├── README.md              # Main documentation
├── QUICKSTART.md          # Quick start guide
├── API.md                 # API documentation
├── KUBERNETES.md          # Kubernetes guide
└── DOCUMENTATION.md       # This file
```

## Key Files

### Configuration Files

- **[.env.example](.env.example)** - Environment variable template
  - Database configuration (optional)
  - Server configuration

- **[go.mod](go.mod)** - Go module dependencies
  - Gin framework
  - PostgreSQL driver
  - Testing libraries

### Deployment Files

- **[Dockerfile](Dockerfile)** - Multi-stage Docker build
  - Build stage with Go 1.25
  - Runtime stage with Alpine Linux
  - Includes templates and static files

- **[docker-compose.yml](docker-compose.yml)** - Docker Compose setup
  - PostgreSQL service
  - Application service
  - Health checks and dependencies

- **[k8s-deployment.yaml](k8s-deployment.yaml)** - Kubernetes manifests
  - Namespace
  - Deployment (2 replicas)
  - NodePort Service (port 30080)
  - Optional Ingress

### Source Code

- **[main.go](main.go)** - Application entry point
  - Initializes mock database
  - Sets up Gin router
  - Loads templates and static files
  - Defines routes

- **[config/mock_database.go](config/mock_database.go)** - In-memory database
  - Thread-safe operations
  - Contact CRUD operations
  - Email uniqueness validation

- **[handlers/](handlers/)** - Request handlers
  - [`home.go`](handlers/home.go) - Home page
  - [`calculator.go`](handlers/calculator.go) - Calculator operations
  - [`contact.go`](handlers/contact.go) - Contact form and list

- **[models/contact.go](models/contact.go)** - Contact data model
  - ID, Name, Surname, Email
  - Timestamps (CreatedAt, UpdatedAt)

### Tests

- **[tests/handlers_test.go](tests/handlers_test.go)** - Handler unit tests
  - Home page test
  - Calculator tests
  - Contact form tests

- **[tests/integration_test.go](tests/integration_test.go)** - Integration tests
  - Full user flow testing
  - End-to-end scenarios

## Common Tasks

### Development

```bash
# Run the application
go run main.go

# Run tests
go test ./tests -v

# Run tests with coverage
go test ./tests -cover

# Build binary
go build -o gospace main.go
```

### Docker

```bash
# Using Docker Compose
docker-compose up -d

# Build Docker image
docker build -t gospace .

# Run Docker container
docker run -p 8080:8080 gospace
```

### Kubernetes

```bash
# Deploy to Kubernetes
kubectl apply -f k8s-deployment.yaml

# Check status
kubectl get all -n gospace

# View logs
kubectl logs -n gospace -l app=gospace -f

# Access application
open http://localhost:30080
```

## Features by Document

### README.md
- Complete feature list
- Installation instructions
- Usage guide
- Testing instructions
- Docker deployment
- Environment variables
- Troubleshooting

### QUICKSTART.md
- 2-minute setup
- Quick testing guide
- Basic customization
- Docker quick start

### API.md
- All endpoints documented
- Request/response formats
- Data models with examples
- Error handling
- Handler references

### KUBERNETES.md
- Deployment guide
- Resource specifications
- Scaling instructions
- Health checks
- Production considerations
- Troubleshooting

## Technology Stack

- **Language**: Go 1.21+
- **Framework**: Gin Web Framework
- **Database**: In-memory (default) or PostgreSQL (optional)
- **Containerization**: Docker
- **Orchestration**: Kubernetes
- **Testing**: Go testing package, testify

## Support

For issues or questions:
1. Check the relevant documentation file
2. Review the troubleshooting sections
3. Check test files for usage examples
4. Open an issue on the project repository

## Contributing

When contributing to documentation:
1. Keep documentation in sync with code changes
2. Update all relevant documentation files
3. Test all code examples
4. Use clear, concise language
5. Include code references with line numbers where applicable

## License

This project is licensed under the MIT License.