# Task Management API (Go + Gin)

## Features

- Create, read, update, and delete tasks  
- Retrieve all tasks or a specific task by ID  
- No database setup required  
- Clean and modular codebase

## Project Structure

go_task_manager/
├── main.go # Application entry point
├── controllers/ # HTTP request handlers
│ └── task_controller.go
├── models/ # Task data model
│ └── task.go
├── data/ # Business logic and in-memory storage
│ └── task_service.go
├── router/ # API route definitions
│ └── router.go
├── docs/ # API documentation
│ └── api_documentation.md
└── go.mod # Go module file


## Getting Started

### Prerequisites

- Go installed (version 1.16 or higher)

### Run the Application

```bash
git clone https://github.com/your-username/task_manager.git
cd task_manager
go run main.go
```

## Documentation
For details on endpoints, request/response formats, and example usage, see:
docs/api_documentation.md
