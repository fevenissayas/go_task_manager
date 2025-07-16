# Task Management API Documentation

Base URL: `http://localhost:8080`

## Endpoints

---

### Create Task

**POST** `/tasks`

**Request Body:**
```
{
  "id": "1",
  "title": "buy fruits",
  "description": "apple, banana, avocado",
  "due_date": "2025-07-25T10:00:00Z",
  "status": "pending"
}
```
**Responses**
- 201 Created: Task created successfully
- 400 Bad Request: Invalid JSON or duplicate ID

### Get All Tasks

**GET** `/tasks`

**Responses**

- 200 OK: List of tasks
- 200 OK + message: "No Task in DB" if list is empty

### 🔹 Get Task by ID

**GET** `/tasks/:id`

**Responses**

- 200 OK: Task details
- 404 Not Found: Task does not exist

### Update Task
**PUT** `/tasks/:id`

**Request Body**

```
{
  "title": "Updated Title",
  "description": "Updated Description",
  "due_date": "2025-08-01T08:00:00Z",
  "status": "completed"
}
```

**Responses**

- 200 OK: Task updated
- 400 Bad Request: Invalid JSON
- 404 Not Found: Task not found

### Delete Task

**DELETE** `/tasks/:id`

**Responses**

- 200 OK: "Task deleted Successfully"
- 404 Not Found: Task not found

### Testing Tools
**You can test the endpoints using:**
- Postman (recommended)
- VS Code REST Client (.http or .rest file)
- curl / Invoke-RestMethod