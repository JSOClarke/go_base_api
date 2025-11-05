# Base CRUD API

A lightweight and opinionated starter template for building CRUD APIs in Go using the **Gin** framework.  
This template provides a clean project layout, middleware setup, JWT authentication, repository patterns, and ready-to-use test structure.

---

## Features

- Layered architecture (handlers, services, repositories, utils)
-  JWT authentication included
-  Environment variable support with `.env`
-  Unit and integration test setup
-  CI-ready test configuration (GitHub Actions)
-  Example CRUD endpoints (User)
-  Simple, extensible base for new Go APIs

---

## Project Structure

```
base-crud-api/
├── cmd/
│   └── main.go            # Entry point
├── internals/
│   ├── handlers/          # HTTP route handlers
│   ├── services/          # Business logic
│   ├── repository/        # Database interactions
│   ├── middleware/        # JWT, logging, etc.
│   ├── models/            # DTOs and data structs
│   └── utils/             # Helpers (JWT, env loading)
├── go.mod
├── .env
└── README.md
```

---

## Getting Started

### Clone the repo

```bash
git clone https://github.com/<your-username>/base-crud-api.git
cd base-crud-api
```

### Install dependencies

```bash
go mod tidy
```

### Set up environment

Create a `.env` file:

```
JWT_SECRET=your_secret_here
DB_URL=postgres://user:pass@localhost:5432/db
```

### Run the server

```bash
go run cmd/main.go
```

### Run tests

```bash
go test ./... -tags=unit -v
```

---

## Example Endpoints

| Method | Endpoint     | Description              |
| ------ | ------------ | ------------------------ |
| POST   | `/signup`    | Register a new user      |
| POST   | `/login`     | Authenticate and get JWT |


---

## Technologies Used

- [Go](https://go.dev/)
- [Gin Gonic](https://github.com/gin-gonic/gin)
- [GoDotEnv](https://github.com/joho/godotenv)
- [JWT-Go](https://github.com/golang-jwt/jwt)
- [Testify](https://github.com/stretchr/testify) for testing

---

## Future Improvements

- Docker setup for local DB testing  
- Swagger/OpenAPI auto-doc generation  
- Support for multiple database backends  

---

## License

MIT License © 2025 [Your Name]
