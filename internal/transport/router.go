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
	mux.Get("/clear-notification", app.handleClearNotification)

	// CRUD Movies
	mux.Post("/movies", app.handleMoviesCreate)
	mux.Put("/movies/{id}", app.handleMoviesUpdate)
	mux.Delete("/movies/{id}", app.handleMovieDelete)
	mux.Get("/movies/{id}/edit", app.handleMovieEditForm)

	// CRUD Genres
	mux.Post("/genres", app.handleGenresCreate)
	mux.Put("/genres/{id}", app.handleGenresUpdate)
	mux.Delete("/genres/{id}", app.handleGenreDelete)
	mux.Get("/genres/{id}/edit", app.handleGenreEditForm)

	// CRUD Celebrities
	mux.Post("/celebrities", app.handleCelebritiesCreate)
	mux.Put("/celebrities/{id}", app.handleCelebritiesUpdate)
	mux.Delete("/celebrities/{id}", app.handleCelebrityDelete)
	mux.Get("/celebrities/{id}/edit", app.handleCelebrityEditForm)

	return mux
}
