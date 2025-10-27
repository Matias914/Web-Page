package api

import (
	"net/http"

	"github.com/Matias914/Web-Page/internal/service"
)

// @Summary    	Agrega un rating de un usuario sobre una película
// @Description Agrega un rating de un usuario dado por ID a una película también dada por ID.
// @Tags       	Ratings
// @Accept     	json
// @Produce    	json
// @Param      	userID  path      int      				true   	"ID del usuario"
// @Param      	rating	body      service.RatingData	true   	"Datos del rating"
// @Success    	201    	{object}  service.Rating				"Rating agregada exitosamente"
// @Failure    	400    	{object}  JsonError  					"Error de validación de los parámetros"
// @Failure    	404    	{object}  JsonError  					"Error por rating no encontrada"
// @Failure    	409    	{object}  JsonError  					"Error de duplicación de recursos
// @Failure    	500    	{object}  JsonError  					"Error interno del servidor"
// @Router     	/users/{userID}/ratings         		[post]
func (api *API) handlePostRating(w http.ResponseWriter, r *http.Request) {
	instancePostWithIDRequestTemplate(w, r, api, service.RatingData{}, api.RatingService.AddRating)
}

// @Summary    	Obtiene un rating de un usuario sobre una película
// @Description Obtiene un rating de un usuario dado por ID sobre una película también dada por ID.
// @Tags       	Ratings
// @Accept     	json
// @Produce    	json
// @Param      	userID    	path     	int      		true   	"ID del usuario"
// @Param      	movieID    	path     	int     		true   	"ID de la película"
// @Success    	200    		{object}	service.Rating			"Rating obtenida exitosamente"
// @Success    	400    		{object} 	JsonError  				"Error de validación de los parámetros"
// @Failure    	404    		{object}  	JsonError		  		"Error por rating no encontrada"
// @Failure    	500    		{object}  	JsonError  				"Error interno del servidor"
// @Router     	/users/{userID}/ratings/{movieID} 		[get]
func (api *API) handleGetRating(w http.ResponseWriter, r *http.Request) {
	instanceGetDoubleIDRequestTemplate(w, r, api, api.RatingService.GetRating)
}

// @Summary    	Actualiza un rating de un usuario sobre una película
// @Description Actualiza un rating de un usuario dado por ID sobre una película dada también por ID.
// @Tags       	Ratings
// @Accept     	json
// @Produce    	json
// @Param      	userID    	path      int  							true   	"ID del usuario"
// @Param      	movieID 	path      int      						true   	"ID de la película"
// @Param      	rating    	body      service.UpdatableRatingData   true   	"Datos actualizables del rating"
// @Success    	200    		{object}  service.Rating						"Rating actualizada exitosamente"
// @Success    	400    		{object}  JsonError  							"Error de validación de los parámetros"
// @Failure    	404    		{object}  JsonError		  						"Error por rating no encontrada"
// @Failure    	500    		{object}  JsonError  							"Error interno del servidor"
// @Router     	/users/{userID}/ratings/{movieID} 					[put]
func (api *API) handlePutRating(w http.ResponseWriter, r *http.Request) {
	instancePutDoubleIDRequestTemplate(w, r, api, service.UpdatableRatingData{}, api.RatingService.UpdateRating)
}

// @Summary    	Borra un rating de un usuario sobre una película
// @Description Borra un rating de un usuario dado por ID sobre una película dada también por ID permanentemente.
// @Tags       	Ratings
// @Accept     	json
// @Produce    	json
// @Param      	userID    	path      int  			true   	"ID del usuario"
// @Param      	movieID    	path      int       	true   	"ID de la película"
// @Success    	204    							 			"Rating borrada exitosamente"
// @Success    	400    		{object}  JsonError  			"Error de validación de los parámetros"
// @Failure    	404    		{object}  JsonError		  		"Error por rating no encontrada"
// @Failure    	500    		{object}  JsonError  			"Error interno del servidor"
// @Router     	/users/{userID}/ratings/{movieID} 	[delete]
func (api *API) handleDeleteRating(w http.ResponseWriter, r *http.Request) {
	instanceDeleteDoubleIDRequestTemplate(w, r, api, api.RatingService.DeleteRating)
}

// @Summary    	Lista los ratings de un usuario paginados
// @Description Obtiene la lista de ratings asociados a un usuario dada por ID, con paginación.
// @Tags       	Ratings
// @Accept     	json
// @Produce    	json
// @Param		page   query	 int			true	"Número de página"
// @Param       rows   query	 int			true	"Cantidad de filas por página"
// @Param      	userID path      int   			true   	"ID del usuario"
// @Success    	200    {array}   service.Rating			"Lista de ratings exitosa"
// @Failure    	400    {object}  JsonError  			"Error de validación de los parámetros"
// @Failure    	500    {object}  JsonError  			"Error interno del servidor"
// @Router     	/users/{userID}/ratings			[get]
func (api *API) handleGetUserRatings(w http.ResponseWriter, r *http.Request) {
	instanceGetWithIDRequestTemplate(w, r, api, api.RatingService.GetUserRatingsList)
}

// @Summary    	Lista los ratings de una película paginados
// @Description Obtiene la lista de ratings asociados a una película dada por ID, con paginación.
// @Tags       	Ratings
// @Accept     	json
// @Produce    	json
// @Param		page   		query	  int				true	"Número de página"
// @Param       rows   		query	  int				true	"Cantidad de filas por página"
// @Param      	movieID 	path      int   			true   	"ID de la película"
// @Success    	200    		{array}   service.Rating			"Lista de ratings exitosa"
// @Failure    	400    		{object}  JsonError  				"Error de validación de los parámetros"
// @Failure    	500    		{object}  JsonError  				"Error interno del servidor"
// @Router     	/movies/{movieID}/ratings				[get]
func (api *API) handleGetMovieRatings(w http.ResponseWriter, r *http.Request) {
	instanceGetWithIDRequestTemplate(w, r, api, api.RatingService.GetMovieRatingsList)
}
