package transport

import (
	"net/http"
)

func (app *Application) handleHome(w http.ResponseWriter, r *http.Request) {
	if err := app.handleResponse(w, http.StatusOK, "index", nil); err != nil {
		app.handleServerError(w, err)
	}
}

func (app *Application) handleNotFound(w http.ResponseWriter, r *http.Request) {
	if err := app.handleResponse(w, http.StatusNotFound, "404", nil); err != nil {
		app.handleServerError(w, err)
	}
}

func (app *Application) handleCatalog(w http.ResponseWriter, r *http.Request) {
	if err := app.handleResponse(w, http.StatusOK, "catalog", nil); err != nil {
		app.handleServerError(w, err)
	}
}

func (app *Application) handleManagement(w http.ResponseWriter, r *http.Request) {
	if err := app.handleResponse(w, http.StatusOK, "management", nil); err != nil {
		app.handleServerError(w, err)
	}
}
