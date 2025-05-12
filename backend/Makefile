# Backend Makefile

# Variables
APP_NAME = backend-server
MAIN_PATH = ./cmd/server
BUILD_DIR = ./build
MIGRATION_DIR = ./migrations
DB_URL = "postgres://postgres:postgres@localhost:5432/app_db?sslmode=disable"

# Colors for terminal output
GREEN = \033[0;32m
YELLOW = \033[0;33m
RED = \033[0;31m
NC = \033[0m # No Color

.PHONY: help
help: ## Display this help message
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

# Server commands
.PHONY: run
run: ## Run the server
	@echo "$(GREEN)Starting server...$(NC)"
	@go run $(MAIN_PATH)/main.go

.PHONY: build
build: ## Build the server binary
	@echo "$(GREEN)Building server...$(NC)"
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)/main.go
	@echo "$(GREEN)Build complete: $(BUILD_DIR)/$(APP_NAME)$(NC)"

# Database migration commands
.PHONY: migrate-create
migrate-create: ## Create a new migration file (usage: make migrate-create name=migration_name)
	@if [ -z "$(name)" ]; then \
		echo "$(RED)Error: migration name is required$(NC)"; \
		echo "Usage: make migrate-create name=migration_name"; \
		exit 1; \
	fi
	@mkdir -p $(MIGRATION_DIR)
	@echo "$(GREEN)Creating migration $(name)...$(NC)"
	@migrate create -ext sql -dir $(MIGRATION_DIR) -seq $(name)

.PHONY: migrate-up
migrate-up: ## Run all pending migrations
	@echo "$(GREEN)Running migrations up...$(NC)"
	@migrate -path $(MIGRATION_DIR) -database $(DB_URL) up

.PHONY: migrate-down
migrate-down: ## Rollback the last migration
	@echo "$(GREEN)Rolling back the last migration...$(NC)"
	@migrate -path $(MIGRATION_DIR) -database $(DB_URL) down 1

.PHONY: migrate-reset
migrate-reset: ## Rollback all migrations
	@echo "$(GREEN)Rolling back all migrations...$(NC)"
	@migrate -path $(MIGRATION_DIR) -database $(DB_URL) down

.PHONY: migrate-force
migrate-force: ## Force migration version (usage: make migrate-force version=x)
	@if [ -z "$(version)" ]; then \
		echo "$(RED)Error: version is required$(NC)"; \
		echo "Usage: make migrate-force version=x"; \
		exit 1; \
	fi
	@echo "$(GREEN)Forcing migration version to $(version)...$(NC)"
	@migrate -path $(MIGRATION_DIR) -database $(DB_URL) force $(version)

.PHONY: migrate-status
migrate-status: ## Check migration status
	@echo "$(GREEN)Checking migration status...$(NC)"
	@migrate -path $(MIGRATION_DIR) -database $(DB_URL) version

# Database commands
.PHONY: db-create
db-create: ## Create the database
	@echo "$(GREEN)Creating database...$(NC)"
	@echo "This is a placeholder. Implement database creation logic here."
	# Example: psql -U postgres -c "CREATE DATABASE app_db;"

.PHONY: db-drop
db-drop: ## Drop the database
	@echo "$(GREEN)Dropping database...$(NC)"
	@echo "This is a placeholder. Implement database drop logic here."
	# Example: psql -U postgres -c "DROP DATABASE IF EXISTS app_db;"

.PHONY: db-reset
db-reset: ## Reset the database (drop, create, migrate, seed)
	@echo "$(GREEN)Resetting database...$(NC)"
	@make db-drop
	@make db-create
	@make migrate-up
	@make db-seed
	@echo "$(GREEN)Database reset complete.$(NC)"

.PHONY: db-seed
db-seed: ## Seed the database with initial data
	@echo "$(GREEN)Seeding database...$(NC)"
	@echo "This is a placeholder. Implement database seeding logic here."
	# Example: go run ./cmd/seed/main.go

# Testing commands
.PHONY: test
test: ## Run tests
	@echo "$(GREEN)Running tests...$(NC)"
	@go test -v ./...

.PHONY: test-coverage
test-coverage: ## Run tests with coverage
	@echo "$(GREEN)Running tests with coverage...$(NC)"
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)Coverage report generated: coverage.html$(NC)"

# Linting and formatting
.PHONY: lint
lint: ## Run linter
	@echo "$(GREEN)Running linter...$(NC)"
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run ./...; \
	else \
		echo "$(YELLOW)golangci-lint not installed. Run: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest$(NC)"; \
		exit 1; \
	fi

.PHONY: fmt
fmt: ## Format code
	@echo "$(GREEN)Formatting code...$(NC)"
	@go fmt ./...

# Cleanup
.PHONY: clean
clean: ## Clean build artifacts
	@echo "$(GREEN)Cleaning build artifacts...$(NC)"
	@rm -rf $(BUILD_DIR)
	@go clean
	@echo "$(GREEN)Cleanup complete.$(NC)"

# Default target
.DEFAULT_GOAL := help