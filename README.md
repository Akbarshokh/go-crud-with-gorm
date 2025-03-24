# 🎬 Movie CRUD API

A RESTful movie management API built with:

- Go 1.23
- Gin
- GORM (PostgreSQL)
- Uber Fx (Dependency Injection)
- JWT Authentication
- Swagger docs
- Docker / Docker Compose

---

## 🧩 Features

- JWT-based authentication (`/auth/register`, `/auth/login`)
- Full CRUD for movies
- Dependency injection with UberFx
- GORM ORM integration with PostgreSQL
- Swagger documentation with Bearer Auth support
- Clean Architecture structure

---


## ⚙️ Getting Started

### 1. 📁 Setup `.env`

Copy `.env.example` and rename to `.env`, then update:

```env
DB_HOST=movie_postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=movie_db

JWT_SECRET=super-secret-key
HTTP_PORT=:8080
LOG_LEVEL=debug

```
---

## 2. 🐳 Run with Docker

```dockerfile
docker compose up --build
```

---

## 3. 📚 Swagger Docs

```
http://localhost:8080/swagger/index.html
```