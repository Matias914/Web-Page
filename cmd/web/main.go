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
	// Se carga el archivo '.env' para trabajar en local sin contenedores
	_ = godotenv.Load()

	// Se cargan los datos de entorno en el struct de Configuracion
	cfg := config.LoadConfig()

	// Se crea un pool de conexiones a la base
	db, err := postgres.NewDB(cfg.DSN)
	if err != nil {
		log.Fatalf("error: no se pudo conectar a la base de datos: %v", err)
	}
	// Nos aseguramos de que el pool de conexiones se cierre de forma limpia
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("error al cerrar la conexión con la base de datos: %v", err)
		}
	}(db)

	log.Println("conexión a la base de datos establecida exitosamente")

	// Se crean las dependencias de los handlers
	renderer, err := handler.NewRenderer("./web/templates/")
	if err != nil {
		log.Fatalf("no pudo crearse una instancia de Renderer: %v", err)
	}
	app := &handler.Application{
		Queries:  sqlc.New(db),
		Renderer: renderer,
	}

	// Se crea un servidor con timeouts
	server := &http.Server{
		// Host Implicito, por defecto, 0.0.0.0
		Addr: ":" + cfg.AppPort,
		// GetRouter devuelve el Handler Mux
		// WithLogging devuelve el handler que invoca el Handler Mux
		// WithRecovery devuelve el Handler que invoca el Handler WithLogging
		// El Servidor deriva sus consultas al único Handler WithRecovery
		Handler:      middleware.WithRecovery(middleware.WithLogging(app.GetRouter())),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	// Logs de inicio y del resultado de ListenAndServe
	// Mostramos el DSN para verificar que se cargó bien la info de la DB
	log.Printf("servidor ejecutándose en http://localhost:%s", cfg.AppPort)
	log.Printf("conectando a la base de datos con DSN: '%s'", cfg.DSN)

	log.Fatal(server.ListenAndServe())
}
