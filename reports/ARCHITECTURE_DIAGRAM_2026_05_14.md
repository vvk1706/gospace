# GoSpace Application Architecture Diagram
**Date:** 2026-05-14  
**Version:** 1.0  
**Purpose:** Complete architectural overview of the GoSpace application

---

## Table of Contents
1. [High-Level Architecture](#high-level-architecture)
2. [Package Structure](#package-structure)
3. [Data Models & Relationships](#data-models--relationships)
4. [Handler Functions & Routes](#handler-functions--routes)
5. [Database Schema](#database-schema)
6. [Request Flow](#request-flow)
7. [Deployment Architecture](#deployment-architecture)

---

## High-Level Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                         CLIENT LAYER                             │
│  (Web Browser - HTML/CSS, No JavaScript)                        │
└────────────────────────┬────────────────────────────────────────┘
                         │ HTTP Requests
                         ▼
┌─────────────────────────────────────────────────────────────────┐
│                      WEB SERVER LAYER                            │
│  ┌──────────────────────────────────────────────────────────┐  │
│  │              Gin Web Framework (Go)                       │  │
│  │  - Router                                                 │  │
│  │  - Middleware                                             │  │
│  │  - Template Engine                                        │  │
│  └──────────────────────────────────────────────────────────┘  │
└────────────────────────┬────────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────────┐
│                    APPLICATION LAYER                             │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐         │
│  │   Handlers   │  │    Models    │  │    Config    │         │
│  │  (Business   │  │  (Data       │  │  (Settings)  │         │
│  │   Logic)     │  │  Structures) │  │              │         │
│  └──────┬───────┘  └──────┬───────┘  └──────┬───────┘         │
│         │                  │                  │                  │
│         └──────────────────┴──────────────────┘                 │
│                            │                                     │
└────────────────────────────┼─────────────────────────────────────┘
                             │ GORM ORM
                             ▼
┌─────────────────────────────────────────────────────────────────┐
│                      DATABASE LAYER                              │
│  ┌──────────────────────────────────────────────────────────┐  │
│  │              PostgreSQL 15                                │  │
│  │  - Contacts Table                                         │  │
│  │  - Calculator History Table                               │  │
│  │  - Agents Table                                           │  │
│  │  - Tools Table                                            │  │
│  └──────────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────┘
```

---

## Package Structure

```
gospace/
│
├── main.go                          # Application Entry Point
│   ├── Initializes Gin router
│   ├── Loads configuration
│   ├── Connects to database
│   ├── Registers routes
│   └── Starts HTTP server
│
├── config/                          # Configuration Package
│   ├── config.go                    # Environment configuration
│   │   ├── Config struct
│   │   ├── LoadConfig()
│   │   ├── GetDSN()
│   │   └── getEnv()
│   │
│   ├── database.go                  # Database connector
│   │   └── InitDB()
│   │
│   └── mock_database.go             # Legacy (deprecated)
│
├── models/                          # Data Models Package
│   ├── contact.go
│   │   ├── Contact struct
│   │   ├── NewContact()
│   │   └── TableName()
│   │
│   ├── calculator_history.go
│   │   ├── CalculatorHistory struct
│   │   └── NewCalculatorHistory()
│   │
│   ├── agent.go
│   │   └── Agent struct
│   │
│   └── tool.go
│       └── Tool struct
│
├── handlers/                        # HTTP Handlers Package
│   ├── handler.go                   # Base handler
│   │   ├── Handler struct
│   │   └── NewHandler()
│   │
│   ├── home.go                      # Home page
│   │   └── Home()
│   │
│   ├── calculator.go                # Calculator handlers
│   │   ├── Calculator()
│   │   ├── CalculateResult()
│   │   ├── ListCalculatorHistory()
│   │   ├── DeleteCalculatorHistory()
│   │   ├── performCalculation()
│   │   └── validateOperation()
│   │
│   ├── contact.go                   # Contact handlers
│   │   ├── ContactForm()
│   │   ├── SubmitContact()
│   │   └── ListContacts()
│   │
│   ├── agents.go                    # Agent handlers
│   │   ├── ListAgents()
│   │   ├── AddAgentForm()
│   │   ├── CreateAgent()
│   │   ├── EditAgentForm()
│   │   ├── UpdateAgent()
│   │   └── DeleteAgent()
│   │
│   └── tools.go                     # Tool handlers
│       ├── ListTools()
│       ├── AddToolForm()
│       ├── CreateTool()
│       ├── EditToolForm()
│       ├── UpdateTool()
│       └── DeleteTool()
│
├── templates/                       # HTML Templates
│   ├── home.html
│   ├── calculator.html
│   ├── calculator_history.html
│   ├── contact.html
│   ├── contacts_list.html
│   ├── agents.html
│   ├── add_agent.html
│   ├── edit_agent.html
│   ├── tools.html
│   ├── add_tool.html
│   └── edit_tool.html
│
└── static/                          # Static Assets
    └── css/
        └── style.css
```

---

## Data Models & Relationships

### Entity Relationship Diagram

```
┌─────────────────────────┐
│       Contact           │
├─────────────────────────┤
│ ID (PK)        uint     │
│ Name           string   │
│ Surname        string   │
│ Email          string   │ ← UNIQUE INDEX
│ CreatedAt      time     │
│ UpdatedAt      time     │
└─────────────────────────┘

┌─────────────────────────┐
│  CalculatorHistory      │
├─────────────────────────┤
│ ID (PK)        uint     │
│ Num1           float64  │
│ Num2           float64  │
│ Operation      string   │
│ Result         float64  │
│ Version        int      │
│ CreatedAt      time     │
│ UpdatedAt      time     │
└─────────────────────────┘

┌─────────────────────────┐
│        Agent            │
├─────────────────────────┤
│ ID (PK)        uint     │
│ CreatedAt      time     │
│ UpdatedAt      time     │
│ DeletedAt      time     │ ← Soft Delete
│ Name           string   │
│ Description    text     │
│ Category       string   │
│ Version        string   │
│ Author         string   │
│ Repository     string   │
│ Tags           text     │
│ Status         string   │
│ SourceCode     text     │ ← NEW
└─────────────────────────┘

┌─────────────────────────┐
│         Tool            │
├─────────────────────────┤
│ ID (PK)        uint     │
│ CreatedAt      time     │
│ UpdatedAt      time     │
│ DeletedAt      time     │ ← Soft Delete
│ Name           string   │
│ Description    text     │
│ Category       string   │
│ Version        string   │
│ Author         string   │
│ Repository     string   │
│ Tags           text     │
│ Language       string   │
│ Status         string   │
│ SourceCode     text     │ ← NEW
└─────────────────────────┘

Note: Currently no foreign key relationships between tables.
Each entity is independent.
```

### Model Structures Detail

```go
// Contact Model
type Contact struct {
    ID        uint      `json:"id"`
    Name      string    `json:"name"`
    Surname   string    `json:"surname"`
    Email     string    `gorm:"uniqueIndex;not null" json:"email"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// CalculatorHistory Model
type CalculatorHistory struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Num1      float64   `json:"num1"`
    Num2      float64   `json:"num2"`
    Operation string    `json:"operation"`
    Result    float64   `json:"result"`
    Version   int       `gorm:"default:0" json:"version"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// Agent Model
type Agent struct {
    ID          uint           `gorm:"primarykey" json:"id"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
    Name        string         `gorm:"not null;size:255" json:"name"`
    Description string         `gorm:"type:text" json:"description"`
    Category    string         `gorm:"size:100" json:"category"`
    Version     string         `gorm:"size:50" json:"version"`
    Author      string         `gorm:"size:255" json:"author"`
    Repository  string         `gorm:"size:500" json:"repository"`
    Tags        string         `gorm:"type:text" json:"tags"`
    Status      string         `gorm:"size:50;default:'active'" json:"status"`
    SourceCode  string         `gorm:"type:text" json:"source_code"`
}

// Tool Model
type Tool struct {
    ID          uint           `gorm:"primarykey" json:"id"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
    Name        string         `gorm:"not null;size:255" json:"name"`
    Description string         `gorm:"type:text" json:"description"`
    Category    string         `gorm:"size:100" json:"category"`
    Version     string         `gorm:"size:50" json:"version"`
    Author      string         `gorm:"size:255" json:"author"`
    Repository  string         `gorm:"size:500" json:"repository"`
    Tags        string         `gorm:"type:text" json:"tags"`
    Language    string         `gorm:"size:50" json:"language"`
    Status      string         `gorm:"size:50;default:'active'" json:"status"`
    SourceCode  string         `gorm:"type:text" json:"source_code"`
}
```

---

## Handler Functions & Routes

### Route Mapping

```
┌──────────┬─────────────────────────────┬────────────────────────────┐
│ Method   │ Route                       │ Handler Function           │
├──────────┼─────────────────────────────┼────────────────────────────┤
│ GET      │ /                           │ Home()                     │
│ GET      │ /calculator                 │ Calculator()               │
│ POST     │ /calculator                 │ CalculateResult()          │
│ GET      │ /calculator/history         │ ListCalculatorHistory()    │
│ POST     │ /calculator/history/:id/del │ DeleteCalculatorHistory()  │
│ GET      │ /contact                    │ ContactForm()              │
│ POST     │ /contact                    │ SubmitContact()            │
│ GET      │ /contacts                   │ ListContacts()             │
│ GET      │ /agents                     │ ListAgents()               │
│ GET      │ /agents/add                 │ AddAgentForm()             │
│ POST     │ /agents/add                 │ CreateAgent()              │
│ GET      │ /agents/:id/edit            │ EditAgentForm()            │
│ POST     │ /agents/:id/edit            │ UpdateAgent()              │
│ POST     │ /agents/:id/delete          │ DeleteAgent()              │
│ GET      │ /tools                      │ ListTools()                │
│ GET      │ /tools/add                  │ AddToolForm()              │
│ POST     │ /tools/add                  │ CreateTool()               │
│ GET      │ /tools/:id/edit             │ EditToolForm()             │
│ POST     │ /tools/:id/edit             │ UpdateTool()               │
│ POST     │ /tools/:id/delete           │ DeleteTool()               │
└──────────┴─────────────────────────────┴────────────────────────────┘
```

### Handler Dependencies

```
┌─────────────────────────────────────────────────────────────────┐
│                        Handler Struct                            │
│  ┌──────────────────────────────────────────────────────────┐  │
│  │  type Handler struct {                                    │  │
│  │      DB *gorm.DB  ← Database connection                   │  │
│  │  }                                                         │  │
│  └──────────────────────────────────────────────────────────┘  │
│                                                                  │
│  All handler methods are attached to this struct:               │
│  func (h *Handler) HandlerName(c *gin.Context) { ... }         │
└─────────────────────────────────────────────────────────────────┘

Handler Creation Flow:
main.go → NewHandler(db) → Handler{DB: db} → Attach to routes
```

### Handler Function Relationships

```
Calculator Module:
┌──────────────────┐
│  Calculator()    │ ← Display form
└────────┬─────────┘
         │
         ▼
┌──────────────────┐     ┌─────────────────────┐
│ CalculateResult()│────→│ performCalculation()│ (helper)
└────────┬─────────┘     └─────────────────────┘
         │                ┌─────────────────────┐
         ├───────────────→│ validateOperation() │ (helper)
         │                └─────────────────────┘
         │
         ├─→ Save to DB (CalculatorHistory)
         │
         └─→ Redirect to /calculator/history
                          │
                          ▼
              ┌──────────────────────────┐
              │ ListCalculatorHistory()  │
              └────────┬─────────────────┘
                       │
                       ▼
              ┌──────────────────────────┐
              │DeleteCalculatorHistory() │
              └──────────────────────────┘

Contact Module:
┌──────────────────┐
│  ContactForm()   │ ← Display form
└────────┬─────────┘
         │
         ▼
┌──────────────────┐
│ SubmitContact()  │ → Save to DB (Contact)
└────────┬─────────┘
         │
         └─→ Show success/error
         
┌──────────────────┐
│  ListContacts()  │ ← Display all contacts
└──────────────────┘

Agent Module:
┌──────────────────┐
│   ListAgents()   │ ← Display all agents
└────────┬─────────┘
         │
         ├─→ Edit button → EditAgentForm()
         │                      │
         │                      ▼
         │                 UpdateAgent() → Save to DB
         │
         ├─→ Add button → AddAgentForm()
         │                      │
         │                      ▼
         │                 CreateAgent() → Save to DB
         │
         └─→ Delete button → DeleteAgent() → Remove from DB

Tool Module:
┌──────────────────┐
│   ListTools()    │ ← Display all tools
└────────┬─────────┘
         │
         ├─→ Edit button → EditToolForm()
         │                      │
         │                      ▼
         │                 UpdateTool() → Save to DB
         │
         ├─→ Add button → AddToolForm()
         │                      │
         │                      ▼
         │                 CreateTool() → Save to DB
         │
         └─→ Delete button → DeleteTool() → Remove from DB
```

---

## Database Schema

### PostgreSQL Tables

```sql
-- contacts table
CREATE TABLE contacts (
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    surname    VARCHAR(255) NOT NULL,
    email      VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_contacts_email ON contacts(email);

-- calculator_histories table
CREATE TABLE calculator_histories (
    id         SERIAL PRIMARY KEY,
    num1       DOUBLE PRECISION NOT NULL,
    num2       DOUBLE PRECISION NOT NULL,
    operation  VARCHAR(50) NOT NULL,
    result     DOUBLE PRECISION NOT NULL,
    version    INTEGER DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- agents table
CREATE TABLE agents (
    id          SERIAL PRIMARY KEY,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMP,
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    category    VARCHAR(100),
    version     VARCHAR(50),
    author      VARCHAR(255),
    repository  VARCHAR(500),
    tags        TEXT,
    status      VARCHAR(50) DEFAULT 'active',
    source_code TEXT
);

CREATE INDEX idx_agents_deleted_at ON agents(deleted_at);

-- tools table
CREATE TABLE tools (
    id          SERIAL PRIMARY KEY,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMP,
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    category    VARCHAR(100),
    version     VARCHAR(50),
    author      VARCHAR(255),
    repository  VARCHAR(500),
    tags        TEXT,
    language    VARCHAR(50),
    status      VARCHAR(50) DEFAULT 'active',
    source_code TEXT
);

CREATE INDEX idx_tools_deleted_at ON tools(deleted_at);
```

### GORM Connection Configuration

```go
// Connection Pool Settings
MaxIdleConns: 10   // Maximum idle connections
MaxOpenConns: 100  // Maximum open connections

// Connection String Format
host=%s port=%s user=%s password=%s dbname=%s sslmode=%s
```

---

## Request Flow

### Complete Request Flow Diagram

```
┌─────────────┐
│   Browser   │
└──────┬──────┘
       │ HTTP Request (GET/POST)
       ▼
┌─────────────────────────────────────────┐
│         Gin Router (main.go)            │
│  - Parse URL                            │
│  - Match route                          │
│  - Extract parameters                   │
└──────┬──────────────────────────────────┘
       │
       ▼
┌─────────────────────────────────────────┐
│      Handler Function                   │
│  (handlers/*.go)                        │
│  1. Extract form data / URL params      │
│  2. Validate input                      │
│  3. Business logic                      │
└──────┬──────────────────────────────────┘
       │
       ▼
┌─────────────────────────────────────────┐
│         GORM ORM Layer                  │
│  - Build SQL query                      │
│  - Execute query                        │
│  - Map results to structs               │
└──────┬──────────────────────────────────┘
       │
       ▼
┌─────────────────────────────────────────┐
│      PostgreSQL Database                │
│  - Execute SQL                          │
│  - Return results                       │
└──────┬──────────────────────────────────┘
       │
       ▼
┌─────────────────────────────────────────┐
│      Handler (Response)                 │
│  - Prepare data for template            │
│  - Render HTML template                 │
│  - OR Redirect to another page          │
└──────┬──────────────────────────────────┘
       │
       ▼
┌─────────────────────────────────────────┐
│      Gin Template Engine                │
│  - Load template file                   │
│  - Inject data                          │
│  - Generate HTML                        │
└──────┬──────────────────────────────────┘
       │
       ▼
┌─────────────┐
│   Browser   │ ← HTML Response
└─────────────┘
```

### Example: Calculator Flow

```
User submits calculation form
         │
         ▼
POST /calculator
         │
         ▼
CalculateResult() handler
         │
         ├─→ Extract num1, num2, operation from form
         │
         ├─→ Validate inputs (parseFloat, validateOperation)
         │
         ├─→ performCalculation(num1, num2, operation)
         │       │
         │       └─→ Returns result or error
         │
         ├─→ Create CalculatorHistory record
         │       │
         │       └─→ NewCalculatorHistory(num1, num2, op, result)
         │
         ├─→ Save to database
         │       │
         │       └─→ h.DB.Create(history)
         │
         └─→ Redirect to /calculator/history
                 │
                 ▼
         ListCalculatorHistory() handler
                 │
                 ├─→ Query all history records
                 │       │
                 │       └─→ h.DB.Order("created_at DESC").Find(&history)
                 │
                 └─→ Render calculator_history.html template
                         │
                         └─→ Display results to user
```

---

## Deployment Architecture

### Docker Deployment

```
┌─────────────────────────────────────────────────────────────┐
│                    Docker Host                               │
│                                                              │
│  ┌────────────────────┐      ┌─────────────────────────┐   │
│  │  gospace Container │      │  gospace_db Container   │   │
│  │  ┌──────────────┐  │      │  ┌──────────────────┐  │   │
│  │  │ Go App       │  │      │  │  PostgreSQL 15   │  │   │
│  │  │ Port: 8080   │  │◄────►│  │  Port: 5432      │  │   │
│  │  └──────────────┘  │      │  └──────────────────┘  │   │
│  └────────────────────┘      │  Volume: postgres_data │   │
│           │                  └─────────────────────────┘   │
│           │                                                 │
└───────────┼─────────────────────────────────────────────────┘
            │
            ▼
    Host Port 8080
            │
            ▼
    http://localhost:8080

Networks:
- gospace_default (bridge)
- mybridge (custom bridge for cluster access)
```

### Kubernetes Deployment

```
┌─────────────────────────────────────────────────────────────────┐
│                    Kubernetes Cluster (Kind)                     │
│                                                                  │
│  ┌──────────────────────────────────────────────────────────┐  │
│  │              Namespace: gospace-app                       │  │
│  │  ┌────────────────────────────────────────────────────┐  │  │
│  │  │         DaemonSet: gospace                         │  │  │
│  │  │  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────┐ │  │  │
│  │  │  │ Pod      │ │ Pod      │ │ Pod      │ │ Pod  │ │  │  │
│  │  │  │ worker1  │ │ worker2  │ │ worker3  │ │worker4│ │  │  │
│  │  │  └────┬─────┘ └────┬─────┘ └────┬─────┘ └───┬──┘ │  │  │
│  │  └───────┼────────────┼────────────┼────────────┼────┘  │  │
│  │          └────────────┴────────────┴────────────┘       │  │
│  │                         │                                │  │
│  │                         ▼                                │  │
│  │          ┌──────────────────────────────┐               │  │
│  │          │  Service: gospace-service    │               │  │
│  │          │  Type: NodePort              │               │  │
│  │          │  Port: 8080 → 30081          │               │  │
│  │          └──────────────────────────────┘               │  │
│  └──────────────────────────────────────────────────────────┘  │
│                                                                  │
│  ┌──────────────────────────────────────────────────────────┐  │
│  │              Namespace: gospace-db                        │  │
│  │  ┌────────────────────────────────────────────────────┐  │  │
│  │  │         StatefulSet: postgres                      │  │  │
│  │  │  ┌──────────────────────────────────────────────┐  │  │  │
│  │  │  │  Pod: postgres-0                             │  │  │  │
│  │  │  │  ┌────────────────────────────────────────┐  │  │  │  │
│  │  │  │  │  PostgreSQL 15                         │  │  │  │  │
│  │  │  │  │  Port: 5432                            │  │  │  │  │
│  │  │  │  └────────────────────────────────────────┘  │  │  │  │
│  │  │  └──────────────────────────────────────────────┘  │  │  │
│  │  │                         │                           │  │  │
│  │  │                         ▼                           │  │  │
│  │  │          ┌──────────────────────────────┐          │  │  │
│  │  │          │  Service: postgres-service   │          │  │  │
│  │  │          │  Type: ClusterIP             │          │  │  │
│  │  │          │  Port: 5432                  │          │  │  │
│  │  │          └──────────────────────────────┘          │  │  │
│  │  │                         │                           │  │  │
│  │  │                         ▼                           │  │  │
│  │  │          ┌──────────────────────────────┐          │  │  │
│  │  │          │  PVC: postgres-pvc           │          │  │  │
│  │  │          │  Size: 1Gi                   │          │  │  │
│  │  │          └──────────────────────────────┘          │  │  │
│  │  └────────────────────────────────────────────────────┘  │  │
│  └──────────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────┘
                         │
                         ▼
                 NodePort 30081
                         │
                         ▼
             http://localhost:30081
```

### Network Topology

```
┌─────────────────────────────────────────────────────────────┐
│                    mybridge Network                          │
│                    (172.20.0.0/16)                          │
│                                                              │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐     │
│  │  portainer   │  │  gospace_db  │  │   pgadmin    │     │
│  │  172.20.0.2  │  │  172.20.0.5  │  │  172.20.0.6  │     │
│  └──────────────┘  └──────────────┘  └──────────────┘     │
│                                                              │
│  ┌──────────────────────────────────────────────────────┐  │
│  │         mk-control-plane (Kind)                      │  │
│  │         172.20.0.10 (STATIC IP)                      │  │
│  │  ┌────────────────────────────────────────────────┐ │  │
│  │  │  Kubernetes Services (NodePort)                │ │  │
│  │  │  - GoSpace App: 30081                          │ │  │
│  │  │  - PostgreSQL: 30432                           │ │  │
│  │  │  - Portainer Agent: 30778                      │ │  │
│  │  └────────────────────────────────────────────────┘ │  │
│  └──────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
```

---

## Component Interaction Matrix

```
┌──────────────┬─────────┬─────────┬─────────┬──────────┬──────────┐
│ Component    │ Config  │ Models  │Handlers │ Database │Templates │
├──────────────┼─────────┼─────────┼─────────┼──────────┼──────────┤
│ main.go      │   ✓     │    ✓    │    ✓    │    ✓     │    ✓     │
│ config/*     │   -     │    -    │    -    │    ✓     │    -     │
│ models/*     │   -     │    -    │    -    │    -     │    -     │
│ handlers/*   │   -     │    ✓    │    -    │    ✓     │    ✓     │
│ templates/*  │   -     │    -    │    -    │    -     │    -     │
└──────────────┴─────────┴─────────┴─────────┴──────────┴──────────┘

Legend:
✓ = Direct dependency/interaction
- = No direct interaction
```

---

## Data Flow Summary

### Create Operation (e.g., Add Agent)
```
User → Form → POST /agents/add → CreateAgent() → 
GORM → PostgreSQL → Success → Redirect → ListAgents() → 
Template → HTML → User
```

### Read Operation (e.g., List Contacts)
```
User → GET /contacts → ListContacts() → GORM → 
PostgreSQL → Results → Template → HTML → User
```

### Update Operation (e.g., Edit Tool)
```
User → GET /tools/:id/edit → EditToolForm() → GORM → 
PostgreSQL → Load Data → Template → Form → User →
POST /tools/:id/edit → UpdateTool() → GORM → 
PostgreSQL → Update → Redirect → ListTools()
```

### Delete Operation (e.g., Delete Calculator History)
```
User → POST /calculator/history/:id/delete → 
DeleteCalculatorHistory() → GORM → PostgreSQL → 
Delete → Redirect → ListCalculatorHistory()
```

---

## Technology Stack Summary

```
┌─────────────────────────────────────────────────────────────┐
│                    Technology Stack                          │
├─────────────────────────────────────────────────────────────┤
│ Language:        Go 1.25                                    │
│ Web Framework:   Gin v1.12.0                                │
│ ORM:             GORM v1.31.1                               │
│ Database:        PostgreSQL 15                              │
│ DB Driver:       pgx/v5 (via GORM)                          │
│ Template Engine: Go html/template (via Gin)                 │
│ Containerization: Docker                                     │
│ Orchestration:   Kubernetes (Kind for local)                │
│ Testing:         testify v1.11.1                            │
└─────────────────────────────────────────────────────────────┘
```

---

**Document Version:** 1.0  
**Last Updated:** 2026-05-14  
**Author:** Bob (AI Assistant)  
**Status:** Complete Architecture Documentation ✅