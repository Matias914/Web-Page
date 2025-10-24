package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Matias914/Web-Page/internal/service"
)

// @Summary    	Agrega un género a una película
// @Description Agrega un género dado por ID a una película también dada por ID.
// @Tags       	Categories
// @Accept     	json
// @Produce    	json
// @Param      	movieID path      int   					true	"ID de la película"
// @Param      	genreID	body      service.AddCategoryInput 	true   	"ID del género"
// @Success    	204    												"Género agregado exitosamente"
// @Failure    	400    	{object}  JsonError  						"Error de validación de los parámetros"
// @Failure    	404    	{object}  JsonError  						"Error de película o género no encontrado"
// @Failure    	409    	{object}  JsonError  						"Error de duplicación de recursos
// @Failure    	500    	{object}  JsonError  						"Error interno del servidor"
// @Router     	/movies/{movieID}/genres	 				[post]
func (api *API) handlePostMovieGenre(w http.ResponseWriter, r *http.Request) {
	movieID, err := api.handleSingleIdentifierParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	var genre service.AddCategoryInput
	if err := json.NewDecoder(r.Body).Decode(&genre); err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	if err := api.ValidationService.Validate(genre); err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	err = api.CategoryService.AddMovieGenre(r.Context(), movieID, genre.GenreID)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrMovieNotFound):

			api.handleErrorResponse(w, http.StatusNotFound, err)
		case errors.Is(err, service.ErrGenreNotFound):
			api.handleErrorResponse(w, http.StatusNotFound, err)
		case errors.Is(err, service.ErrCategoryDuplicated):
			api.handleErrorResponse(w, http.StatusConflict, err)
		default:
			api.handleErrorResponse(w, http.StatusInternalServerError, err)
		}
		return
	}
	api.handleJsonResponse(w, http.StatusNoContent, nil)
}

// @Summary    	Verifica que la película tiene asociado un género
// @Description Verifica si una película dada por ID tiene un determinado género también dado por ID.
// @Tags       	Categories
// @Accept     	json
// @Produce    	json
// @Param      	movieID    	path     	int      	true   	"ID de la película"
// @Param      	genreID    	path     	int     	true   	"ID del género"
// @Success    	204    							 			"Existe la relación entre la película y el género"
// @Success    	400    		{object} 	JsonError  			"Error de validación de los parámetros"
// @Failure    	404    		{object}  	JsonError		  	"Error de relación inexistente entre la película y el género"
// @Failure    	500    		{object}  	JsonError  			"Error interno del servidor"
// @Router     	/movies/{movieID}/genres/{genreID} 	[get]
func (api *API) handleGetMovieGenre(w http.ResponseWriter, r *http.Request) {
	movieID, genreID, err := api.handleDoubleIdentifierParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	err = api.CategoryService.HasMovieWithGenre(r.Context(), movieID, genreID)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrCategoryNotFound):
			api.handleErrorResponse(w, http.StatusNotFound, err)
		default:
			api.handleErrorResponse(w, http.StatusInternalServerError, err)
		}
		return
	}
	api.handleJsonResponse(w, http.StatusNoContent, nil)
}

// @Summary    	Borra un género de una película
// @Description Borra un género dado por ID asociado a una película dada también por ID permanentemente.
// @Tags       	Categories
// @Accept     	json
// @Produce    	json
// @Param      	movieID    	path      int  			true   	"ID de la película"
// @Param      	genreID    	path      int       	true   	"ID del género"
// @Success    	204    							 			"Género borrado exitosamente"
// @Success    	400    		{object}  JsonError  			"Error de validación de los parámetros"
// @Failure    	404    		{object}  JsonError		  		"Error de relación inexistente entre la película y el género"
// @Failure    	500    		{object}  JsonError  			"Error interno del servidor"
// @Router     	/movies/{movieID}/genres/{genreID} 	[delete]
func (api *API) handleDeleteMovieGenre(w http.ResponseWriter, r *http.Request) {
	movieID, genreID, err := api.handleDoubleIdentifierParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	err = api.CategoryService.DeleteMovieGenre(r.Context(), movieID, genreID)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrCategoryNotFound):
			api.handleErrorResponse(w, http.StatusNotFound, err)
		default:
			api.handleErrorResponse(w, http.StatusInternalServerError, err)
		}
		return
	}
	api.handleJsonResponse(w, http.StatusNoContent, nil)
}

// @Summary    	Lista géneros de una película paginados
// @Description Obtiene la lista de géneros asociados a una película dada por ID, con paginación.
// @Tags       	Categories
// @Accept     	json
// @Produce    	json
// @Param		page   query	 int			true	"Número de página"
// @Param       rows   query	 int			true	"Cantidad de filas por página"
// @Param      	id     path      int   			true   	"ID de la película"
// @Success    	200    {array}   service.Genre 			"Lista de géneros exitosa"
// @Failure    	400    {object}  JsonError  			"Error de validación de los parámetros"
// @Failure    	500    {object}  JsonError  			"Error interno del servidor"
// @Router     	/movies/{id}/genres 			[get]
func (api *API) handleGetMovieGenres(w http.ResponseWriter, r *http.Request) {
	movieID, err := api.handleSingleIdentifierParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	page, rows, err := api.handlePageAndRowsParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	genres, err := api.GenreService.GetMovieGenresList(r.Context(), movieID, page, rows)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrMovieNotFound):
			api.handleErrorResponse(w, http.StatusBadRequest, err)
		default:
			api.handleErrorResponse(w, http.StatusInternalServerError, err)
		}
		return
	}
	api.handleJsonResponse(w, http.StatusOK, genres)
}

// @Summary    	Lista películas de un género paginadas
// @Description Obtiene la lista de películas asociadas a un género dado por ID, con paginación.
// @Tags       	Categories
// @Accept     	json
// @Produce    	json
// @Param		page   query	 int			true	"Número de página"
// @Param       rows   query	 int			true	"Cantidad de filas por página"
// @Param      	id     path      int   			true   	"ID del género"
// @Success    	200    {array}   service.Genre 			"Lista de películas exitosa"
// @Failure    	400    {object}  JsonError  			"Error de validación de los parámetros"
// @Failure    	500    {object}  JsonError  			"Error interno del servidor"
// @Router     	/genres/{id}/movies 			[get]
func (api *API) handleGetGenreMovies(w http.ResponseWriter, r *http.Request) {
	genreID, err := api.handleSingleIdentifierParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	page, rows, err := api.handlePageAndRowsParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	movies, err := api.MovieService.GetGenreMoviesList(r.Context(), genreID, page, rows)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrGenreNotFound):
			api.handleErrorResponse(w, http.StatusBadRequest, err)
		default:
			api.handleErrorResponse(w, http.StatusInternalServerError, err)
		}
		return
	}
	api.handleJsonResponse(w, http.StatusOK, movies)
}
