package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/Matias914/Web-Page/docs"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/Matias914/Web-Page/internal/config"
	"github.com/Matias914/Web-Page/internal/middleware"
	"github.com/Matias914/Web-Page/internal/service"
	"github.com/Matias914/Web-Page/internal/storage/postgres"
	"github.com/Matias914/Web-Page/internal/storage/postgres/sqlc"
	"github.com/Matias914/Web-Page/internal/transport"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

// @title			API de Películas UNICEN
// @version			1.0
// @description		Esta es la API REST para la aplicación de gestión de películas.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

func main() {
	_ = godotenv.Load()
	cfg := config.LoadConfig()

	/* ----------------------------------------------------------------------------- */
	/*						    	DATABASE CONNECTION	    				         */
	/* ----------------------------------------------------------------------------- */

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

	/* ----------------------------------------------------------------------------- */
	/*						      	    SERVICES    	      				         */
	/* ----------------------------------------------------------------------------- */

	validationService, err := service.NewValidationService()
	if err != nil {
		log.Fatalf("error: no se pudo crear el servicio de validacion: %v", err)
	}

	movieService := service.MovieService{Queries: sqlc.New(db)}
	genreService := service.GenreService{Queries: sqlc.New(db)}
	celebrityService := service.CelebrityService{Queries: sqlc.New(db)}

	/* ----------------------------------------------------------------------------- */
	/*						      	   APPLICATION  	      				         */
	/* ----------------------------------------------------------------------------- */

	app := &transport.Application{
		MovieService:     &movieService,
		GenreService:     &genreService,
		CelebrityService: &celebrityService,

		ValidationService: &validationService,
	}

	/* ----------------------------------------------------------------------------- */
	/*						       	      SERVER  	        				         */
	/* ----------------------------------------------------------------------------- */

	mainRouter := chi.NewRouter()
	mainRouter.Mount("/", app.GetRouter())
	mainRouter.Get("/swagger/*", httpSwagger.WrapHandler)

	server := &http.Server{
		Addr:         ":" + cfg.AppPort,
		Handler:      middleware.WithRecovery(middleware.WithLogging(mainRouter)),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("servidor ejecutándose en http://localhost:%s", cfg.AppPort)
	log.Printf("conectando a la base de datos con DSN: '%s'", cfg.DSN)
	log.Fatal(server.ListenAndServe())
}
