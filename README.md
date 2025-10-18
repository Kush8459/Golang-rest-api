```
# Go RESTful API with SQLite

A simple and lightweight RESTful API built with Go (Golang) and SQLite.  
This project demonstrates how to perform basic CRUD (Create, Read, Update, Delete) operations using Go’s standard libraries and an SQLite backend.

## Features

- RESTful API endpoints (GET, POST, PUT, DELETE)  
- SQLite database for data persistence  
- Clean and modular project structure  
- JSON-based request and response handling  
- Error handling and response formatting  

## Tech Stack

- **Language:** Go (Golang)  
- **Database:** SQLite  
- **Libraries Used:**  
  - `net/http` – for HTTP server & routing  
  - `database/sql` – for database operations  
  - `github.com/mattn/go-sqlite3` – SQLite driver  

## Project Structure

```

├── cmd/
│   └── students-api           # entry point of the API
├── internal/                  # internal packages (handlers, models, database code)
├── .gitignore
├── go.mod
├── go.sum
└── README.md

````

## Setup & Installation

### 1. Clone the repository  
```bash
git clone https://github.com/Kush8459/Golang-rest-api.git
cd Golang-rest-api
````

### 2. Download dependencies

```bash
go mod tidy
```

### 3. Run the server

```bash
go run ./cmd/students-api
```

By default, your API server will run at:

```
http://localhost:8082
```

## API Endpoints

| Method | Endpoint         | Description                |
| ------ | ---------------- | -------------------------- |
| GET    | `/students`      | Get all students           |
| GET    | `/students/{id}` | Get student by ID          |
| POST   | `/students`      | Create a new student       |
| PUT    | `/students/{id}` | Update an existing student |
| DELETE | `/students/{id}` | Delete a student           |

> Replace “students” with your actual resource/entity if different.

## Example Request (POST)

**POST** `/students`

```json
{
  "name": "John Doe",
  "email": "abc@email.com",
  "age": 22
}
```

```
