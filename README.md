# User DOB API – Go Backend Service

A production-style RESTful API built with **Go (Fiber)** and **PostgreSQL** that manages users and dynamically calculates their age based on date of birth.

This project focuses on **clean architecture**, **type-safe database access**, and **real-world backend best practices**, making it suitable for backend engineering interviews and assignments.

---

##  Key Highlights

- Clean layered architecture (Handler → Service → Repository)
- Type-safe SQL queries using **SQLC**
- PostgreSQL migrations with **golang-migrate**
- Dynamic age calculation using Go’s `time` package
- Proper HTTP status codes and validation
- Structured, scalable project layout

---

## 🛠️ Tech Stack

- **Go (Golang)**
- **GoFiber** – HTTP web framework
- **PostgreSQL** – Relational database
- **SQLC** – Compile-time safe SQL
- **golang-migrate** – Database migrations
- **go-playground/validator** – Request validation
- **Uber Zap** – Structured logging

---

##  Project Structure
cmd/server/main.go # Application entry point
config/ # Database configuration
db/
├── migrations/ # SQL migration files
└── sqlc/ # SQLC generated code
internal/
├── handler/ # HTTP handlers
├── service/ # Business logic
├── repository/ # Database access layer
├── routes/ # Route definitions
├── models/ # Domain models
└── logger/ # Logging setup

---

## 🗄️ Database Schema

**users**

| Column | Type | Description |
|------|------|------------|
| id | SERIAL | Primary key |
| name | TEXT | User name |
| dob | DATE | Date of birth |

---

## 🔌 API Endpoints

### ➕ Create User
**POST** `/users`

```json

{
  "name": "Alice",
  "dob": "1990-05-10"
}
```

## Other Endpoints

- **GET `/users/{id}`** – Fetch user details with dynamically calculated age  
- **GET `/users`** – List all users with age calculation

---

## How to Run the Project

### 
1. Clone the repository
```bash
git clone https://github.com/akshayhat/user-dob-api.git
cd user-dob-api
```
###
2. Create PostgreSQL database
``` sql
CREATE DATABASE user_dob_db;
```
3. Run migrations  
   migrate -path db/migrations -database "postgres://<user>:<password>@localhost:5432/user_dob_db?sslmode=disable" up

4. Start server  
   go run cmd/server/main.go

##  Why This Project

This project was built to demonstrate clean backend architecture, database correctness, and business logic implementation in Go.

The scope was intentionally kept focused to emphasize quality, maintainability, and real-world practices rather than excessive features.
