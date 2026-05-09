# Todo CRUD REST API

A Go REST API for managing todo items with in-memory storage.

## Features

- Full CRUD operations (Create, Read, Update, Delete)
- In-memory storage using Go maps
- JSON serialization
- Basic middleware (logging, error handling, CORS)
- RESTful API design

## Project Structure

```
day2/
├── main.go           # Main entry point
├── handlers/         # HTTP handlers
├── models/           # Data models
├── middleware/       # Custom middleware
└── storage/          # In-memory storage
```

## API Endpoints

### Todos

| Method | Endpoint      | Description                    |
|--------|---------------|--------------------------------|
| GET    | /api/todos    | Get all todos                 |
| GET    | /api/todos/:id| Get a specific todo           |
| POST   | /api/todos    | Create a new todo             |
| PUT    | /api/todos/:id| Update an existing todo        |
| DELETE | /api/todos/:id| Delete a todo                 |
| GET    | /api/health   | Health check endpoint         |

## Usage

### Run the server

```bash
go run main.go
```

Server will start on `http://localhost:8080`

### API Examples

#### Create a todo
```bash
curl -X POST http://localhost:8080/api/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go","description":"Study Go programming","priority":1}'
```

#### Get all todos
```bash
curl http://localhost:8080/api/todos
```

#### Get a specific todo
```bash
curl http://localhost:8080/api/todos/1
```

#### Update a todo
```bash
curl -X PUT http://localhost:8080/api/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go Advanced","description":"Deep dive into Go","completed":true}'
```

#### Delete a todo
```bash
curl -X DELETE http://localhost:8080/api/todos/1
```

## Todo Model

```json
{
  "id": 1,
  "title": "Learn Go",
  "description": "Study Go programming",
  "completed": false,
  "priority": 1,
  "created_at": "2024-01-15T10:00:00Z",
  "updated_at": "2024-01-15T10:00:00Z"
}
```

## Middleware

- **Logger**: Logs all incoming requests with method, path, and response time
- **Error Handler**: Catches panics and returns proper error responses
- **CORS**: Enables Cross-Origin Resource Sharing
- **JSON Content-Type**: Sets Content-Type header for all responses
