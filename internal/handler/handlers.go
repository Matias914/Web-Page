package handler

import (
	"net/http"
)

// Home es una funcion que maneja las solicitudes que se dirigen al root o a URL desconocidas.
func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.handleClientError(w, http.StatusNotFound)
		return
	}
	if err := app.handleResponse(w, http.StatusOK, "index", nil); err != nil {
		app.handleServerError(w, err)
	}
}
