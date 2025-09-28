# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`uniar` is a CLI tool and web server for managing UNI'S ON AIR scene card collections and database. It uses SQLite for data storage and is written in Go with Cobra for CLI commands and Gin for the web server.

## Development Commands

### Build
```bash
make build                    # Build the CLI binary
make air-cmd                  # Build with hot reload dependencies (for development)
make docker-build            # Build Docker image
```

### Database
```bash
make db-init                 # Remove the database file (reset)
make db-migrate             # Apply schema migrations using sqlite3def
make db-dump                # Export database to sql/seed.sql
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
```

### Documentation & Release
```bash
make doc                  # Generate README from command documentation
make prerelease          # Prepare for release (dump DB, update docs, generate changelog)
make release             # Create and push release
```

## Architecture

### Core Components

- **cmd/** - Cobra commands for CLI interface
  - Each command is in its own file (list_*.go, regist_*.go, setup_*.go)
  - Main entry point in cmd/uniar/main.go

- **repository/** - Database layer using sqlc
  - Auto-generated from SQL queries in sql/queries/
  - Mock interfaces generated for testing
  - Uses SQLite with go-sqlite3 driver

- **service/** - Business logic layer
  - Handles scene cards, members, photographs, producers
  - Orchestrates repository calls

- **handler/** - HTTP handlers for web server
  - Gin framework routes and middleware
  - OAuth2 integration for authentication

- **entity/** - Domain models
  - Scene, Member, Photograph, Producer entities

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
- sqlc for SQL code generation
- sqlite3 as database
- Air for hot reload during development
- mockgen for test mocks