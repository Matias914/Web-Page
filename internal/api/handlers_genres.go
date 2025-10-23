package api

import (
	"net/http"
)

// @Summary    Lista todos los géneros de una película
// @Description Obtiene la lista de géneros asociados a una película específica.
// @Tags       Movies, Genres
// @Accept     json
// @Produce    json
// @Param      id1    path      int       true   "ID de la película"
// @Success    200    {array}   service.Genre "Lista de géneros exitosa"
// @Failure    400    {object}  api.JsonError  "Error de validación en el ID de la película"
// @Failure    500    {object}  api.JsonError  "Error interno del servidor"
// @Router     /movies/{id1}/genres [get]
func (api *API) handleGetMovieGenres(w http.ResponseWriter, r *http.Request) {
	movieID, err := api.handleSingleIdentifierParsing(r)
	page, rows, err := api.handlePageAndRowsParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	genres, err := api.GenreService.GetMovieGenresList(r.Context(), movieID, page, rows)
	if err != nil {
		api.handleErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	api.handleJsonResponse(w, http.StatusOK, genres)
}

// @Summary    Obtiene un género asociado a una película (Por su ID de género)
// @Description Obtiene la información de un género específico asociado a una película, dado su ID de género.
// @Tags       Movies, Genres
// @Accept     json
// @Produce    json
// @Param      id1    path      int       true   "ID de la película"
// @Param      id2    path      int       true   "ID del género"
// @Success    200    {object}  service.Genre  "Género encontrado exitosamente"
// @Failure    400    {object}  api.JsonError  "Error de validación en el ID"
// @Failure    404    {object}  api.JsonError  "Género no encontrado para esa película"
// @Router     /movies/{id1}/genres/{id2} [get]
func (api *API) handleGetMovieGenre(w http.ResponseWriter, r *http.Request) {
	//movieID, genreID, err := api.handleDoubleIdentifierParsing(r)
	//page, rows, err := api.handlePageAndRowsParsing(r)
	//if err != nil {
	//	api.handleErrorResponse(w, http.StatusBadRequest, err)
	//	return
	//}
	//
	//// Lógica para obtener el género por ID y verificar que esté asociado a la película
	//genre, err := api.GenreService.GetMovieGenresList(r.Context(), movieID, genreID, page, rows)
	//if err != nil {
	//	// Asumimos que el Service devolvería un error que se mapea a 404 si no se encuentra
	//	api.handleErrorResponse(w, http.StatusNotFound, err)
	//	return
	//}
	//api.handleJsonResponse(w, http.StatusOK, genre)
}

func (api *API) handlePostMovieGenre(w http.ResponseWriter, r *http.Request) {
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

func (api *API) handleDeleteMovieGenre(w http.ResponseWriter, r *http.Request) {
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
