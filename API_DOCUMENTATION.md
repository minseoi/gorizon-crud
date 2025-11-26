# Gorizon CRUD API Documentation

RESTful API for ingredient management with MySQL database.

## Setup

### 1. Database Configuration

Copy the environment example file:
```bash
cp .env.example .env
```

Update the `.env` file with your MySQL credentials:
```env
DB_USER=root
DB_PASSWORD=your_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=gorizon_db
```

### 2. Create Database

Before running the application, create the database:
```sql
CREATE DATABASE gorizon_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 3. Run the Application

```bash
cd source
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### 1. Get All Ingredients (모든 식재료 조회)
**GET** `/ingredients`

Returns a list of all ingredients.

**Response (200 OK):**
```json
[
  {
    "ID": 1,
    "CreatedAt": "2024-11-26T10:00:00Z",
    "UpdatedAt": "2024-11-26T10:00:00Z",
    "DeletedAt": null,
    "name": "토마토"
  },
  {
    "ID": 2,
    "CreatedAt": "2024-11-26T10:00:00Z",
    "UpdatedAt": "2024-11-26T10:00:00Z",
    "DeletedAt": null,
    "name": "양파"
  }
]
```

**Example:**
```bash
curl http://localhost:8080/ingredients
```

---

### 2. Create Ingredient (식재료 추가)
**POST** `/ingredients`

Creates a new ingredient.

**Request Body:**
```json
{
  "name": "감자"
}
```

**Validation Rules:**
- `name`: Required, non-empty string

**Response (201 Created):**
```json
{
  "ID": 3,
  "CreatedAt": "2024-11-26T10:00:00Z",
  "UpdatedAt": "2024-11-26T10:00:00Z",
  "DeletedAt": null,
  "name": "감자"
}
```

**Response (400 Bad Request):**
```json
{
  "error": "Name is required"
}
```

**Example:**
```bash
curl -X POST http://localhost:8080/ingredients \
  -H "Content-Type: application/json" \
  -d '{"name":"감자"}'
```

---

### 3. Update Ingredient Name (식재료 이름 변경)
**PUT** `/ingredients/:id`

Updates the name of an existing ingredient.

**Request Body:**
```json
{
  "name": "고구마"
}
```

**Validation Rules:**
- `name`: Required, non-empty string

**Response (200 OK):**
```json
{
  "ID": 1,
  "CreatedAt": "2024-11-26T10:00:00Z",
  "UpdatedAt": "2024-11-26T10:30:00Z",
  "DeletedAt": null,
  "name": "고구마"
}
```

**Response (404 Not Found):**
```json
{
  "error": "Ingredient not found"
}
```

**Example:**
```bash
curl -X PUT http://localhost:8080/ingredients/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"고구마"}'
```

---

### 4. Delete Ingredient (식재료 제거)
**DELETE** `/ingredients/:id`

Deletes an ingredient by ID (soft delete using GORM).

**Response (200 OK):**
```json
{
  "message": "Ingredient deleted successfully"
}
```

**Response (404 Not Found):**
```json
{
  "error": "Ingredient not found"
}
```

**Example:**
```bash
curl -X DELETE http://localhost:8080/ingredients/1
```

---

## Error Responses

All error responses follow this format:

```json
{
  "error": "Error message description"
}
```

### Common HTTP Status Codes

- `200 OK`: Request successful
- `201 Created`: Resource created successfully
- `400 Bad Request`: Invalid request body or parameters
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server error

---

## Database Schema

### Ingredients Table

```sql
CREATE TABLE ingredients (
  id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  created_at DATETIME(3),
  updated_at DATETIME(3),
  deleted_at DATETIME(3),
  name VARCHAR(100) NOT NULL UNIQUE,
  INDEX idx_ingredients_deleted_at (deleted_at)
);
```

**Fields:**
- `id`: Auto-incrementing primary key
- `created_at`: Timestamp when record was created
- `updated_at`: Timestamp when record was last updated
- `deleted_at`: Soft delete timestamp (NULL if not deleted)
- `name`: Ingredient name (max 100 characters, unique)

---

## Features

- **CRUD Operations**: Complete Create, Read, Update, Delete functionality
- **MySQL Integration**: Uses GORM ORM for database operations
- **Soft Deletes**: Deleted records are marked, not removed
- **Auto Migration**: Database schema is automatically created
- **Input Validation**: Request validation with meaningful error messages
- **Environment Configuration**: Database settings via environment variables
- **Seed Data**: Initial test data is automatically created
- **Error Handling**: Comprehensive error handling with proper HTTP status codes

---

## Development

### Project Structure

```
source/
├── main.go                     # Application entry point
├── db/
│   └── db.go                  # Database connection and initialization
├── models/
│   └── ingredient.go          # Ingredient model with GORM tags
├── handlers/
│   └── ingredient_handler.go  # CRUD handler functions
└── routes/
    └── routes.go              # API route definitions
```

### Dependencies

- `github.com/labstack/echo/v4` - Web framework
- `gorm.io/gorm` - ORM library
- `gorm.io/driver/mysql` - MySQL driver for GORM

### Running Tests

```bash
# Example test commands
curl http://localhost:8080/ingredients
curl -X POST http://localhost:8080/ingredients -H "Content-Type: application/json" -d '{"name":"감자"}'
curl -X PUT http://localhost:8080/ingredients/1 -H "Content-Type: application/json" -d '{"name":"고구마"}'
curl -X DELETE http://localhost:8080/ingredients/1
```
