# Backend Service

This is the backend service for the application.

## Prerequisites

- Go 1.21 or higher
- PostgreSQL
- [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) for database migrations

## Getting Started

### Installing Dependencies

```bash
go mod download
```

### Running the Server

```bash
make run
```

The server will start on port 8080.

## Database Operations

### Setting Up the Database

```bash
# Create the database
make db-create

# Run migrations
make migrate-up

# Seed the database with initial data
make db-seed
```

### Database Migrations

```bash
# Create a new migration
make migrate-create name=migration_name

# Apply all pending migrations
make migrate-up

# Rollback the last migration
make migrate-down

# Check migration status
make migrate-status

# Reset the database (drop, create, migrate, seed)
make db-reset
```

## Development

### Building the Application

```bash
make build
```

The binary will be created in the `build` directory.

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage
```

### Code Quality

```bash
# Format code
make fmt

# Run linter
make lint
```

### Cleaning Up

```bash
make clean
```

## Available Make Commands

Run `make help` to see all available commands:

```
Usage: make [target]

Targets:
  build           Build the server binary
  clean           Clean build artifacts
  db-create       Create the database
  db-drop         Drop the database
  db-reset        Reset the database (drop, create, migrate, seed)
  db-seed         Seed the database with initial data
  fmt             Format code
  help            Display this help message
  lint            Run linter
  migrate-create  Create a new migration file (usage: make migrate-create name=migration_name)
  migrate-down    Rollback the last migration
  migrate-force   Force migration version (usage: make migrate-force version=x)
  migrate-reset   Rollback all migrations
  migrate-status  Check migration status
  migrate-up      Run all pending migrations
  run             Run the server
  test            Run tests
  test-coverage   Run tests with coverage
```