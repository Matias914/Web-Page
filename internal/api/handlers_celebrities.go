package api

import "net/http"

func (api *API) handleGetMovieCelebrities(w http.ResponseWriter, r *http.Request) {
	page, rows, err := api.handlePageAndRowsParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	movieId, err := api.handleSingleIdentifierParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	genres, err := api.GenreService.GetMovieGenresList(r.Context(), movieId, page, rows)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	api.handleJsonResponse(w, http.StatusOK, genres)
}

func (api *API) handleGetMovieCelebrityRoles(w http.ResponseWriter, r *http.Request) {
	page, rows, err := api.handlePageAndRowsParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	movieId, err := api.handleSingleIdentifierParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	genres, err := api.GenreService.GetMovieGenresList(r.Context(), movieId, page, rows)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	api.handleJsonResponse(w, http.StatusOK, genres)
}

func (api *API) handlePutMovieCelebrityRoles(w http.ResponseWriter, r *http.Request) {
	page, rows, err := api.handlePageAndRowsParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	movieId, err := api.handleSingleIdentifierParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	genres, err := api.GenreService.GetMovieGenresList(r.Context(), movieId, page, rows)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	api.handleJsonResponse(w, http.StatusOK, genres)
}

func (api *API) handlePostMovieCelebrityRole(w http.ResponseWriter, r *http.Request) {
	page, rows, err := api.handlePageAndRowsParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	movieId, err := api.handleSingleIdentifierParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	genres, err := api.GenreService.GetMovieGenresList(r.Context(), movieId, page, rows)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	api.handleJsonResponse(w, http.StatusOK, genres)
}

func (api *API) handleDeleteMovieCelebrityRole(w http.ResponseWriter, r *http.Request) {
	page, rows, err := api.handlePageAndRowsParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	movieId, err := api.handleSingleIdentifierParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	genres, err := api.GenreService.GetMovieGenresList(r.Context(), movieId, page, rows)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	api.handleJsonResponse(w, http.StatusOK, genres)
}
