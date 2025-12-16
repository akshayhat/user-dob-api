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

## 📁 Project Structure
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

