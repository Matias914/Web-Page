# ---------------------
# Etapa de construcción
# ---------------------

FROM golang:1.24.6-alpine AS builder
LABEL authors="matiasortiz"

# Establecer variables para un binario estático
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go mod verify

COPY . .

RUN go build -ldflags="-s -w" -o /app/main ./cmd/web

# -----------------------
# Etapa final minimalista
# -----------------------

FROM alpine:3.18

RUN apk add --no-cache ca-certificates \
    && addgroup -S app \
    && adduser -S -G app app

WORKDIR /app

COPY --from=builder --chown=app:app /app/main /app/main
COPY --from=builder --chown=app:app /app/web ./web

EXPOSE 8080

USER app

CMD ["/app/main"]