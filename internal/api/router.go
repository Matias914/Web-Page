package api

import (
	"github.com/go-chi/chi/v5"
)

func (api *API) GetRouter() *chi.Mux {
	mux := chi.NewRouter()

	mux.Route("/movies", func(sub chi.Router) {
		sub.Get("/", api.handleGetMovies)
		sub.Post("/", api.handlePostMovie)

		sub.Route("/{id1}", func(sub chi.Router) {
			sub.Get("/", api.handleGetMovie)
			sub.Put("/", api.handlePutMovie)
			sub.Delete("/", api.handleDeleteMovie)

			sub.Route("/genres", func(sub chi.Router) {
				sub.Get("/", api.handleGetMovieGenres)
				sub.Post("/", api.handlePostMovieGenre)
				sub.Get("/{id2}", api.handleGetMovieGenre)
				sub.Delete("/{id2}", api.handleDeleteMovieGenre)
			})

			sub.Route("/celebrities", func(sub chi.Router) {
				sub.Get("/", api.handleGetMovieCelebrities)
				sub.Post("/", api.handlePostMovieCelebrityRole)
				sub.Get("/{id2}", api.handleGetMovieCelebrityRoles)
				sub.Put("/{id2}", api.handlePutMovieCelebrityRoles)
				sub.Delete("/{id2}", api.handleDeleteMovieCelebrityRole)
			})
		})
	})

	return mux
}
