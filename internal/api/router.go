package api

import (
	"github.com/go-chi/chi/v5"
)

func (api *API) GetRouter() *chi.Mux {
	mux := chi.NewRouter()

	// Celebrities
	mux.Route("/celebrities", func(sub chi.Router) {
		sub.Get("/", api.handleGetCelebrities)
		sub.Post("/", api.handlePostCelebrity)
		sub.Route("/{id1}", func(sub chi.Router) {
			sub.Get("/", api.handleGetCeleberity)
			sub.Put("/", api.handlePutCelebrity)
			sub.Delete("/", api.handleDeleteCelebrity)
			// Celebrities - Role
			sub.Route("/roles", func(sub chi.Router) {
				sub.Get("/", api.handleGetCelebrityRoles)
			})
		})
	})

	// Genres
	mux.Route("/genres", func(sub chi.Router) {
		sub.Get("/", api.handleGetGenres)
		sub.Post("/", api.handlePostGenre)
		sub.Route("/{id1}", func(sub chi.Router) {
			sub.Get("/", api.handleGetGenre)
			sub.Put("/", api.handlePutGenre)
			sub.Delete("/", api.handleDeleteGenre)
			// Genres - Categories
			sub.Route("/categories", func(sub chi.Router) {
				sub.Get("/", api.handleGetGenreMovies)
			})
		})
	})

	// Movies
	mux.Route("/movies", func(sub chi.Router) {
		sub.Get("/", api.handleGetMovies)
		sub.Post("/", api.handlePostMovie)
		sub.Route("/{id1}", func(sub chi.Router) {
			sub.Get("/", api.handleGetMovie)
			sub.Put("/", api.handlePutMovie)
			sub.Delete("/", api.handleDeleteMovie)
			// Movies - Categories
			sub.Route("/categories", func(sub chi.Router) {
				sub.Get("/", api.handleGetMovieGenres)
				sub.Post("/", api.handlePostCategory)
				sub.Get("/{id2}", api.handleGetCategory)
				sub.Delete("/{id2}", api.handleDeleteCategory)
			})
			// Movies - Ratings
			sub.Route("/ratings", func(sub chi.Router) {
				sub.Get("/", api.handleGetMovieRatings)
			})
			// Movies - Reviews
			sub.Route("/reviews", func(sub chi.Router) {
				sub.Get("/", api.handleGetMovieReviews)
			})
			// Movies - Roles
			sub.Route("/roles", func(sub chi.Router) {
				sub.Get("/", api.handleGetMovieRoles)
				sub.Post("/", api.handlePostRole)
				sub.Get("/{id2}", api.handleGetRole)
				sub.Put("/{id2}", api.handlePutRole)
				sub.Delete("/{id2}", api.handleDeleteRole)
			})
		})
	})

	// Users
	mux.Route("/users", func(sub chi.Router) {
		sub.Get("/", api.handleGetUsers)
		sub.Post("/", api.handlePostUser)
		sub.Route("/{id1}", func(sub chi.Router) {
			sub.Get("/", api.handleGetUser)
			sub.Put("/", api.handlePutUser)
			sub.Delete("/", api.handleDeleteUser)
			// Users - Ratings
			sub.Route("/ratings", func(sub chi.Router) {
				sub.Get("/", api.handleGetUserRatings)
				sub.Post("/", api.handlePostRating)
				sub.Get("/{id2}", api.handleGetRating)
				sub.Put("/{id2}", api.handlePutRating)
				sub.Delete("/{id2}", api.handleDeleteRating)
			})
			// Users - Reviews
			sub.Route("/reviews", func(sub chi.Router) {
				sub.Get("/", api.handleGetUserReviews)
				sub.Post("/", api.handlePostReview)
				sub.Get("/{id2}", api.handleGetReview)
				sub.Put("/{id2}", api.handlePutReview)
				sub.Delete("/{id2}", api.handleDeleteReview)
			})
		})
	})

	return mux
}
