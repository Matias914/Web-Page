package transport

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Application) GetRouter() *chi.Mux {
	mux := chi.NewRouter()
	fsr := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fsr))

	mux.NotFound(app.handleNotFound)
	mux.MethodNotAllowed(app.handleNotFound)

	mux.Get("/", app.handleIndexPage)
	mux.Get("/catalog", app.handleCatalogPage)
	mux.Get("/management", app.handleControlPage)

	// Movies
	mux.Post("/movies", app.handleMoviesCreate)

	// Genres
	mux.Post("/genres", app.handleGenresCreate)

	// Celebrities
	mux.Post("/celebrities", app.handleCelebritiesCreate)

	return mux
}
