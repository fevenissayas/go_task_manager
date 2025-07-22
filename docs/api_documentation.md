# Task Management API Documentation

Base URL: `http://localhost:8080`

## Overview

A RESTful API for user registration, authentication, user roles, and task management with JWT authentication.

## Authentication

All protected routes require a JWT in the header:
```
Authorization: Bearer <your_token_here>
```

## User Roles

- **Admin:** Full access (create, update, delete tasks, promote users)
- **User:** Can view tasks only

---

## User Endpoints

### Register
`POST /register`

**Request Body:**
```json
{
  "username": "user",
  "password": "pass"
}
```
- First user becomes admin
- **Responses:**  
  `201`: Success  
  `400`: User exists

---

### Login
`POST /login`

**Request Body:**
```json
{
  "username": "user",
  "password": "pass"
}
```
- Returns JWT token  
- **Responses:**  
  `401`: Invalid credentials

---

### Promote User
`POST /promote` (Admin only)

**Request Body:**
```json
{
  "user_id": "<id>",
  "new_role": "admin"
}
```
- **Responses:**  
  `200`: Success  
  `403`: Admin only  
  `404`: Not found

---

## Task Endpoints

> **Note:** JWT required for all.  
> - Admins: create, update, delete  
> - All users: view

---

### Create Task
`POST /tasks`

**Request Body:**
```json
{
  "id": "1",
  "title": "buy fruits",
  "description": "apple, banana, avocado",
  "due_date": "2025-07-25T10:00:00Z",
  "status": "pending"
}
```
**Responses:**  
`201`: Task created  
`500`: Internal Server Error

---

### Get All Tasks
`GET /tasks`

**Responses:**  
`200`: List of tasks  
`200`: Message "No Task in DB" if list is empty  
`500`: Internal Server Error

---

### Get Task by ID
`GET /tasks/:id`

**Responses:**  
`200`: Task details  
`500`: Internal Server Error

---

### Update Task
`PUT /tasks/:id`

**Request Body:**
```json
{
  "title": "Updated Title",
  "description": "Updated Description",
  "due_date": "2025-08-01T08:00:00Z",
  "status": "completed"
}
```
**Responses:**  
`200`: Task updated  
`400`: Invalid JSON  
`404`: Task not found

---

### Delete Task
`DELETE /tasks/:id`

**Responses:**  
`200`: "Task deleted Successfully"  
`404`: Task not found

---

## Testing Tools

You can test the endpoints using:
- Postman (recommended)
- VS Code REST Client (.http or .rest file)
- curl / Invoke-RestMethod
