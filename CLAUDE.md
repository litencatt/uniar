# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

**Language Guidelines:**

- All content in CLAUDE.md should be written in English for consistency
- Chat responses should be in Japanese for user communication
- Think in English when analyzing problems and solutions

## Project Overview

`uniar` is a CLI tool and web server for managing UNI'S ON AIR scene card collections and database. It uses SQLite for data storage and is written in Go with Cobra for CLI commands and Gin for the web server.

## Development Commands

### Build
```bash
make build                    # Build the CLI binary
make air-cmd                  # Build with hot reload dependencies (for development)
make docker-build            # Build Docker image
```

### Docker Development
```bash
docker compose up -d          # Start development environment
docker compose down           # Stop and remove containers
docker compose restart app    # Restart application container

# NOTE: Container automatically starts with 'uniar server' command on startup
# Server will be accessible at http://localhost:8090
# Code changes are automatically reflected due to volume mounting (.:/app)
# No need to rebuild Docker image after code changes
docker compose exec app go run cmd/uniar/main.go [command]  # Run commands in container
```

### Database
```bash
make db-init                 # Remove the database file (reset)
make db-migrate             # Apply schema migrations using sqlite3def
make db-dump                # Export database to sql/seed.sql
make db-setup               # Initialize and migrate database (combines db-init + db-migrate)

# Docker database commands
make db-init-docker         # Initialize database in Docker container
make db-migrate-docker      # Apply schema migrations in Docker container
make db-setup-docker        # Initialize and migrate database in Docker container

sqlite3def ~/.uniar/uniar.db -f sql/schema.sql  # Apply schema manually
```

### Code Generation
```bash
sqlc generate               # Generate repository code from SQL queries
make gen-mock              # Generate mock interfaces for testing
```

### Development Server
```bash
air                        # Run with hot reload (uses .air.toml config)
./uniar server            # Run server directly
```

### Testing
```bash
go test ./...             # Run all tests
go test ./service/...     # Run service layer tests
go test -v ./...          # Run tests with verbose output
go test -race -coverprofile=coverage.out ./...  # Run tests with race detection and coverage

# Docker testing
docker compose exec app go test -v ./...  # Run tests in Docker container
```

### Linting
```bash
make lint                 # Run golangci-lint locally
make lint-docker          # Run golangci-lint in Docker container
golangci-lint run         # Run linter directly
```

### Documentation & Release
```bash
make doc                  # Generate README from command documentation
make prerelease          # Prepare for release (dump DB, update docs, generate changelog)
make release             # Create and push release
```

### CI/CD
```bash
# GitHub Actions workflows are configured in .github/workflows/
# - ci.yml: Runs tests, linting, and build checks on push/PR
# - Includes test coverage reporting with codecov
```

## Architecture

### Core Components

- **cmd/** - Cobra commands for CLI interface
  - Each command is in its own file (list_*.go, regist_*.go, setup_*.go)
  - Main entry point in cmd/uniar/main.go

- **repository/** - Database layer using sqlc
  - Auto-generated from SQL queries in sql/queries/
  - Mock interfaces generated for testing
  - Uses SQLite with modernc.org/sqlite driver (pure Go, CGO-free)

- **service/** - Business logic layer
  - Handles scene cards, members, photographs, producers
  - Orchestrates repository calls

- **handler/** - HTTP handlers for web server
  - Gin framework routes and middleware
  - OAuth2 integration for authentication

- **entity/** - Domain models
  - Scene, Member, Photograph, Producer entities
  - Includes CalcTotal method for score calculations

- **sql/** - Database schemas and queries
  - schema.sql defines table structure
  - queries/ contains sqlc query definitions
  - seed.sql contains initial data

### Database Schema

The application manages UNI'S ON AIR game data including:
- Scene cards with colors, photographs, members, and rankings
- Music and live performance data
- Member and group information
- Producer collections tracking owned cards

### Key Dependencies

- Cobra for CLI commands
- Gin for HTTP server
- sqlc for SQL code generation (with field rename configuration)
- modernc.org/sqlite as pure Go SQLite driver (CGO-free for cross-compilation)
- Air for hot reload during development
- mockgen for test mocks
- sqlmock for database testing
- golangci-lint for code quality checks

### sqlc Configuration

Field naming is configured in `sqlc.yaml` with rename rules:
```yaml
rename:
  memberid: "MemberID"
  photographid: "PhotographID"
  producerid: "ProducerID"
```

This ensures proper Go naming conventions (PascalCase) for generated struct fields.

## Templates Structure

### Template Directory Organization

Templates are organized in the `templates/` directory with the following structure:
```
templates/
├── admin/           # Admin panel templates
├── define/          # Template definitions and common components
├── error/           # Error page templates
├── members/         # Member-related templates
├── regist/          # Registration templates
├── scenes/          # Scene-related templates
└── top/             # Top page templates
```

### Template Loading Configuration

The server loads templates using the pattern `templates/**/*.go.tmpl` (two levels deep) in `cmd/server.go`:
```go
r.LoadHTMLGlob("templates/**/*.go.tmpl")
```

**IMPORTANT NOTES:**

1. **Template File Placement**: All template files MUST be placed in subdirectories under `templates/`, not directly in the `templates/` root directory.

2. **File Naming Convention**: Template files use the `.go.tmpl` extension and should follow the naming pattern `[prefix_]description.go.tmpl`.

3. **Server Restart for Template Changes**: After adding new templates or modifying existing ones, restart the server for changes to take effect. Binary rebuild is only required when Go source code (.go files) changes:
   ```bash
   # For template changes only - restart server
   docker compose restart app

   # For Go code changes - rebuild binary
   docker compose exec app go build -o uniar cmd/uniar/main.go
   ```

4. **Template Loading Issues**: If you encounter template loading errors like `pattern matches no files`, ensure:
   - All templates are in proper subdirectories (not root `templates/`)
   - The glob pattern matches your file structure
   - The server has been restarted after template changes

### Admin Templates

Admin panel templates are located in `templates/admin/` and include:
- `admin_dashboard.go.tmpl` - Admin dashboard
- `admin_music_*.go.tmpl` - Music management templates
- `admin_photograph_*.go.tmpl` - Photograph management templates
- `admin_scene_*.go.tmpl` - Scene management templates

Admin templates support CRUD operations with simplified HTTP methods (GET for read/forms, POST for all actions).