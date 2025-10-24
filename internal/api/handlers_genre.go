package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Matias914/Web-Page/internal/service"
)

// @Summary		Lista géneros paginados
// @Description	Obtiene una lista de géneros, ordenada por fecha de estreno descendente, con paginación.
// @Tags		Genres
// @Accept		json
// @Produce		json
// @Param		page	query		int				true	"Número de página"
// @Param       rows	query		int				true	"Cantidad de filas por página"
// @Success     200		{array}		service.Genre			"Lista de géneros exitosa"
// @Failure		400		{object}	JsonError	 			"Error de validación de los parámetros"
// @Failure		500		{object}	JsonError	 			"Error interno del servidor"
// @Router		/genres [get]
func (api *API) handleGetGenres(w http.ResponseWriter, r *http.Request) {
	page, rows, err := api.handlePageAndRowsParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	genres, err := api.GenreService.GetGenresList(r.Context(), page, rows)
	if err != nil {
		api.handleErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	api.handleJsonResponse(w, http.StatusOK, genres)
}

// @Summary		Agrega un género
// @Description	Agrega un nuevo género al catálogo.
// @Tags		Genres
// @Accept		json
// @Produce		json
// @Param		movie	body		service.AddGenreInput  true	"Número de página"
// @Success     201		{object}	service.Genre				"Género agregado exitosamente"
// @Failure		400		{object}	JsonError 					"Error de validación de los parámetros"
// @Failure		409		{object}	JsonError 					"Error de duplicacion de recursos"
// @Failure		500		{object}	JsonError 					"Error interno del servidor"
// @Router		/genres [post]
func (api *API) handlePostGenre(w http.ResponseWriter, r *http.Request) {
	var genre service.AddGenreInput
	if err := json.NewDecoder(r.Body).Decode(&genre); err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	if err := api.ValidationService.Validate(genre); err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	result, err := api.GenreService.AddGenre(r.Context(), genre)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidGenre):
			api.handleErrorResponse(w, http.StatusBadRequest, err)
		case errors.Is(err, service.ErrGenreDuplicated):
			api.handleErrorResponse(w, http.StatusConflict, err)
		default:
			api.handleErrorResponse(w, http.StatusInternalServerError, err)
		}
		return
	}
	api.handleJsonResponse(w, http.StatusCreated, result)
}

// @Summary		Obtiene un género
// @Description	Obtiene toda la información asociada a un género dado por ID.
// @Tags		Genres
// @Accept		json
// @Produce		json
// @Param		id		path		int		true	"ID del género a obtener"
// @Success     200		{object}	service.Genre	"Género obtenida exitosamente"
// @Failure     400		{object}	JsonError		"Error de validación de los parámetros"
// @Failure		404		{object}	JsonError	 	"Error por género no encontrado"
// @Failure		500		{object}	JsonError	 	"Error interno del servidor"
// @Router		/genres/{id} 		[get]
func (api *API) handleGetGenre(w http.ResponseWriter, r *http.Request) {
	id, err := api.handleSingleIdentifierParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	genre, err := api.GenreService.GetGenre(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrGenreNotFound):
			api.handleErrorResponse(w, http.StatusNotFound, err)
		default:
			api.handleErrorResponse(w, http.StatusInternalServerError, err)
		}
		return
	}
	api.handleJsonResponse(w, http.StatusOK, genre)
}

// @Summary		Actualiza los datos de un género
// @Description	Actualiza toda la información de un género dado por ID. Todos los campos son requeridos.
// @Tags		Genres
// @Accept		json
// @Produce		json
// @Param		id		path		int		      			  true	"ID del género a actualizar"
// @Param		movie	body		service.UpdateGenreInput  true	"Número de página"
// @Success     200		{object}	service.Genre					"Género actualizada exitosamente"
// @Failure		400		{object}	JsonError	 					"Error de validación de los parámetros"
// @Failure		404		{object}	JsonError	 					"Error por género no encontrado"
// @Failure		409		{object}	JsonError	 					"Error de duplicación de recursos"
// @Failure		500		{object}	JsonError	 					"Error interno del servidor"
// @Router		/genres/{id} 		[put]
func (api *API) handlePutGenre(w http.ResponseWriter, r *http.Request) {
	id, err := api.handleSingleIdentifierParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	var genre service.UpdateGenreInput
	if err := json.NewDecoder(r.Body).Decode(&genre); err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	if err := api.ValidationService.Validate(genre); err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	result, err := api.GenreService.UpdateGenre(r.Context(), id, genre)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidGenre):
			api.handleErrorResponse(w, http.StatusBadRequest, err)
		case errors.Is(err, service.ErrGenreDuplicated):
			api.handleErrorResponse(w, http.StatusConflict, err)
		case errors.Is(err, service.ErrGenreNotFound):
			api.handleErrorResponse(w, http.StatusNotFound, err)
		default:
			api.handleErrorResponse(w, http.StatusInternalServerError, err)
		}
		return
	}
	api.handleJsonResponse(w, http.StatusOK, result)
}

// @Summary		Borra un género
// @Description	Borra un género dado por ID permanentemente del catálogo.
// @Tags		Genres
// @Accept		json
// @Produce		json
// @Param		id		path   int	true		"ID del género a eliminar"
// @Success     204								"Género eliminada exitosamente"
// @Failure		400		{object}	JsonError 	"Error de validación de los parámetros"
// @Failure		404		{object}	JsonError 	"Error por género no encontrado"
// @Failure		500		{object}	JsonError 	"Error interno del servidor"
// @Router		/genres/{id} 		[delete]
func (api *API) handleDeleteGenre(w http.ResponseWriter, r *http.Request) {
	id, err := api.handleSingleIdentifierParsing(r)
	if err != nil {
		api.handleErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	err = api.GenreService.DeleteGenre(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrGenreReferenced):
			api.handleErrorResponse(w, http.StatusBadRequest, err)
		case errors.Is(err, service.ErrGenreNotFound):
			api.handleErrorResponse(w, http.StatusNotFound, err)
		default:
			api.handleErrorResponse(w, http.StatusInternalServerError, err)
		}
		return
	}
	api.handleJsonResponse(w, http.StatusNoContent, nil)
}
