package transport

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Application) GetRouter() *chi.Mux {
	mux := chi.NewRouter()
	fsr := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fsr))

	mux.HandleFunc("/", app.handleHome)
	mux.HandleFunc("/catalog", app.handleCatalog)
	mux.HandleFunc("/management", app.handleManagement)
	mux.NotFound(app.handleNotFound)
	return mux
}
