# Gorizon CRUD

RESTful API server with MySQL database integration for ingredient management.
Built with Go, Echo framework, and GORM ORM.

---

## Features

- Complete CRUD operations (Create, Read, Update, Delete)
- MySQL database integration with GORM
- RESTful API design
- Input validation
- Error handling
- Auto-migration
- Soft deletes
- Environment-based configuration

---

## Quick Start

### Prerequisites

- Go 1.19 or higher
- MySQL 5.7 or higher
- Docker (optional)

### 1. Setup Database

Create a MySQL database:
```sql
CREATE DATABASE gorizon_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 2. Configure Environment

Copy the environment example:
```bash
cp .env.example .env
```

Edit `.env` with your MySQL credentials:
```env
DB_USER=root
DB_PASSWORD=your_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=gorizon_db
```

### 3. Install Dependencies

```bash
cd source
go mod download
```

### 4. Run the Server

```bash
go run main.go
```

The API server will start at `http://localhost:8080`

---

## API Documentation

See [API_DOCUMENTATION.md](API_DOCUMENTATION.md) for complete API reference.

### Quick API Examples

```bash
# Get all ingredients (모든 식재료 조회)
curl http://localhost:8080/ingredients

# Create ingredient (식재료 추가)
curl -X POST http://localhost:8080/ingredients \
  -H "Content-Type: application/json" \
  -d '{"name":"감자"}'

# Update ingredient name (식재료 이름 변경)
curl -X PUT http://localhost:8080/ingredients/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"고구마"}'

# Delete ingredient (식재료 제거)
curl -X DELETE http://localhost:8080/ingredients/1
```

---

## Docker Usage (Optional)

### 1. Run the Development Container

Execute the following script to start the container:
```bash
./run.sh
```

### 2. Access the Container

```bash
docker exec -it gorizon /bin/bash
```

### 3. Install Dependencies Inside Container

```bash
go mod download
```