package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Matias914/Web-Page/config"
	"github.com/Matias914/Web-Page/handlers"
	"github.com/Matias914/Web-Page/middleware"
)

func main() {
	// Se establece un puerto y host
	c := config.LoadConfig()
	// Se crea un servidor con timeouts
	server := &http.Server{
		Addr: c.Host + ":" + c.Port,
		// Routes devuelve el Handler Mux
		// WithRecovery devuelve el Handler que invoca el Handler Mux
		// WithLogging devuelve el handler que invoca el Handler WithRecovery
		// El Servidor deriva sus consultas al único Handler WithLogging
		Handler:      middleware.WithLogging(middleware.WithRecovery(handlers.GetRouter())),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	// TODO ver qué hacer con ese http que me molesta
	// Logs de inicio y del resultado de ListenAndServe
	log.Printf("Servidor ejecutándose en http://%s:%s", c.Host, c.Port)
	log.Fatal(server.ListenAndServe())
}
