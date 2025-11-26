# Gorizon CRUD API Documentation

## API Endpoints

### 1. 모든 식재료 조회 (Get All Ingredients)

모든 식재료 목록을 조회합니다.

**Endpoint**
```
GET /ingredients
```

**Request**
- Method: `GET`
- Headers: None required

**Response**

Success (200 OK):
```json
[
  {
    "ID": 1,
    "CreatedAt": "2024-01-01T00:00:00Z",
    "UpdatedAt": "2024-01-01T00:00:00Z",
    "DeletedAt": null,
    "name": "토마토"
  },
  {
    "ID": 2,
    "CreatedAt": "2024-01-01T00:00:00Z",
    "UpdatedAt": "2024-01-01T00:00:00Z",
    "DeletedAt": null,
    "name": "양파"
  }
]
```

---

### 2. 식재료 추가 (Create Ingredient)

새로운 식재료를 추가합니다.

**Endpoint**
```
POST /ingredients
```

**Request**
- Method: `POST`
- Headers:
  - `Content-Type: application/json`
- Body:
  ```json
  {
    "name": "사과"
  }
  ```

**Request Fields**
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| name | string | Yes | 식재료 이름 (최대 100자) |

**Response**

Success (201 Created):
```json
{
  "ID": 3,
  "CreatedAt": "2024-01-01T00:00:00Z",
  "UpdatedAt": "2024-01-01T00:00:00Z",
  "DeletedAt": null,
  "name": "사과"
}
```

Error (400 Bad Request):
```json
{
  "error": "Name is required"
}
```

---

### 3. 식재료 이름 변경 (Update Ingredient Name)

특정 식재료의 이름을 변경합니다.

**Endpoint**
```
PUT /ingredients/:id
```

**Path Parameters**
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | 식재료 ID |

**Request**
- Method: `PUT`
- Headers:
  - `Content-Type: application/json`
- Body:
  ```json
  {
    "name": "케챱"
  }
  ```

**Request Fields**
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| name | string | Yes | 변경할 식재료 이름 (최대 100자) |

**Response**

Success (200 OK):
```json
{
  "ID": 8,
  "CreatedAt": "2024-01-01T00:00:00Z",
  "UpdatedAt": "2024-01-01T12:00:00Z",
  "DeletedAt": null,
  "name": "케챱"
}
```

Error (404 Not Found):
```json
{
  "error": "Ingredient not found"
}
```

Error (400 Bad Request):
```json
{
  "error": "Name is required"
}
```

---

### 4. 식재료 제거 (Delete Ingredient)

특정 식재료를 삭제합니다 (소프트 삭제).

**Endpoint**
```
DELETE /ingredients/:id
```

**Path Parameters**
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | 삭제할 식재료 ID |

**Request**
- Method: `DELETE`
- Headers: None required

**Response**

Success (200 OK):
```json
{
  "message": "Ingredient deleted successfully"
}
```

Error (404 Not Found):
```json
{
  "error": "Ingredient not found"
}
```

---

## HTTP Status Codes

| Status Code | Description |
|-------------|-------------|
| 200 OK | 요청 성공 |
| 201 Created | 리소스 생성 성공 |
| 400 Bad Request | 잘못된 요청 (유효성 검증 실패) |
| 404 Not Found | 리소스를 찾을 수 없음 |
| 500 Internal Server Error | 서버 오류 |

## Error Response Format

모든 에러 응답은 다음 형식을 따릅니다:

```json
{
  "error": "Error message describing what went wrong"
}
```

## Data Model

### Ingredient

```go
{
  "ID": 1,                              // 자동 증가 Primary Key
  "CreatedAt": "2024-01-01T00:00:00Z",  // 생성 일시 (자동)
  "UpdatedAt": "2024-01-01T00:00:00Z",  // 수정 일시 (자동)
  "DeletedAt": null,                    // 삭제 일시 (소프트 삭제)
  "name": "토마토"                       // 식재료 이름 (UNIQUE, NOT NULL)
}
```

**Field Constraints**
- `name`: 최대 100자, 중복 불가, 필수 입력

## Notes

- 모든 삭제는 소프트 삭제로 처리됩니다 (레코드가 물리적으로 삭제되지 않고 `DeletedAt` 필드가 설정됨)
- 식재료 이름은 데이터베이스 레벨에서 중복이 허용되지 않습니다
- 모든 타임스탬프는 ISO 8601 형식으로 반환됩니다
