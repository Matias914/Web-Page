# ---------------------------
# Etapa de construcción
# ---------------------------
FROM golang:1.24.6-alpine AS builder
LABEL authors="matiasortiz"

# Establecer variables para un binario estático
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /app

# Copiar dependencias primero para caching (descomentá si tenés go.mod y go.sum)
# COPY go.mod go.sum ./
# RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación de forma limpia
RUN go build -ldflags="-s -w" -o /app/main .

# ---------------------------
# Etapa final minimalista
# ---------------------------
FROM alpine:3.18

# Instalar certificados y crear usuario no root
RUN apk add --no-cache ca-certificates \
    && addgroup -S app \
    && adduser -S -G app app

WORKDIR /app

# Copiar binario y assets desde builder, asignando propietario
COPY --from=builder --chown=app:app /app/main /app/main
COPY --from=builder --chown=app:app /app/templates /app/templates
COPY --from=builder --chown=app:app /app/errors /app/errors
COPY --from=builder --chown=app:app /app/static /app/static

# Asegurar permisos de ejecución
RUN chmod +x /app/main

# Exponer puerto y definir usuario no root
EXPOSE 8080
USER app

# Comando de ejecución
CMD ["/app/main"]
