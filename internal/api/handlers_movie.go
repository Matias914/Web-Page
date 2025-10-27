package api

import (
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
	instanceGetRequestTemplate(w, r, api, api.MovieService.GetMoviesList)
}

// @Summary		Agrega una película
// @Description	Agrega una nueva película al catálogo. La URL del póster puede omitirse.
// @Tags		Movies
// @Accept		json
// @Produce		json
// @Param		movie	body		service.MovieData  true		"Datos de la película"
// @Success     201		{object}	service.Movie				"Película agregada exitosamente"
// @Failure		400		{object}	JsonError 					"Error de validación de los parámetros"
// @Failure		409		{object}	JsonError 					"Error de duplicacion de recursos"
// @Failure		500		{object}	JsonError 					"Error interno del servidor"
// @Router		/movies [post]
func (api *API) handlePostMovie(w http.ResponseWriter, r *http.Request) {
	instancePostRequestTemplate(w, r, api, service.MovieData{}, api.MovieService.AddMovie)
}

// @Summary		Obtiene una película
// @Description	Obtiene toda la información asociada a una película dada por ID.
// @Tags		Movies
// @Accept		json
// @Produce		json
// @Param		ID		path		int		true	"ID de la película a obtener"
// @Success     200		{object}	service.Movie	"Película obtenida exitosamente"
// @Failure     400		{object}	JsonError		"Error de validación de los parámetros"
// @Failure		404		{object}	JsonError	 	"Error por película no encontrada"
// @Failure		500		{object}	JsonError	 	"Error interno del servidor"
// @Router		/movies/{ID} 		[get]
func (api *API) handleGetMovie(w http.ResponseWriter, r *http.Request) {
	instanceGetIDRequestTemplate(w, r, api, api.MovieService.GetMovie)
}

// @Summary		Actualiza los datos de una película
// @Description	Actualiza toda la información de una película dada por ID. Todos los campos son requeridos.
// @Tags		Movies
// @Accept		json
// @Produce		json
// @Param		ID		path		int		      	   true	"ID de la película a actualizar"
// @Param		movie	body		service.MovieData  true	"Datos actualizables de la película"
// @Success     200		{object}	service.Movie			"Película actualizada exitosamente"
// @Failure		400		{object}	JsonError	 			"Error de validación de los parámetros"
// @Failure		404		{object}	JsonError	 			"Error por película no encontrada"
// @Failure		409		{object}	JsonError	 			"Error de duplicación de recursos"
// @Failure		500		{object}	JsonError	 			"Error interno del servidor"
// @Router		/movies/{ID} 		[put]
func (api *API) handlePutMovie(w http.ResponseWriter, r *http.Request) {
	instancePutIDRequestTemplate(w, r, api, service.MovieData{}, api.MovieService.UpdateMovie)
}

// @Summary		Borra una película
// @Description	Borra una película dada por ID permanentemente del catálogo.
// @Tags		Movies
// @Accept		json
// @Produce		json
// @Param		ID		path   int	true		"ID de la película a eliminar"
// @Success     204								"Película eliminada exitosamente"
// @Failure		400		{object}	JsonError 	"Error de validación de los parámetros"
// @Failure		404		{object}	JsonError 	"Error por película no encontrada"
// @Failure		500		{object}	JsonError 	"Error interno del servidor"
// @Router		/movies/{ID} 		[delete]
func (api *API) handleDeleteMovie(w http.ResponseWriter, r *http.Request) {
	instanceDeleteIDRequestTemplate(w, r, api, api.MovieService.DeleteMovie)
}
