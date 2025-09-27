package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Matias914/Web-Page/internal/storage/postgres/sqlc"
)

// Application contiene las dependencias de la aplicación, como el acceso a
// la base de datos. Esto permite la inyección de dependencias en los handlers.
type Application struct {
	Queries  *sqlc.Queries
	Renderer *Renderer
}

// createResponse crea la respuesta HTTP usando el renderer para generar el contenido '.html' en un buffer.
func (app *Application) handleResponse(w http.ResponseWriter, status int, name string, data interface{}) error {
	buf, err := app.Renderer.renderToBuffer(name, data)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Printf("error al escribir el buffer en la respuesta: %v", err)
	}
	return nil
}

// handleServerError utiliza el createResponse para generar una respuesta 500.
// Caso contrario, maneja el error.
func (app *Application) handleServerError(w http.ResponseWriter, err error) {
	log.Printf("error interno del servidor: %v", err)
	if err := app.handleResponse(w, http.StatusInternalServerError, "500", nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// handleServerError utiliza el createResponse para generar una respuesta 4xx.
// Caso contrario, delega el error.
func (app *Application) handleClientError(w http.ResponseWriter, status int) {
	file := fmt.Sprintf("%d", status)
	if err := app.handleResponse(w, status, file, nil); err != nil {
		app.handleServerError(w, err)
	}
}
