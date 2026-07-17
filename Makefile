# ==========================================================
# Velora Makefile
# ==========================================================

.DEFAULT_GOAL := help

# ==========================================================
# Variables
# ==========================================================

BACKEND_DIR := backend
APP_BINARY := $(BACKEND_DIR)/bin/api

# ==========================================================
# Help
# ==========================================================

.PHONY: help

help:
	@echo ""
	@echo "==============================================="
	@echo " Velora Development Commands"
	@echo "==============================================="
	@echo ""
	@echo "Docker"
	@echo "  make up           Start infrastructure"
	@echo "  make down         Stop infrastructure"
	@echo "  make restart      Restart infrastructure"
	@echo "  make logs         Show docker logs"
	@echo "  make ps           Show containers"
	@echo "  make clean        Remove containers & volumes"
	@echo ""
	@echo "Backend"
	@echo "  make run          Run backend"
	@echo "  make build        Build backend"
	@echo "  make start        Run compiled binary"
	@echo "  make fmt          Format Go code"
	@echo "  make vet          Run go vet"
	@echo "  make tidy         Run go mod tidy"
	@echo "  make test         Run tests"
	@echo ""
	@echo "Database"
	@echo "  make db           Open PostgreSQL"
	@echo "  make redis        Open Redis CLI"
	@echo "  make migrate-up   Run migrations"
	@echo "  make migrate-down Rollback migrations"
	@echo "  make migrate-version  Show current migration version"
	@echo "  make migrate-force    Force migration to specific version"
	@echo ""
	@echo "AI"
	@echo "  make ollama       Open Ollama container"
	@echo ""

# ==========================================================
# Docker
# ==========================================================

.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down

.PHONY: restart
restart:
	docker compose down
	docker compose up -d

.PHONY: logs
logs:
	docker compose logs -f

.PHONY: ps
ps:
	docker compose ps

.PHONY: clean
clean:
	docker compose down -v --remove-orphans

# ==========================================================
# Backend
# ==========================================================

.PHONY: run
run:
	cd $(BACKEND_DIR) && go run ./cmd/api

.PHONY: build
build:
	cd $(BACKEND_DIR) && go build -o bin/api ./cmd/api

.PHONY: start
start:
	$(APP_BINARY)

.PHONY: fmt
fmt:
	cd $(BACKEND_DIR) && go fmt ./...

.PHONY: vet
vet:
	cd $(BACKEND_DIR) && go vet ./...

.PHONY: tidy
tidy:
	cd $(BACKEND_DIR) && go mod tidy

.PHONY: test
test:
	cd $(BACKEND_DIR) && go test ./...

# ==========================================================
# Database
# ==========================================================

.PHONY: db
db:
	docker exec -it velora-postgres psql -U postgres -d velora

.PHONY: redis
redis:
	docker exec -it velora-redis redis-cli

# ==========================================================
# Migration
# ==========================================================

MIGRATE_DB_URL := postgres://postgres:postgres@localhost:15432/velora?sslmode=disable

.PHONY: migrate-up
migrate-up:
	migrate -path $(BACKEND_DIR)/migrations -database "$(MIGRATE_DB_URL)" up

.PHONY: migrate-down
migrate-down:
	migrate -path $(BACKEND_DIR)/migrations -database "$(MIGRATE_DB_URL)" down 1

.PHONY: migrate-version
migrate-version:
	migrate -path $(BACKEND_DIR)/migrations -database "$(MIGRATE_DB_URL)" version

.PHONY: migrate-force
migrate-force:
	migrate -path $(BACKEND_DIR)/migrations -database "$(MIGRATE_DB_URL)" force $(VERSION)

# ==========================================================
# Ollama
# ==========================================================

.PHONY: ollama
ollama:
	docker exec -it velora-ollama bash
