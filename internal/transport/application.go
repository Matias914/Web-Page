package transport

import (
	"log"
	"net/http"
)

type Application struct {
	Renderer *Renderer
}

func (app *Application) handleResponse(w http.ResponseWriter, status int, name string, data interface{}) error {
	buf, err := app.Renderer.renderToBuffer(name, data)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	// Error Irreversible porque solo falla si hubo una desconexion
	// Además, protocolo HTTP prohibe escribir el contenido antes de la cabecera
	if _, err := buf.WriteTo(w); err != nil {
		log.Printf("(handleResponse) no se pudo escribir el contenido del buffer en la respuesta: %v", err)
	}
	return nil
}

func (app *Application) handleServerError(w http.ResponseWriter, err error) {
	log.Printf("(handleServerError) error interno del servidor: %v", err)
	if err := app.handleResponse(w, http.StatusInternalServerError, "500", nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
