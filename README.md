# GoSpace

A comprehensive web application built with Go and the Gin framework, featuring a calculator with history tracking, contact form with PostgreSQL database storage, and a modern responsive UI.

## Features

- **Hello World Home Page**: Landing page with horizontal feature card layout and navigation to all features
- **Calculator**: Perform basic arithmetic operations (addition, subtraction, multiplication, division)
- **Calculator History**: View and manage calculation history with delete functionality
- **Contact Form**: Submit and store contact information (name, surname, email) in PostgreSQL
- **Contact List**: View all stored contacts from the database
- **AI Agents Repository**: Manage AI agents with metadata, source code, and categorization
- **AI Tools Repository**: Manage AI tools, frameworks, and libraries with detailed information
- **Modern Responsive UI**: Horizontal layouts, wide forms (1800px), and consistent navigation
- **No JavaScript**: Pure HTML/CSS with server-side rendering for maximum compatibility and security
- **PostgreSQL Database**: Persistent storage with GORM ORM
- **Unique Contact Emails**: Duplicate contact submissions are rejected at the database layer

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
│   ├── calculator_history.go # Calculator history model
│   ├── agent.go           # AI Agent model
│   └── tool.go            # AI Tool model
├── handlers/               # HTTP request handlers
│   ├── handler.go         # Base handler
│   ├── home.go            # Home page handler
│   ├── calculator.go      # Calculator handlers
│   ├── contact.go         # Contact form handlers
│   ├── agents.go          # AI Agents handlers
│   └── tools.go           # AI Tools handlers
├── templates/              # HTML templates
│   ├── home.html          # Home page template
│   ├── calculator.html    # Calculator page template
│   ├── calculator_history.html # Calculator history template
│   ├── contact.html       # Contact form template
│   ├── contacts_list.html # Contacts list template
│   ├── agents.html        # AI Agents list template
│   ├── add_agent.html     # Add AI Agent form
│   ├── edit_agent.html    # Edit AI Agent form
│   ├── tools.html         # AI Tools list template
│   ├── add_tool.html      # Add AI Tool form
│   └── edit_tool.html     # Edit AI Tool form
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
- PostgreSQL 15+ (local install optional if you use the provided Docker or Kubernetes PostgreSQL deployments)

## Installation

### Option 1: Local Development

1. **Clone the Repository**

```bash
git clone <repository-url>
cd gospace
```

2. **Set up PostgreSQL**

Choose one of the following PostgreSQL options, then update [`.env.example`](.env.example) values in your local [`.env`](.env.example) copy if needed:

**Option A: Run PostgreSQL in Docker**
```bash
docker run --name gospace-postgres \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=gin_webapp \
  -p 5432:5432 \
  -v gospace-postgres-data:/var/lib/postgresql/data \
  -d postgres:15-alpine
```

**Option B: Use the project's Docker Compose stack**
```bash
cp .env.example .env
docker-compose up -d postgres
```

**Option C: Use a local PostgreSQL installation**
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

The easiest way to run the application with PostgreSQL in a container is to use the included [`docker-compose.yml`](docker-compose.yml):

```bash
cp .env.example .env
./deploy-docker.sh
```

Or manually:

```bash
cp .env.example .env
docker build -t gospace:latest .
docker-compose up -d
```

This starts:
- a PostgreSQL container on `localhost:5432`
- the GoSpace application on `http://localhost:8080`

Access at `http://localhost:8080`

### Option 3: Kubernetes

Deploy both PostgreSQL and the application to Kubernetes:

```bash
./deploy-k8s.sh
```

Or manually:

```bash
# Build and push image
docker build -t gospace:latest -t localhost:5000/gospace:latest -t vvk17/gospace:latest .
docker push localhost:5000/gospace:latest
docker push vvk17/gospace:latest

# Deploy PostgreSQL inside the cluster
kubectl apply -f k8s-postgres.yaml
kubectl wait --for=condition=ready pod -l app=postgres -n gospace-db --timeout=120s

# Deploy application
kubectl apply -f k8s-deployment.yaml
```

