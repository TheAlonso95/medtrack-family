# Authentication API Endpoints

This document outlines the authentication API endpoints that need to be implemented in the Go backend to support the frontend authentication features.

## Endpoints

### 1. User Registration

**Endpoint:** `POST /api/auth/signup`

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "securepassword",
  "name": "John Doe"
}
```

**Success Response (200 OK):**
```json
{
  "token": "jwt-token-here",
  "user": {
    "id": "user-id",
    "email": "user@example.com",
    "name": "John Doe"
  }
}
```

**Error Response (400 Bad Request):**
```json
{
  "message": "Email already in use"
}
```

### 2. User Login

**Endpoint:** `POST /api/auth/login`

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "securepassword"
}
```

**Success Response (200 OK):**
```json
{
  "token": "jwt-token-here",
  "user": {
    "id": "user-id",
    "email": "user@example.com",
    "name": "John Doe"
  }
}
```

**Error Response (401 Unauthorized):**
```json
{
  "message": "Invalid email or password"
}
```

### 3. Get Current User

**Endpoint:** `GET /api/auth/me`

**Headers:**
```
Authorization: Bearer jwt-token-here
```

**Success Response (200 OK):**
```json
{
  "id": "user-id",
  "email": "user@example.com",
  "name": "John Doe"
}
```

**Error Response (401 Unauthorized):**
```json
{
  "message": "Unauthorized"
}
```

## Implementation Notes

1. **JWT Token:**
   - The JWT token should be signed with a secure secret key.
   - The token should include the user ID and expiration time.
   - Token expiration should be set to 7 days.

2. **Password Security:**
   - Passwords should be hashed using bcrypt before storing in the database.
   - Minimum password length should be 6 characters.

3. **Database Schema:**
   - Users table should include at minimum: id, email, password (hashed), name, created_at, updated_at.
   - Email addresses should be unique.

4. **Error Handling:**
   - All endpoints should return appropriate HTTP status codes.
   - Error messages should be descriptive but not reveal sensitive information.

5. **CORS Configuration:**
   - The API should be configured to accept requests from the frontend domain.
   - Appropriate CORS headers should be set for all responses.