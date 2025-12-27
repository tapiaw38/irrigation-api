.PHONY: help install-deps run run-dev test cover cover-html migrate-up migrate-down migrate-create init-docker clean

# Colors for terminal output
GREEN  := \033[0;32m
YELLOW := \033[0;33m
BLUE   := \033[0;34m
RESET  := \033[0m

# Database configuration
DB_URL := postgres://postgres:postgres123@localhost:5432/irrigation_db?sslmode=disable
MIGRATIONS_PATH := database/migrations

help: ## Display this help message
	@echo "$(BLUE)Irrigation API - Development Commands$(RESET)"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "$(GREEN)%-18s$(RESET) %s\n", $$1, $$2}'

welcome:
	@echo "$(BLUE)========================================$(RESET)"
	@echo "$(BLUE)  Irrigation API Development$(RESET)"
	@echo "$(BLUE)========================================$(RESET)"
	@echo ""

install-deps: welcome ## Install development dependencies (Air, Delve, migrate)
	@echo "$(YELLOW)Installing Delve (Go debugger)...$(RESET)"
	@go install github.com/go-delve/delve/cmd/dlv@latest
	@echo "$(YELLOW)Installing Air (live reloader)...$(RESET)"
	@go install github.com/air-verse/air@latest
	@echo "$(GREEN)✓ All tools installed successfully$(RESET)"
	@echo ""
	@echo "$(YELLOW)To install migrate CLI:$(RESET)"
	@echo "  Arch Linux:   make pacman-install-migrate"
	@echo "  Debian/Ubuntu: make apt-install-migrate"

apt-install-migrate: ## Install migrate tool (Debian/Ubuntu)
	@echo "$(YELLOW)Installing migrate CLI for Debian/Ubuntu...$(RESET)"
	@curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-amd64.tar.gz | tar xvz
	@sudo mv migrate /usr/local/bin/migrate
	@rm LICENSE README.md
	@echo "$(GREEN)✓ migrate installed successfully$(RESET)"

pacman-install-migrate: ## Install migrate tool (Arch Linux)
	@echo "$(YELLOW)Installing migrate CLI for Arch Linux...$(RESET)"
	@yay -S golang-migrate-bin
	@echo "$(GREEN)✓ migrate installed successfully$(RESET)"

run: welcome ## Run the application directly (without hot reload)
	@echo "$(YELLOW)Starting application...$(RESET)"
	@go run ./main.go

run-dev: welcome ## Run the application with Air hot reload
	@echo "$(YELLOW)Starting Air (hot reload enabled)...$(RESET)"
	@air

test: welcome ## Run all tests with coverage
	@echo "$(YELLOW)Running tests...$(RESET)"
	@go test ./... -cover

cover: welcome ## Show test coverage in function mode
	@echo "$(YELLOW)Generating coverage report...$(RESET)"
	@go test ./... -coverprofile=coverage.out
	@go tool cover -func=coverage.out

cover-html: welcome ## Generate HTML coverage report
	@echo "$(YELLOW)Generating HTML coverage report...$(RESET)"
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)✓ Coverage report generated: coverage.html$(RESET)"

migrate-up: welcome ## Apply all database migrations
	@echo "$(YELLOW)Applying migrations...$(RESET)"
	@migrate -path $(MIGRATIONS_PATH) -database '$(DB_URL)' up
	@echo "$(GREEN)✓ Migrations applied successfully$(RESET)"

migrate-down: welcome ## Revert the last database migration
	@echo "$(YELLOW)Reverting last migration...$(RESET)"
	@migrate -path $(MIGRATIONS_PATH) -database '$(DB_URL)' down 1
	@echo "$(GREEN)✓ Migration reverted successfully$(RESET)"

migrate-create: welcome ## Create a new migration (prompts for name)
	@read -p "Enter migration name: " name; \
	migrate create -dir $(MIGRATIONS_PATH) -ext sql $$name
	@echo "$(GREEN)✓ Migration files created$(RESET)"

init-docker: welcome ## Start PostgreSQL and Redis services
	@echo "$(YELLOW)Starting Docker services (PostgreSQL & Redis)...$(RESET)"
	@docker-compose up -d
	@echo "$(GREEN)✓ Services started$(RESET)"
	@echo ""
	@echo "$(BLUE)Services running:$(RESET)"
	@echo "  PostgreSQL: localhost:5432"
	@echo "  Redis:      localhost:6380"

stop-docker: welcome ## Stop all Docker services
	@echo "$(YELLOW)Stopping Docker services...$(RESET)"
	@docker-compose down
	@echo "$(GREEN)✓ Services stopped$(RESET)"

clean: welcome ## Clean build artifacts and cache
	@echo "$(YELLOW)Cleaning build artifacts...$(RESET)"
	@rm -rf tmp/
	@rm -f coverage.out coverage.html
	@rm -f build-errors.log
	@go clean -cache
	@echo "$(GREEN)✓ Cleanup complete$(RESET)"

# Default target
.DEFAULT_GOAL := help
