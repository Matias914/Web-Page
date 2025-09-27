locals {
    auth = "${getenv("POSTGRES_USER")}:${getenv("POSTGRES_PASSWORD")}"
    host = "localhost:${getenv("POSTGRES_PORT")}"
    dsn  = "postgres://${local.auth}@${local.host}/${getenv("POSTGRES_DB")}?sslmode=disable"
}

env "local" {
    url = local.dsn
    dev = local.dsn

    migration {
      dir = "file://internal/storage/postgres/migrations"
    }

    schema {
      src = "file://internal/storage/postgres/schema/schema.sql"
    }
}