The PostgreSQL database runs in-cluster in the `gospace-db` namespace behind the `postgres-service` service and stores data on the `postgres-pvc` persistent volume claim.

Access via NodePort at `http://localhost:30081`

## Usage

### Home Page
Navigate to `http://localhost:8080` to see the landing page with horizontally arranged feature cards and links to all features.

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

### AI Agents Repository
1. Go to `http://localhost:8080/agents` to view all AI agents
2. Click "Add New Agent" to create a new agent with metadata and source code
3. Edit existing agents by clicking the edit icon
4. Delete agents by clicking the delete icon
5. View agent details including category, version, author, repository, tags, and status

### AI Tools Repository
1. Go to `http://localhost:8080/tools` to view all AI tools
2. Click "Add New Tool" to create a new tool entry
3. Edit existing tools by clicking the edit icon
4. Delete tools by clicking the delete icon
5. View tool details including category, language, version, author, repository, tags, and status

**Note**: All data is stored in PostgreSQL and persists across application restarts.

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check endpoint |
| GET | `/` | Home page |
| GET | `/calculator` | Calculator page |
| POST | `/calculator` | Calculate result |
| GET | `/calculator/history` | View calculation history |
| POST | `/calculator/history/:id/delete` | Delete calculation |
| GET | `/contact` | Contact form page |
| POST | `/contact` | Submit contact |
| GET | `/contacts` | List all contacts |
| GET | `/agents` | List all AI agents |
| GET | `/agents/add` | Add AI agent form |
| POST | `/agents/add` | Create new AI agent |
| GET | `/agents/:id/edit` | Edit AI agent form |
| POST | `/agents/:id/edit` | Update AI agent |
| POST | `/agents/:id/delete` | Delete AI agent |
| GET | `/tools` | List all AI tools |
| GET | `/tools/add` | Add AI tool form |
| POST | `/tools/add` | Create new AI tool |
| GET | `/tools/:id/edit` | Edit AI tool form |
| POST | `/tools/:id/edit` | Update AI tool |
| POST | `/tools/:id/delete` | Delete AI tool |

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

The project includes a [`docker-compose.yml`](docker-compose.yml) that starts PostgreSQL as a containerized dependency for the application.

```bash
# Create local environment file
cp .env.example .env

# Start PostgreSQL and the app
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

To start only PostgreSQL in Docker and run the app locally with [`main.go`](main.go:13):

```bash
cp .env.example .env
docker-compose up -d postgres
go run main.go
```

### Build Docker Image Only

```bash
docker build -t gospace .
```

### Run with Docker (Standalone)

```bash
# Start PostgreSQL separately
docker run --name gospace-postgres \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=gin_webapp \
  -p 5432:5432 \
  -v gospace-postgres-data:/var/lib/postgresql/data \
  -d postgres:15-alpine

# Run the app container against that PostgreSQL instance
docker run -p 8080:8080 \
  -e DB_HOST=host.docker.internal \
  -e DB_PORT=5432 \
  -e DB_USER=postgres \
  -e DB_PASSWORD=postgres \
  -e DB_NAME=gin_webapp \
  -e DB_SSLMODE=disable \
  gospace
```

**Note**: The application requires PostgreSQL for runtime storage. Use [`docker-compose.yml`](docker-compose.yml) for the simplest containerized setup.

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

- **Modern UI**: Horizontal layouts, wide forms (1800px), and consistent navigation
- **AI Repository Management**: Manage AI agents and tools with full CRUD operations
- **Fast Startup**: Application starts in seconds
- **Complete Test Suite**: All features tested with 100% pass rate
- **Responsive Design**: Mobile-friendly with smooth animations
- **Clean Architecture**: Well-organized code structure
- **Production Ready**: Includes Dockerfile, Kubernetes manifests, and deployment scripts
- **Standardized Deployment**: Single image name (`gospace:latest`) for both Docker and Kubernetes

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