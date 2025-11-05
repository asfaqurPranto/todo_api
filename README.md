# 🧾 Golang Todo List API

This project implements a **RESTful API** for managing a to-do list with **user authentication**, built using **Go (Golang)**.  
It includes endpoints for **user registration**, **login**, and **CRUD operations** for to-do items, ensuring a secure and user-friendly experience.  
This project is a completion of the **[Todo List API project on roadmap.sh](https://roadmap.sh/projects/todo-list-api)**.

---

## Features

- **User registration and login**
- **CRUD operations** for to-do items
- **User authentication and authorization** using JWT
- **Refresh token** mechanism for authentication
- **Error handling** and security measures
- **Database integration using Docker**
- Built with **Gin web framework**

---

## 📡 API Endpoints

### 👤 User Authentication

todo-api
├── cmd
│   └── app
│       └── main.go          # Application entry point
├── internal
│   ├── handlers             # HTTP request handlers (controllers)
│   │   ├── user_handler.go
│   │   └── todo_handler.go
│   ├── middleware           # JWT auth middleware
│   ├── services             # Business logic (service layer)
│   ├── models               # Database models (User, Todo)
│   └── config               # Database configuration
