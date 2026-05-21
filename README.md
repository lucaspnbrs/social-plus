# Social-Plus 🧵

> **A high-performance, production-ready forum platform built with Go — designed for scale, clarity, and clean architecture.**

Social-Plus is a text-first social forum where users can create posts, engage through reactions, filter and explore threads — powered by a robust Go backend, PostgreSQL, and a clean RESTful API. Built with software architecture best practices from the ground up.

---

## ✨ Features

- 📝 **Thread & Post System** — Create, edit, delete and visualize forum posts
- ❤️ **Reactions** — Like, unlike and interact with content
- 🔍 **Filter & Search** — Filter threads by category, date, popularity and author
- 📄 **Pagination** — Efficient cursor-based pagination for large datasets
- 🔐 **Authentication** — JWT-based auth with refresh token strategy
- 👤 **User Profiles** — Public profiles with post history and activity
- 🗂️ **Categories & Tags** — Organize threads for better discoverability

---

## 🏗️ Architecture & Technical Decisions

This project is not just functional — it is intentionally engineered. Every structural decision was made to maximize maintainability, testability, and scalability.

### 📁 Project Layout

The project follows a layered architecture with clear separation of concerns, adapted from the [golang-standards/project-layout](https://github.com/golang-standards/project-layout):

```
social-plus/
├── sql/                          # Raw SQL — migrations and seed scripts
│   ├── migrations/               # Versioned schema migrations (up/down)
│   └── seeds/                    # Optional seed data for dev/test
│
├── src/                          # All application source code
│   ├── auth/                     # JWT generation, validation, token refresh logic
│   ├── config/                   # Environment loading, app settings (viper/envconfig)
│   ├── controllers/              # HTTP handlers — thin layer, delegates to repositories
│   ├── db/                       # Database connection, pool config, pgx/sqlx setup
│   ├── middleware/               # Auth guard, CORS, rate limiter, request logger
│   ├── models/                   # Domain entities — pure Go structs, no framework deps
│   ├── repositories/             # Data access layer — all SQL queries live here
│   ├── responses/                # Standardized API response helpers (success/error)
│   ├── router/                   # Route declarations, versioning, middleware binding
│   └── security/                 # Password hashing, input sanitization, token signing
│
├── .env                          # Local environment variables (never committed)
├── .env.example                  # Safe template — committed to repo
├── .gitignore
├── go.mod
├── go.sum
└── main.go                       # Application entrypoint — wires everything together
```

> **Why this matters:** Each folder has a single, obvious responsibility. A new developer can navigate the codebase and know exactly where to find (or add) any piece of logic — without reading documentation first. This is the hallmark of intentional architecture.

---

### 🧩 Design Patterns Applied

#### 1. Repository Pattern
All database interactions are abstracted behind interfaces. Business logic **never** talks directly to PostgreSQL — it talks to a contract.

```go
// src/repositories/users
func (repository users ) Create( user models.User) (uint64, error) {
	statement, erro := repository.database.Prepare(
		"QUERY-THAT-USE-IN-YOUR-DB",
	)
	
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	var lastInserted uint64
	erro = statement.QueryRow(user.Nome, user.Nick, user.Email, user.Pass).Scan(&lastInserted)
	if erro != nil {
		return 0, erro
	}

	return lastInserted, nil

}
```

> **Benefit:** Swap PostgreSQL for any other database without touching a single line of business logic. Enables 100% mockable unit tests.

---

#### 2. Use Case Layer (Application Services)
Business rules live exclusively in `internal/usecase/`. Handlers are kept thin — they parse input, call use cases, return output.

```
HTTP Request → Handler (parse) → UseCase (logic) → Repository (data) → Response
```

> **Benefit:** Logic is framework-agnostic. The same use case can be called from an HTTP handler, a CLI command, or a gRPC server.

---

#### 3. Dependency Injection (Constructor Injection)
No global state. All dependencies are explicit and injected via constructors.

```go
func NewRepositoryFromUsers ( database *sql.DB) *users {
	return &users{database}
}
```

> **Benefit:** Testable, predictable, zero hidden coupling.

---

#### 4. Interface Segregation (Go Interfaces as Contracts)
Go's implicit interface system is used deliberately. Each layer depends only on the minimal interface it needs — not concrete types.

> **Benefit:** Follows the **I** in SOLID. Reduces tight coupling across layers.

---

#### 5. Middleware Chain Pattern
Cross-cutting concerns (auth, logging, CORS, rate limiting) are composed as middleware — fully decoupled from handlers.

```go
func Generate() *mux.Router {
	r := mux.NewRouter()

	return routers.Settings(r)
}
```

---

#### 6. DTO Pattern (Data Transfer Objects)
Domain entities are never exposed directly to the API layer. Dedicated request/response structs control what enters and exits the system.

```
domain.Post  ≠  dto.PostRequest  ≠  dto.PostResponse
```

> **Benefit:** Protects internal models from accidental exposure. Enables independent evolution of API contracts and domain models.

---

#### 7. Versioned Migrations
All schema changes are tracked as sequential SQL migration files, never modified retroactively.

```
sql/
├── sql.sql
└── ...
```

---

## 🛠️ Tech Stack | Future techs

| Layer | Technology |
|---|---|
| Language | Go 1.22+ |
| Web Framework | Chi / Gin / Fiber |
| Database | PostgreSQL 16 |
| ORM / Query Builder | sqlx + raw SQL or pgx |
| Auth | JWT (golang-jwt) |
| Config | envconfig / viper |
| Logging | zerolog |
| Containerization | Docker + Docker Compose |
| Migrations | golang-migrate |
| API Docs | Swagger (swaggo) |
| Testing | testify + gomock |

---

## 🚀 Getting Started

### Prerequisites

- Go 1.22+
- Docker & Docker Compose for containers 
- Make

### Running locally

```bash
# Clone the repo
git clone https://github.com/lucaspnbrs/social-plus.git
cd social-plus

# Copy environment config
cp .env.example .env

# Start services (PostgreSQL + API)
docker-compose up -d

# Run migrations
make migrate-up

# Start the API
make run
```

### Running tests

```bash
make test          # unit tests
make test-cover    # with coverage report
```

---

## 📐 API Overview

```
POST   /api/v1/auth/register
POST   /api/v1/auth/login
POST   /api/v1/auth/refresh

GET    /api/v1/posts              # list & filter posts
POST   /api/v1/posts              # create post
GET    /api/v1/posts/:id          # get single post
PUT    /api/v1/posts/:id          # update post
DELETE /api/v1/posts/:id          # delete post

POST   /api/v1/posts/:id/like     # react to post
DELETE /api/v1/posts/:id/like     # remove reaction

GET    /api/v1/users/:id          # public user profile
GET    /api/v1/users/:id/posts    # posts by user
```

---

## 🗺️ Roadmap

- [x] Core post CRUD
- [x] Authentication (JWT)
- [x] Reactions system
- [ ] Comments & nested threads
- [ ] Real-time notifications (WebSocket)
- [ ] Full-text search (PostgreSQL `tsvector`)
- [ ] Rate limiting per user
- [ ] Admin moderation panel
- [ ] Kubernetes deployment manifests

---

## 🤝 Contributing

Pull requests are welcome. For major changes, open an issue first to discuss what you'd like to change.

Please make sure to update tests accordingly.

---

## 📄 License

[MIT](LICENSE)

---

<p align="center">
  Built with precision by <a href="https://github.com/lucaspnbrs">Lucas Barros</a> · <strong>JL.Code</strong>
</p>
