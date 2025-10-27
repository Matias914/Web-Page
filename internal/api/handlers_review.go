package api

import (
	"net/http"

	"github.com/Matias914/Web-Page/internal/service"
)

// @Summary    	Agrega una review de un usuario sobre una película
// @Description Agrega una review de un usuario dado por ID a una película también dada por ID.
// @Tags       	Reviews
// @Accept     	json
// @Produce    	json
// @Param      	userID  path      int      				true   	"ID del usuario"
// @Param      	review	body      service.ReviewData	true   	"Datos de la review"
// @Success    	201    	{object}  service.Review				"Review agregada exitosamente"
// @Failure    	400    	{object}  JsonError  					"Error de validación de los parámetros"
// @Failure    	404    	{object}  JsonError  					"Error por review no encontrada"
// @Failure    	409    	{object}  JsonError  					"Error de duplicación de recursos
// @Failure    	500    	{object}  JsonError  					"Error interno del servidor"
// @Router     	/users/{userID}/reviews         		[post]
func (api *API) handlePostReview(w http.ResponseWriter, r *http.Request) {
	instancePostWithIDRequestTemplate(w, r, api, service.ReviewData{}, api.ReviewService.AddReview)
}

// @Summary    	Obtiene una review de un usuario sobre una película
// @Description Obtiene una review de un usuario dado por ID sobre una película también dada por ID.
// @Tags       	Reviews
// @Accept     	json
// @Produce    	json
// @Param      	userID    	path     	int      		true   	"ID del usuario"
// @Param      	movieID    	path     	int     		true   	"ID de la película"
// @Success    	200    		{object}	service.Review			"Review obtenida exitosamente"
// @Success    	400    		{object} 	JsonError  				"Error de validación de los parámetros"
// @Failure    	404    		{object}  	JsonError		  		"Error por review no encontrada"
// @Failure    	500    		{object}  	JsonError  				"Error interno del servidor"
// @Router     	/users/{userID}/reviews/{movieID} 		[get]
func (api *API) handleGetReview(w http.ResponseWriter, r *http.Request) {
	instanceGetDoubleIDRequestTemplate(w, r, api, api.ReviewService.GetReview)
}

// @Summary    	Actualiza una review de un usuario sobre una película
// @Description Actualiza una review de un usuario dado por ID sobre una película dada también por ID.
// @Tags       	Reviews
// @Accept     	json
// @Produce    	json
// @Param      	userID    	path      int  							true   	"ID del usuario"
// @Param      	movieID    	path      int  							true   	"ID de la película"
// @Param      	review    	body      service.UpdatableReviewData   true   	"Datos actualizables de la review"
// @Success    	200    		{object}  service.Review						"Review actualizada exitosamente"
// @Success    	400    		{object}  JsonError  							"Error de validación de los parámetros"
// @Failure    	404    		{object}  JsonError		  						"Error por review no encontrada"
// @Failure    	500    		{object}  JsonError  							"Error interno del servidor"
// @Router     	/users/{userID}/reviews/{movieID} 					[put]
func (api *API) handlePutReview(w http.ResponseWriter, r *http.Request) {
	instancePutDoubleIDRequestTemplate(w, r, api, service.UpdatableReviewData{}, api.ReviewService.UpdateReview)
}

// @Summary    	Borra una review de un usuario sobre una película
// @Description Borra una review de un usuario dado por ID sobre una película dada también por ID permanentemente.
// @Tags       	Reviews
// @Accept     	json
// @Produce    	json
// @Param      	userID    	path      int  			true   	"ID del usuario"
// @Param      	movieID    	path      int       	true   	"ID de la película"
// @Success    	204    							 			"Review borrada exitosamente"
// @Success    	400    		{object}  JsonError  			"Error de validación de los parámetros"
// @Failure    	404    		{object}  JsonError		  		"Error por review no encontrada"
// @Failure    	500    		{object}  JsonError  			"Error interno del servidor"
// @Router     	/users/{userID}/reviews/{movieID} 	[delete]
func (api *API) handleDeleteReview(w http.ResponseWriter, r *http.Request) {
	instanceDeleteDoubleIDRequestTemplate(w, r, api, api.ReviewService.DeleteReview)
}

// @Summary    	Lista los reviews de un usuario paginados
// @Description Obtiene la lista de reviews asociados a un usuario dada por ID, con paginación.
// @Tags       	Reviews
// @Accept     	json
// @Produce    	json
// @Param		page   query	 int			true	"Número de página"
// @Param       rows   query	 int			true	"Cantidad de filas por página"
// @Param      	userID path      int   			true   	"ID del usuario"
// @Success    	200    {array}   service.Review			"Lista de reviews exitosa"
// @Failure    	400    {object}  JsonError  			"Error de validación de los parámetros"
// @Failure    	500    {object}  JsonError  			"Error interno del servidor"
// @Router     	/users/{userID}/reviews			[get]
func (api *API) handleGetUserReviews(w http.ResponseWriter, r *http.Request) {
	instanceGetWithIDRequestTemplate(w, r, api, api.ReviewService.GetUserReviewsList)
}

// @Summary    	Lista los reviews de una película paginados
// @Description Obtiene la lista de reviews asociados a una película dada por ID, con paginación.
// @Tags       	Reviews
// @Accept     	json
// @Produce    	json
// @Param		page   		query	  int				true	"Número de página"
// @Param       rows   		query	  int				true	"Cantidad de filas por página"
// @Param      	movieID 	path      int   			true   	"ID de la película"
// @Success    	200    		{array}   service.Review			"Lista de reviews exitosa"
// @Failure    	400    		{object}  JsonError  				"Error de validación de los parámetros"
// @Failure    	500    		{object}  JsonError  				"Error interno del servidor"
// @Router     	/movies/{movieID}/reviews				[get]
func (api *API) handleGetMovieReviews(w http.ResponseWriter, r *http.Request) {
	instanceGetWithIDRequestTemplate(w, r, api, api.ReviewService.GetMovieReviewsList)
}
