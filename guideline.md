# AI Agent Guideline

## Project Identity

- Name: Go Fiber Blog
- Type: Go backend API starter
- Main purpose: provide a base for blog or commerce-style APIs using Fiber, GORM, and MySQL

## Core Stack

- Go 1.25.5
- Fiber v3 for HTTP routing and middleware
- GORM for ORM and schema migration
- MySQL as the database
- Godotenv for local environment loading
- Air for local hot reload

## Runtime Flow

1. `cmd/main.go` creates the Fiber app.
2. `database.ConnectDB()` opens the MySQL connection.
3. GORM auto-migrates `model.Product` and `model.User`.
4. `router.SetUpRoutes(app)` registers `/api/v1` routes.
5. The server listens on port `8000`.

## Current Code Layout

- `cmd/main.go`: application bootstrap
- `config/config.go`: reads environment variables from `.env`
- `database/database.go`: creates the GORM connection and runs migrations
- `controller/frontendController.go`: current health check controller
- `model/product.go`: product schema
- `model/user.go`: user schema
- `router/router.go`: API group and route registration
- `middleware/`: reserved for shared middleware, currently empty

## Existing Behavior

- API prefix is `/api/v1`
- Only active endpoint is `GET /api/v1/health`
- Request logging is enabled on the API group
- CORS is enabled globally

## Configuration Expectations

- Required env variables: `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`
- Optional or future-use env variable: `SECRET`
- `.env.example` should stay aligned with the code expectations
- Application port is hardcoded to `8000` unless refactored

## Development Conventions

- Keep the existing package split unless a broader refactor is requested
- Add new database models under `model/`
- Add new handlers under `controller/`
- Register routes centrally in `router/router.go`
- Add shared middleware under `middleware/`
- Keep changes small and consistent with the current code style

## Known Issues And Constraints

- `config.Config()` loads `.env` on every call
- `database.ConnectDB()` currently logs the DSN, which can expose credentials
- There is no service layer or repository layer yet
- Validation tags exist on models, but request validation is not wired into handlers
- The project has no tests yet

## Air / Hot Reload Notes

- Use `air` from the project root
- The Air config must remain cross-platform
- Prefer `build.entrypoint` and OS-specific overrides instead of hardcoded shell wrappers
- Do not reintroduce Windows-only `cmd /C` execution for the run command

## When Adding Features

For a new resource such as product category:

1. Add a model in `model/`
2. Include it in auto-migration
3. Add controller handlers
4. Register routes under `/api/v1`
5. Add validation and tests if the task includes behavior changes

## Safe Assumptions For Future Agents

- This repo is backend-only at the moment
- MySQL must be available before the app can start successfully
- Hot reload is expected to work through `air`
- Documentation should stay practical and onboarding-focused