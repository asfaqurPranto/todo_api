# 🧾 Golang Todo List API

This project implements a **RESTful API** for managing a to-do list with **user authentication**, built using **Go (Golang)**.  
It includes endpoints for **user registration**, **login**, and **CRUD operations** for to-do items, ensuring a secure and user-friendly experience.  
This project is a completion of the **[Todo List API project on roadmap.sh](https://roadmap.sh/projects/todo-list-api)**.

---

##  Features

-  **User registration and login**
-  **CRUD operations** for to-do items
-  **User authentication and authorization** using JWT
-  **Refresh token** mechanism for authentication
-  **Error handling** and security measures
-  **Database integration using Docker**
-  Built with **Gin web framework**

---

## 📡 API Endpoints

### 👤 User Authentication

#### Register a New User

POST /register

**Request:**
```json
{
  "name": "Pranto",
  "email": "pranto@gmail.com",
  "password": "password"
}

Response:

{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
}

User Login

POST /login

Request:

{
  "email": "pranto@gmail.com",
  "password": "password"
}

Response:

{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
}

🧾 Todo Management
Create a To-Do Item

POST /todos

Request:

{
  "completed": false,
  "title": "Buy groceries",
  "description": "Buy milk, eggs, and bread"
}

Response:

{
  "id": 1,
  "completed": false,
  "title": "Buy groceries",
  "description": "Buy milk, eggs, and bread"
}

Update a To-Do Item

PUT /todos/1

Request:

{
  "completed": true,
  "title": "Buy groceries",
  "description": "Buy milk, eggs, bread, and cheese"
}

Response:

{
  "id": 1,
  "completed": true,
  "title": "Buy groceries",
  "description": "Buy milk, eggs, bread, and cheese"
}

Delete a To-Do Item

DELETE /todos/1

Response:
204 No Content
Get To-Do Items 

Project Directory Structure

todo-api
├── cmd
│   └── app
│       └── main.go              # Application entry point
├── internal
│   ├── handlers                 # HTTP request handlers (controllers)
│   │   ├── user_handler.go
│   │   └── todo_handler.go
│   ├── middleware               # JWT auth middleware
│   ├── services                 # Business logic (service layer)
│   ├── models                   # Database models (User, Todo)
│   └── config                   # Database configuration
