package api

import (
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
	instanceGetRequestTemplate(w, r, api, api.GenreService.GetGenresList)
}

// @Summary		Agrega un género
// @Description	Agrega un nuevo género al catálogo.
// @Tags		Genres
// @Accept		json
// @Produce		json
// @Param		genre	body		service.GenreData  true		"Datos del género"
// @Success     201		{object}	service.Genre				"Género agregado exitosamente"
// @Failure		400		{object}	JsonError 					"Error de validación de los parámetros"
// @Failure		409		{object}	JsonError 					"Error de duplicacion de recursos"
// @Failure		500		{object}	JsonError 					"Error interno del servidor"
// @Router		/genres [post]
func (api *API) handlePostGenre(w http.ResponseWriter, r *http.Request) {
	instancePostRequestTemplate(w, r, api, service.GenreData{}, api.GenreService.AddGenre)
}

// @Summary		Obtiene un género
// @Description	Obtiene toda la información asociada a un género dado por ID.
// @Tags		Genres
// @Accept		json
// @Produce		json
// @Param		ID		path		int		true	"ID del género a obtener"
// @Success     200		{object}	service.Genre	"Género obtenida exitosamente"
// @Failure     400		{object}	JsonError		"Error de validación de los parámetros"
// @Failure		404		{object}	JsonError	 	"Error por género no encontrado"
// @Failure		500		{object}	JsonError	 	"Error interno del servidor"
// @Router		/genres/{ID} 		[get]
func (api *API) handleGetGenre(w http.ResponseWriter, r *http.Request) {
	instanceGetIDRequestTemplate(w, r, api, api.GenreService.GetGenre)
}

// @Summary		Actualiza los datos de un género
// @Description	Actualiza toda la información de un género dado por ID. Todos los campos son requeridos.
// @Tags		Genres
// @Accept		json
// @Produce		json
// @Param		ID		path		int		      		  true	"ID del género a actualizar"
// @Param		genre	body		service.GenreData	  true	"Datos actualizables del género"
// @Success     200		{object}	service.Genre				"Género actualizado exitosamente"
// @Failure		400		{object}	JsonError	 				"Error de validación de los parámetros"
// @Failure		404		{object}	JsonError	 				"Error por género no encontrado"
// @Failure		409		{object}	JsonError	 				"Error de duplicación de recursos"
// @Failure		500		{object}	JsonError	 				"Error interno del servidor"
// @Router		/genres/{ID} 		[put]
func (api *API) handlePutGenre(w http.ResponseWriter, r *http.Request) {
	instancePutIDRequestTemplate(w, r, api, service.GenreData{}, api.GenreService.UpdateGenre)
}

// @Summary		Borra un género
// @Description	Borra un género dado por ID permanentemente del catálogo.
// @Tags		Genres
// @Accept		json
// @Produce		json
// @Param		ID		path   int	true		"ID del género a eliminar"
// @Success     204								"Género eliminado exitosamente"
// @Failure		400		{object}	JsonError 	"Error de validación de los parámetros"
// @Failure		404		{object}	JsonError 	"Error por género no encontrado"
// @Failure		500		{object}	JsonError 	"Error interno del servidor"
// @Router		/genres/{ID} 		[delete]
func (api *API) handleDeleteGenre(w http.ResponseWriter, r *http.Request) {
	instanceDeleteIDRequestTemplate(w, r, api, api.GenreService.DeleteGenre)
}
