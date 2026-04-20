# Quick Start Guide

This guide will help you get the Gin Web Application up and running in **under 2 minutes**.

## Prerequisites

- Go 1.21 or higher installed
- Git

**No database required!** The application uses in-memory storage.

## Quick Setup (2 minutes)

### 1. Clone and Navigate

```bash
git clone <repository-url>
cd gospace
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Run the Application

```bash
go run main.go
```

You should see:
```
Using in-memory mock database (no PostgreSQL required)
Starting server on port 8080...
No database setup required - using in-memory storage
Access the application at http://localhost:8080
```

### 4. Access the Application

Open your browser and go to `http://localhost:8080`

**That's it!** No database setup, no configuration files, no complex installation.

## Testing the Features

### 1. Home Page
- Navigate to `http://localhost:8080`
- You'll see three feature cards

### 2. Calculator
- Click "Try Calculator" or go to `http://localhost:8080/calculator`
- Enter two numbers (e.g., 10 and 5)
- Select an operation (add, subtract, multiply, divide)
- Click "Calculate" to see the result

### 3. Contact Form
- Click "Add Contact" or go to `http://localhost:8080/contact`
- Fill in:
  - Name: Your first name
  - Surname: Your last name
  - Email: Your email address
- Click "Submit"
- You'll see a success message

### 4. View Contacts
- Click "View All" or go to `http://localhost:8080/contacts`
- See all contacts stored in memory

**Note**: Data is stored in memory and will be cleared when you restart the application.

## Running Tests

```bash
# Run all tests
go test ./tests -v

# All tests should pass!
```

## Building for Production

```bash
# Build the binary
go build -o gospace main.go

# Run the binary
./gospace
```

## Docker Deployment

### Using Docker Compose (Recommended)

```bash
# Start all services (app + PostgreSQL)
docker-compose up -d

# View logs
docker-compose logs -f app

# Stop services
docker-compose down
```

### Using Docker Only

```bash
# Build the image
docker build -t gospace .

# Run the container
docker run -p 8080:8080 gospace
```

## Customization

### Change Port

```bash
export PORT=3000
go run main.go
```

### Build with Docker

```bash
docker build -t gospace .
docker run -p 8080:8080 gospace
```

## Key Features

✅ **Zero Configuration** - No database setup required  
✅ **Fast Startup** - Application starts in seconds  
✅ **Complete Functionality** - Calculator, forms, and data storage  
✅ **Modern UI** - Responsive design with animations  
✅ **Fully Tested** - Comprehensive test suite included  

## What's Stored in Memory?

- Contact records (name, surname, email)
- Automatic ID generation
- Duplicate email validation
- Thread-safe operations

## Next Steps

- Read the full [`README.md`](README.md) for detailed documentation
- Check [`API.md`](API.md) for API endpoint details
- Review [`KUBERNETES.md`](KUBERNETES.md) for Kubernetes deployment
- Explore the code structure in the project directories
- Modify templates in [`templates/`](templates/) directory
- Customize styles in [`static/css/style.css`](static/css/style.css)
- Add new handlers in [`handlers/`](handlers/) directory

## Troubleshooting

### Port 8080 already in use
```bash
export PORT=3000
go run main.go
```

### Templates not found
Make sure you're running the application from the project root directory.

### Module errors
```bash
go mod tidy
go mod download
```

## Support

For issues or questions, check the README.md or open an issue on the project repository.

---

**Enjoy building with Gin!** 🚀