include .env
export

# ==============================================================================
# CONFIGURACIÓN
# ==============================================================================

GO     ?= go
AIR    ?= air
SQLC   ?= sqlc
DOCKER ?= docker-compose
ATLAS  ?= atlas

# Variables del Proyecto
APP_NAME := App Web en Go
BIN_DIR  := tmp
BIN      := $(BIN_DIR)/main

# Variables de Base de Datos
SCHEMA_FILE     := internal/storage/postgres/schema/schema.sql
MIGRATIONS_PATH := internal/storage/postgres/migrations

.DEFAULT_GOAL := help

# ==============================================================================
# AYUDA
# ==============================================================================

.PHONY: help
help:
	@echo "Gestión del Proyecto ($(APP_NAME))"
	@echo "--------------------------------------------------"
	@echo "Uso: make [comando]"
	@echo ""
	@echo "Comandos Principales:"
	@echo "  dev           - Inicia DB, aplica migraciones y corre el servidor en modo desarrollo."
	@echo "  prod  		   - Construye y levanta toda la aplicación (app y db) en Docker."
	@echo "  prod-down     - Detiene los contenedores de producción y remueve contenedores huérfanos."
	@echo "  server        - Corre el servidor con hot-reload (Air)."
	@echo ""
	@echo "Comandos de Base de Datos (Docker):"
	@echo "  db-up         - Inicia el contenedor de la base de datos."
	@echo "  db-down       - Detiene el contenedor de la base de datos."
	@echo "  db-nuke       - Detiene y elimina los volúmenes de la base de datos."
	@echo ""
	@echo "Comandos de Migraciones (Atlas):"
	@echo "  migrate-diff  - Crea un nuevo archivo de migración (requiere NAME)."
	@echo "  migrate-up    - Aplica todas las migraciones pendientes."
	@echo "  migrate-set   - Revierte a una migración anterior (requiere VERSION)."
	@echo ""
	@echo "Comandos de Desarrollo:"
	@echo "  sqlc-gen      - Genera código Go desde las queries SQL."
	@echo "  build         - Compila el binario de la aplicación."
	@echo "  run           - Compila y ejecuta el binario."
	@echo "  test          - Ejecuta todas las pruebas."
	@echo "  tidy          - Ordena y verifica las dependencias de Go."
	@echo "  clean         - Elimina el directorio de binarios '$(BIN_DIR)'."
	@echo "  docker-clean  - Limpieza completa del proyecto actual en Docker."
	@echo "  docker-nuke   - Elimina contenedores y volúmenes de Docker no utilizados."


# ==============================================================================
# CICLO DE VIDA DE DESARROLLO
# ==============================================================================

.PHONY: dev prod prod-down server

dev: db-up sqlc-gen migrate-up server

prod:
	@echo "🐋  Construyendo y levantando la aplicación en modo producción..."
	@$(DOCKER) up -d --build

prod-down:
	@echo "🐳  Deteniendo el entorno de producción (Docker)..."
	@$(DOCKER) down --remove-orphans

server:
	@echo "🚀  Iniciando servidor con hot-reload..."
	@$(AIR)

# ==============================================================================
# GESTIÓN DE LA BASE DE DATOS
# ==============================================================================

.PHONY: db-up wait-db db-down db-nuke docker-clean docker-nuke
db-up:
	@echo "🐘  Levantando la base de datos con Docker..."
	@$(DOCKER) up -d db
	@$(MAKE) wait-db

wait-db:
	@echo "⏳  Esperando a que la base de datos esté lista para aceptar conexiones..."
	@until $(DOCKER) exec db pg_isready -U "$(POSTGRES_USER)" -d "$(POSTGRES_DB)" -q; do \
		sleep 1; \
	done
	@echo "✅  ¡Base de datos lista!"

db-down:
	@echo "✋  Deteniendo la base de datos..."
	@$(DOCKER) down

db-nuke:
	@echo "🔥  Eliminando la base de datos y sus volúmenes..."
	@$(DOCKER) down -v

docker-clean:
	@echo "💣  Limpiando todo lo relacionado con este proyecto en Docker..."
	@$(DOCKER) down -v --rmi 'local' --remove-orphans

docker-nuke:
	@echo "🔥  Destruyendo sistema Docker..."
	@$(DOCKER) down --remove-orphans
	@docker system prune -af
	@docker volume prune -f

# ==============================================================================
# MIGRACIONES (ATLAS & SQLC)
# ==============================================================================

.PHONY: sqlc-gen migrate-diff migrate-up migrate-down
sqlc-gen:
	@echo "🧬  Generando código Go con sqlc..."
	@$(SQLC) generate

migrate-diff:
	@if [ -z "$(NAME)" ]; then echo "Error: La variable NAME es requerida. Ej: make migrate-diff NAME=create_users_table"; exit 1; fi
	@echo "🔍  Creando nuevo archivo de migración llamado '$(NAME)'..."
	@$(ATLAS) migrate diff $(NAME) --to "file://$(SCHEMA_FILE)" --dev-url "docker://postgres/15/dev?search_path=public" --env local

migrate-up:
	@echo "📈  Aplicando migraciones pendientes..."
	@$(ATLAS) migrate apply --env local

migrate-set:
	@if [ -z "$(VERSION)" ]; then \
		echo "Error: VERSION es requerido. Ej: make migrate-set VERSION=1"; \
		exit 1; \
	fi
	@echo "📉  Revertiendo la última migración..."
	@$(ATLAS) migrate set $(VERSION) --env local


# ==============================================================================
# COMPILACIÓN Y PRUEBAS
# ==============================================================================

.PHONY: build run test tidy clean
build:
	@echo "🛠️  Compilando el binario en $(BIN)..."
	@$(GO) build -o $(BIN) ./cmd/web

run: build
	@echo "▶️  Ejecutando el binario..."
	@./$(BIN)

test:
	@echo "🧪  Ejecutando pruebas..."
	@$(GO) test -v ./...

tidy:
	@echo "📦  Ordenando dependencias de Go..."
	@$(GO) mod tidy
	@$(GO) mod verify

clean:
	@echo "🗑️  Limpiando el directorio $(BIN_DIR)..."
	@rm -rf $(BIN_DIR)