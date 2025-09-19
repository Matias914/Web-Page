# ---------------------------
# Etapa de construcción (Builder)
# ---------------------------
FROM golang:1.24.6-alpine AS builder
LABEL authors="matiasortiz"

# Establecer variables para un binario estático
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

# 1. Copiamos solo los archivos de dependencias para aprovechar la caché de Docker.
#    Esta capa solo se reconstruye si go.mod o go.sum cambian.
COPY go.mod go.sum ./

# 2. Descargamos las dependencias.
RUN go mod download
RUN go mod verify

# 3. Copiamos el resto del código fuente.
#    Esta capa se reconstruye si cualquier archivo .go cambia.
COPY . .

# 4. Construimos la aplicación apuntando al main correcto dentro de cmd/web.
RUN go build -ldflags="-s -w" -o /app/main ./cmd/web

# ---------------------------
# Etapa final minimalista
# ---------------------------
FROM alpine:3.18

# Instalar certificados y crear usuario no root
RUN apk add --no-cache ca-certificates \
    && addgroup -S app \
    && adduser -S -G app app

WORKDIR /app

# Copiamos el binario compilado desde la etapa de construcción
COPY --from=builder --chown=app:app /app/main /app/main

# Copiamos la carpeta 'web' completa que contiene templates y static.
# Esto simplifica las múltiples líneas de COPY que tenías antes.
COPY --from=builder --chown=app:app /app/web ./web

# Exponer puerto y definir usuario no root
EXPOSE 8080
USER app

# Comando de ejecución
CMD ["/app/main"]