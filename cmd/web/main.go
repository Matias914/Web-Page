package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/Matias914/Web-Page/internal/config"
	"github.com/Matias914/Web-Page/internal/handler"
	"github.com/Matias914/Web-Page/internal/middleware"
	"github.com/Matias914/Web-Page/internal/storage/postgres"
	"github.com/Matias914/Web-Page/internal/storage/postgres/sqlc"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := config.LoadConfig()

	db, err := postgres.NewDB(cfg.DSN)
	if err != nil {
		log.Fatalf("error: no se pudo conectar a la base de datos: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("error al cerrar la conexión con la base de datos: %v", err)
		}
	}(db)
	log.Println("conexión a la base de datos establecida exitosamente")

	renderer, err := handler.NewRenderer("./web/templates/")
	if err != nil {
		log.Fatalf("no pudo crearse una instancia de Renderer: %v", err)
	}
	app := &handler.Application{
		Queries:  sqlc.New(db),
		Renderer: renderer,
	}

	server := &http.Server{
		Addr:         ":" + cfg.AppPort,
		Handler:      middleware.WithRecovery(middleware.WithLogging(app.GetRouter())),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	log.Printf("servidor ejecutándose en http://localhost:%s", cfg.AppPort)
	log.Printf("conectando a la base de datos con DSN: '%s'", cfg.DSN)
	log.Fatal(server.ListenAndServe())
}
