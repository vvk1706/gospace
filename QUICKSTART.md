# Quick Start Guide

This guide will help you get GoSpace up and running in **under 2 minutes**.

## Prerequisites

- Go 1.21 or higher installed
- Git

**Database required:** The application uses PostgreSQL as its primary runtime database for persistent storage.

## Quick Setup (2 minutes)

### 1. Clone and Navigate

```bash
git clone <repository-url>
cd gospace
```

### 2. Set Up PostgreSQL

You have three supported options:

**Option A: Docker Compose (Recommended)**
```bash
cp .env.example .env

# Start PostgreSQL and application
docker-compose up -d

# Access at http://localhost:8080
```

**Option B: PostgreSQL container only**
```bash
cp .env.example .env

docker run --name gospace-postgres \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=gin_webapp \
  -p 5432:5432 \
  -v gospace-postgres-data:/var/lib/postgresql/data \
  -d postgres:15-alpine

go mod download
go run main.go
```

**Option C: Local PostgreSQL**
```bash
# Install PostgreSQL, then create database if needed
createdb gin_webapp

# Copy environment file
cp .env.example .env

# Edit .env with your database credentials
go mod download
go run main.go
```

You should see:
```
Connected to PostgreSQL database
Database migration completed
Starting server on port 8080...
Using PostgreSQL database for persistent storage
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
- View calculation history at `http://localhost:8080/calculator/history`
- Delete past calculations from the history page

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
- See all contacts stored in PostgreSQL

**Note**: Data is stored in PostgreSQL and persists across application restarts.

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
# Start PostgreSQL in a container
docker run --name gospace-postgres \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=gin_webapp \
  -p 5432:5432 \
  -v gospace-postgres-data:/var/lib/postgresql/data \
  -d postgres:15-alpine

# Build the image
docker build -t gospace .

# Run the application container
docker run -p 8080:8080 \
  -e DB_HOST=host.docker.internal \
  -e DB_PORT=5432 \
  -e DB_USER=postgres \
  -e DB_PASSWORD=postgres \
  -e DB_NAME=gin_webapp \
  -e DB_SSLMODE=disable \
  gospace
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

✅ **PostgreSQL Database** - Persistent data storage with GORM
✅ **Fast Startup** - Application starts in seconds
✅ **Complete Functionality** - Calculator with history, forms, and data storage
✅ **Modern UI** - Responsive design with pure HTML/CSS (no JavaScript)
✅ **Fully Tested** - Comprehensive test suite with 31 passing tests

## What's Stored in Database?

- Contact records (name, surname, email)
- Calculator history (operations and results)
- Automatic ID generation
- Duplicate email validation
- Optimistic locking for concurrent updates

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