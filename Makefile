# Variables
AIR ?= air
GO ?= go

# Archivos y directorios
BIN_DIR := tmp
BIN := $(BIN_DIR)/main

.PHONY: help server app build run test tidy clean

help:
	@echo "Comandos disponibles:"
	@echo "  server     Inicia el servidor con Air (hot-reload)"
	@echo "  app        Alias de 'server'"
	@echo "  build      Compila el binario en $(BIN)"
	@echo "  run        Ejecuta el binario compilado"
	@echo "  test       Ejecuta las pruebas"
	@echo "  tidy       Ejecuta 'go mod tidy'"
	@echo "  clean      Elimina artefactos de build"
	@echo "  db-up      Levanta el servicio de base de datos (docker-compose)"
	@echo "  db-down    Detiene y elimina los contenedores de docker-compose"
	@echo "  dev        Levanta la base de datos y el servidor para desarrollo"


server:
	$(AIR)

# Alias para iniciar la app (si la app es el servidor con Air)
app: server

build:
	$(GO) build -o $(BIN) ./cmd/web

run: build
	./$(BIN)

test:
	$(GO) test -v ./...

tidy:
	$(GO) mod tidy

clean:
	rm -rf $(BIN_DIR)

db-up:
	docker-compose up -d db

dev: db-up
	$(MAKE) server

db-down:
	docker-compose down