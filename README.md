# Go Fiber Blog

Go Fiber Blog is a backend API starter built with Go, Fiber v3, GORM, and MySQL. The current codebase is intentionally small: it establishes the application bootstrap, database connection, model migration, API versioning, request logging, CORS, and a health endpoint. It is a good base for adding business features such as authentication, product management, user management, and product categories.

## Tech Stack

- Language: Go 1.25.5
- HTTP framework: Fiber v3
- Database ORM: GORM
- Database driver: MySQL (`gorm.io/driver/mysql` and `github.com/go-sql-driver/mysql`)
- Environment loading: `github.com/joho/godotenv`
- Hot reload for development: Air

## Current Project Structure

```text
.
|-- cmd/
|   `-- main.go
|-- config/
|   `-- config.go
|-- controller/
|   `-- frontendController.go
|-- database/
|   `-- database.go
|-- middleware/
|-- model/
|   |-- product.go
|   `-- user.go
|-- router/
|   `-- router.go
|-- .air.toml
|-- .env
|-- .env.example
|-- go.mod
`-- go.sum
```

## How The Application Works

1. `cmd/main.go` creates the Fiber app, connects to MySQL, enables CORS, registers routes, and starts the server on port `8000`.
2. `database/database.go` reads environment variables, builds the MySQL DSN, opens the GORM connection, and runs auto-migrations for the existing models.
3. `router/router.go` creates the API group under `/api/v1`, enables request logging for that group, and exposes the health check endpoint.
4. `controller/frontendController.go` currently contains the health check handler.
5. `model/product.go` and `model/user.go` define the current database tables.

## Existing API Surface

- `GET /api/v1/health`

Successful response:

```json
{
  "message": "Fiber is running ----->"
}
```

## Environment Variables

The code currently expects these variables:

```env
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=secret
DB_NAME=go_fiber_blog
SECRET=replace-me
```

Notes:

- `DB_HOST` is required by the database connection code.
- The server port is currently hardcoded to `8000` in `cmd/main.go`.
- `SECRET` exists in the template but is not used yet.

## Local Setup

### Prerequisites

- Go 1.25.x
- MySQL server running locally or remotely
- Air installed for hot reload:

```bash
go install github.com/air-verse/air@latest
```

### First Run

1. Copy `.env.example` to `.env` and fill in valid MySQL credentials.
2. Make sure the target database already exists.
3. Install dependencies if needed:

```bash
go mod tidy
```

4. Start the application with normal Go:

```bash
go run ./cmd
```

5. Or start it with hot reload:

```bash
air
```

The application will be available at `http://localhost:8000`.

## Development Workflow

### Run without hot reload

```bash
go run ./cmd
```

### Run with hot reload

```bash
air
```

### Verify the project builds

```bash
go build ./...
```

### Format code

```bash
gofmt -w ./cmd ./config ./controller ./database ./middleware ./model ./router
```

## Architecture Notes

- The project follows a simple layered structure: bootstrap, config, database, models, controllers, and routes.
- Routing is centralized in `router/router.go`.
- Database bootstrapping and migration are centralized in `database/database.go`.
- Models embed `gorm.Model`, so each table gets `ID`, timestamps, and soft delete support.
- Validation tags exist on the models, but request validation logic is not implemented yet.
- The `middleware` directory exists but is currently empty.

## Known Gaps

- `.env.example` originally missed `DB_HOST`; this should always be present for new developers.
- `config.Config()` reloads `.env` on every lookup. That works, but it is not efficient for larger applications.
- `database.ConnectDB()` prints the full DSN, which can expose credentials in logs.
- The server port is hardcoded instead of coming from configuration.
- Route handlers for products, users, and auth are stubbed as comments only.

## Adding A New Feature

Example: adding product categories.

1. Create a new model such as `model/category.go`.
2. Add the model to `database.ConnectDB()` auto-migration.
3. Create a controller file such as `controller/categoryController.go`.
4. Add routes in `router/router.go`, for example under `/api/v1/categories`.
5. If needed, add reusable request/auth logic under `middleware/`.
6. Add tests once handler and database logic are introduced.

Suggested category model starting point:

```go
package model

import "gorm.io/gorm"

type Category struct {
    gorm.Model
    Name        string `gorm:"size:100;not null;unique" json:"name"`
    Description string `gorm:"size:255" json:"description"`
}
```

Suggested route layout:

```text
/api/v1/categories
/api/v1/categories/:id
```

## Recommended Next Improvements

1. Move environment loading to application startup and cache config values.
2. Hide sensitive connection details from logs.
3. Add service and repository layers if the domain grows.
4. Introduce request DTOs and validation for create/update handlers.
5. Add unit and integration tests.
6. Make the app port configurable through `.env`.

## Quick Start For A New Developer

1. Read this file once.
2. Create `.env` from `.env.example`.
3. Start MySQL and create the database.
4. Run `air`.
5. Confirm `GET /api/v1/health` works.
6. Add your model, controller, and route in the same structure already used by the project.