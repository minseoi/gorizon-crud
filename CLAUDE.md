# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Gorizon CRUD is a RESTful API server for ingredient management built with Go, Echo framework, and GORM ORM. The application uses MySQL for persistence and follows a clean architecture pattern with separated concerns.

## Development Commands

### Setup and Dependencies
```bash
cd source
go mod download          # Install dependencies
cp ../.env.example ../.env  # Create environment configuration
```

### Running the Application
```bash
cd source
go run main.go           # Run development server (starts on :8080)
go build -o gorizon-api main.go  # Build production binary
./gorizon-api            # Run production binary
```

### Testing API Endpoints
Use the REST client file for testing:
- File: `doc/api.rest`
- Use VS Code REST Client extension or similar tools
- All endpoints are pre-configured with examples

### Manual Testing
```bash
# Test all ingredients endpoint
curl http://localhost:8080/ingredients

# Test create ingredient
curl -X POST http://localhost:8080/ingredients \
  -H "Content-Type: application/json" \
  -d '{"name":"토마토"}'
```

## Architecture

### Project Structure
```
source/
├── main.go                     # Entry point: Echo initialization + routing
├── db/db.go                    # Database: connection, migration, seeding
├── models/ingredient.go        # Data model with GORM tags
├── handlers/ingredient_handler.go  # HTTP handlers for CRUD operations
└── routes/routes.go            # Route registration
```

### Key Architecture Patterns

**Layered Architecture:**
- `main.go` → Application bootstrap
- `routes/` → HTTP routing configuration
- `handlers/` → Request/response handling + validation
- `models/` → Data structures with GORM annotations
- `db/` → Database connection + schema management

**Database Pattern:**
- GORM manages all database operations
- Auto-migration on startup creates/updates schema
- Soft deletes enabled (records marked as deleted, not removed)
- Seed data automatically inserted if table is empty

**Configuration:**
- Environment variables loaded from `.env` file
- Database credentials: `DB_USER`, `DB_PASSWORD`, `DB_HOST`, `DB_PORT`, `DB_NAME`
- Default values provided in `db/getEnv()` function

### API Design

**Current Endpoints (Ingredient CRUD):**
- `GET /ingredients` - List all ingredients
- `POST /ingredients` - Create new ingredient (requires: `name`)
- `PUT /ingredients/:id` - Update ingredient name (requires: `name`)
- `DELETE /ingredients/:id` - Soft delete ingredient

**Response Patterns:**
- Success: HTTP 200/201 with data
- Not Found: HTTP 404 with `{"error": "message"}`
- Bad Request: HTTP 400 with `{"error": "message"}`
- Server Error: HTTP 500 with `{"error": "message"}`

### Database Schema

**Ingredient Model:**
```go
type Ingredient struct {
    gorm.Model              // ID, CreatedAt, UpdatedAt, DeletedAt
    Name string             // VARCHAR(100), NOT NULL, UNIQUE
}
```

The `gorm.Model` embeds standard fields:
- `ID`: Auto-incrementing primary key
- `CreatedAt`, `UpdatedAt`: Automatic timestamps
- `DeletedAt`: Soft delete support

## Important Implementation Details

### Adding New Models
1. Create model in `source/models/` with `gorm.Model` embedded
2. Add GORM tags for validation and constraints
3. Update `db/db.go` to include new model in `AutoMigrate()`
4. Create handlers in `source/handlers/`
5. Register routes in `source/routes/routes.go`

### Database Connection
- Connection established in `db.Initialize()`
- Called from `main.go` before route registration
- Fatal error if connection fails (application won't start)
- Uses connection pooling via GORM defaults

### Validation Pattern
- Validation happens in handlers before database operations
- Empty string checks for required fields
- Return HTTP 400 with descriptive error messages
- Example: `if newIngredient.Name == "" { return error }`

### Error Handling
- Database errors return HTTP 500
- Not found errors return HTTP 404
- Validation errors return HTTP 400
- All errors include JSON `{"error": "message"}` response

## Environment Setup

### Required Environment Variables
```bash
DB_USER=root              # MySQL username
DB_PASSWORD=password      # MySQL password
DB_HOST=localhost         # Database host
DB_PORT=3306             # MySQL port
DB_NAME=gorizon_db       # Database name
```

### Database Initialization
```sql
CREATE DATABASE gorizon_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

## Deployment Notes

### Production Build
```bash
cd source
go build -o gorizon-api main.go
```

### Systemd Service (Linux)
The application can be run as a systemd service. Binary runs on port 8080 by default. Ensure MySQL is running and accessible before starting the service.

### Docker Alternative
Use `./run.sh` script for containerized development environment. Access container with `docker exec -it gorizon /bin/bash`.

## Code Conventions

### Import Organization
Standard library → Third-party → Local packages:
```go
import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/minseoi/gorizon/db"
)
```

### Handler Function Pattern
All handlers follow this pattern:
1. Extract/bind request data
2. Validate input
3. Perform database operation
4. Return JSON response with appropriate HTTP status

### Model Naming
- Struct names: Singular (e.g., `Ingredient`)
- Table names: Plural, auto-generated by GORM (e.g., `ingredients`)
- JSON fields: lowercase (controlled by `json` tags)
