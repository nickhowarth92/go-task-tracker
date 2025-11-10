🧩 go-task-tracker

A simple Task Tracker API built in Go (Golang) using a clean MVC + Service + Repository architecture — inspired by Laravel’s structure.

This project is a learning experiment to explore Go for backend development, focusing on scalability, simplicity, and clean code organization.

⚙️ Tech Stack

Language: Go 1.25+

Frameworks/Libraries:

Gorilla Mux
 — Routing

GORM
 — ORM for database interaction

SQLite — Local development database

Architecture: MVC + Service + Repository (Laravel-style)

Platform: macOS (developed natively, no Docker required)

🏗️ Project Structure

go-task-tracker/
├── main.go # Entry point
├── go.mod / go.sum # Dependencies
│
├── /config/ # App configuration
│ └── db.go # Database connection (SQLite)
│
├── /models/ # Database models
│ ├── user.go # User model (has many Tasks)
│ └── task.go # Task model (belongs to User)
│
├── /repositories/ # Data access layer
│ └── task_repository.go # CRUD operations for Tasks
│
├── /services/ # Business logic layer
│ └── task_service.go # Orchestrates logic and validation
│
├── /handlers/ # Controller layer
│ └── task_handler.go # Handles HTTP requests/responses
│
└── /routes/ # API endpoints
└── routes.go # Route definitions (using Gorilla Mux)

🧱 Setup Instructions
1. Install Go

If you haven’t already:
brew install go

2. Clone the Repository

git clone https://github.com/YOURUSERNAME/go-task-tracker.git
cd go-task-tracker

3. Install Dependencies

go get github.com/gorilla/mux
go get gorm.io/gorm
go get gorm.io/driver/sqlite
go mod tidy

4. Run the Application

go run main.go

✅ You should see:
✅ Database connected
🚀 Server running at http://localhost:8080

🌐 API Endpoints
Method	Endpoint	Description
GET	/tasks	Get all tasks
POST	/tasks	Create a new task
PUT	/tasks/{id}	Update a task
DELETE	/tasks/{id}	Delete a task
GET	/users	Get all users (with tasks)
POST	/users	Create a user

💾 Example JSON (Task)

{
"title": "Learn Go",
"description": "Build a simple task tracker",
"completed": false,
"user_id": 1
}