package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Matias914/Web-Page/internal/service"
)

// @Summary		Lista películas paginadas
// @Description	Obtiene una lista de películas, ordenada por fecha de estreno descendente, con paginación.
// @Tags		Movies
// @Accept		json
// @Produce		json
// @Param		page	query		int		true	"Número de página"
// @Param       rows	query		int		true	"Cantidad de filas por página"
// @Success     200		{array}		service.Movie	"Lista de películas exitosa"
// @Failure		400		{object}	JsonError	 	"Error de validación de los parámetros"
// @Failure		500		{object}	JsonError	 	"Error interno del servidor"
// @Router		/movies [get]
func (api *API) handleGetMovies(w http.ResponseWriter, r *http.Request) {
	page, rows, err := api.handlePageAndRowsParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	movies, err := api.MovieService.GetMoviesList(r.Context(), page, rows)
	if err != nil {
		api.handleErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	api.handleJsonResponse(w, http.StatusOK, movies)
}

// @Summary		Agrega una película
// @Description	Agrega una nueva película al catálogo. La URL del póster puede omitirse.
// @Tags		Movies
// @Accept		json
// @Produce		json
// @Param		movie	body		service.AddMovieInput  true	"Número de página"
// @Success     201		{object}	service.Movie				"Película agregada exitosamente"
// @Failure		400		{object}	JsonError 					"Error de validación de los parámetros"
// @Failure		409		{object}	JsonError 					"Error de duplicacion de recursos"
// @Failure		500		{object}	JsonError 					"Error interno del servidor"
// @Router		/movies [post]
func (api *API) handlePostMovie(w http.ResponseWriter, r *http.Request) {
	var movie service.AddMovieInput
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	if err := api.ValidationService.Validate(movie); err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	result, err := api.MovieService.AddMovie(r.Context(), movie)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidMovie):
			api.handleErrorResponse(w, http.StatusBadRequest, err)
		case errors.Is(err, service.ErrMovieDuplicated):
			api.handleErrorResponse(w, http.StatusConflict, err)
		default:
			api.handleErrorResponse(w, http.StatusInternalServerError, err)
		}
		return
	}
	api.handleJsonResponse(w, http.StatusCreated, result)
}

// @Summary		Obtiene una película
// @Description	Obtiene toda la información asociada a una película dada por ID.
// @Tags		Movies
// @Accept		json
// @Produce		json
// @Param		id		path		int		true	"ID de la película a obtener"
// @Success     200		{object}	service.Movie	"Película obtenida exitosamente"
// @Failure     400		{object}	JsonError		"Error de validación de los parámetros"
// @Failure		404		{object}	JsonError	 	"Error por película no encontrada"
// @Failure		500		{object}	JsonError	 	"Error interno del servidor"
// @Router		/movies/{id} 		[get]
func (api *API) handleGetMovie(w http.ResponseWriter, r *http.Request) {
	id, err := api.handleSingleIdentifierParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	movie, err := api.MovieService.GetMovie(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrMovieNotFound):
			api.handleErrorResponse(w, http.StatusNotFound, err)
		default:
			api.handleErrorResponse(w, http.StatusInternalServerError, err)
		}
		return
	}
	api.handleJsonResponse(w, http.StatusOK, movie)
}

// @Summary		Actualiza los datos de una película
// @Description	Actualiza toda la información de una película dada por ID. Todos los campos son requeridos.
// @Tags		Movies
// @Accept		json
// @Produce		json
// @Param		id		path		int		      			  true	"ID de la película a actualizar"
// @Param		movie	body		service.UpdateMovieInput  true	"Número de página"
// @Success     200		{object}	service.Movie					"Película actualizada exitosamente"
// @Failure		400		{object}	JsonError	 					"Error de validación de los parámetros"
// @Failure		404		{object}	JsonError	 					"Error por película no encontrada"
// @Failure		409		{object}	JsonError	 					"Error de duplicación de recursos"
// @Failure		500		{object}	JsonError	 					"Error interno del servidor"
// @Router		/movies/{id} 		[put]
func (api *API) handlePutMovie(w http.ResponseWriter, r *http.Request) {
	id, err := api.handleSingleIdentifierParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	var movie service.UpdateMovieInput
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	if err := api.ValidationService.Validate(movie); err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	result, err := api.MovieService.UpdateMovie(r.Context(), id, movie)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidMovie):
			api.handleErrorResponse(w, http.StatusBadRequest, err)
		case errors.Is(err, service.ErrMovieDuplicated):
			api.handleErrorResponse(w, http.StatusConflict, err)
		case errors.Is(err, service.ErrMovieNotFound):
			api.handleErrorResponse(w, http.StatusNotFound, err)
		default:
			api.handleErrorResponse(w, http.StatusInternalServerError, err)
		}
		return
	}
	api.handleJsonResponse(w, http.StatusOK, result)
}

// @Summary		Borra una película
// @Description	Borra una película dada por ID permanentemente del catálogo.
// @Tags		Movies
// @Accept		json
// @Produce		json
// @Param		id		path   int	true		"ID de la película a eliminar"
// @Success     204								"Película eliminada exitosamente"
// @Failure		400		{object}	JsonError 	"Error de validación de los parámetros"
// @Failure		404		{object}	JsonError 	"Error por película no encontrada"
// @Failure		500		{object}	JsonError 	"Error interno del servidor"
// @Router		/movies/{id} 		[delete]
func (api *API) handleDeleteMovie(w http.ResponseWriter, r *http.Request) {
	id, err := api.handleSingleIdentifierParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	err = api.MovieService.DeleteMovie(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrMovieReferenced):
			api.handleErrorResponse(w, http.StatusBadRequest, err)
		case errors.Is(err, service.ErrMovieNotFound):
			api.handleErrorResponse(w, http.StatusNotFound, err)
		default:
			api.handleErrorResponse(w, http.StatusInternalServerError, err)
		}
		return
	}
	api.handleJsonResponse(w, http.StatusNoContent, nil)
}
