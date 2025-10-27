package api

import (
	"net/http"

	"github.com/Matias914/Web-Page/internal/service"
)

// @Summary    	Agrega un rol de una celebridad en una película
// @Description Agrega un rol de una celebridad dada por ID a una película también dada por ID.
// @Tags       	Roles
// @Accept     	json
// @Produce    	json
// @Param      	movieID path      int      			true   	"ID de la película"
// @Param      	role	body      service.RoleData	true   	"Datos del rol"
// @Success    	201    	{object}  service.Role				"Rol agregado exitosamente"
// @Failure    	400    	{object}  JsonError  				"Error de validación de los parámetros"
// @Failure    	404    	{object}  JsonError  				"Error por rol no encontrado"
// @Failure    	409    	{object}  JsonError  				"Error de duplicación de recursos
// @Failure    	500    	{object}  JsonError  				"Error interno del servidor"
// @Router     	/movies/{movieID}/roles         	[post]
func (api *API) handlePostRole(w http.ResponseWriter, r *http.Request) {
	instancePostWithIDRequestTemplate(w, r, api, service.RoleData{}, api.RoleService.AddRole)
}

// @Summary    	Obtiene el rol de una celebridad en una película
// @Description Obtiene el rol de una celebridad dada por ID en una película también dada por ID.
// @Tags       	Roles
// @Accept     	json
// @Produce    	json
// @Param      	movieID    	path     	int      		true   	"ID de la película"
// @Param      	celebrityID path     	int     		true   	"ID de la celebridad"
// @Success    	200    		{array}		service.Role			"Lista de roles exitosa"
// @Success    	400    		{object} 	JsonError  				"Error de validación de los parámetros"
// @Failure    	404    		{object}  	JsonError		  		"Error por rol no encontrado"
// @Failure    	500    		{object}  	JsonError  				"Error interno del servidor"
// @Router     	/movies/{movieID}/roles/{celebrityID} 	[get]
func (api *API) handleGetRole(w http.ResponseWriter, r *http.Request) {
	instanceGetDoubleIDRequestTemplate(w, r, api, api.RoleService.GetRole)
}

// @Summary    	Actualiza un rol de una celebridad en una película
// @Description Actualiza un rol de una celebridad dada por ID en una película dada también por ID.
// @Tags       	Roles
// @Accept     	json
// @Produce    	json
// @Param      	movieID    	path      int  							true   	"ID de la película"
// @Param      	celebrityID path      int  							true   	"ID de la celebridad"
// @Param      	rol		    body      service.UpdatableRoleData   	true   	"Datos actualizables del rol"
// @Success    	200    		{object}  service.Role							"Rol actualizado exitosamente"
// @Success    	400    		{object}  JsonError  							"Error de validación de los parámetros"
// @Failure    	404    		{object}  JsonError		  						"Error por rating no encontrada"
// @Failure    	500    		{object}  JsonError  							"Error interno del servidor"
// @Router     	/movies/{movieID}/roles/{celebrityID} 				[put]
func (api *API) handlePutRole(w http.ResponseWriter, r *http.Request) {
	instancePutDoubleIDRequestTemplate(w, r, api, service.UpdatableRoleData{}, api.RoleService.UpdateRole)
}

// @Summary    	Borra un rol de una celebridad en una película
// @Description Borra un rol de una celebridad dada por ID en una película dada también por ID permanentemente.
// @Tags       	Roles
// @Accept     	json
// @Produce    	json
// @Param      	movieID    	path      int  			true   	"ID de la película"
// @Param      	celebrityID path      int       	true   	"ID de la celebridad"
// @Success    	204    							 			"Rol borrado exitosamente"
// @Success    	400    		{object}  JsonError  			"Error de validación de los parámetros"
// @Failure    	404    		{object}  JsonError		  		"Error por rol no encontrado"
// @Failure    	500    		{object}  JsonError  			"Error interno del servidor"
// @Router     	/movies/{movieID}/roles/{celebrityID} 		[delete]
func (api *API) handleDeleteRole(w http.ResponseWriter, r *http.Request) {
	instanceDeleteDoubleIDRequestTemplate(w, r, api, api.RoleService.DeleteRole)
}

// @Summary    	Lista los roles de una película paginados
// @Description Obtiene la lista de roles asociados a una película dada por ID, con paginación.
// @Tags       	Roles
// @Accept     	json
// @Produce    	json
// @Param		page   		query	  int				true	"Número de página"
// @Param       rows   		query	  int				true	"Cantidad de filas por página"
// @Param      	movieID 	path      int   			true   	"ID de la película"
// @Success    	200    		{array}   service.Role				"Lista de roles exitosa"
// @Failure    	400    		{object}  JsonError  				"Error de validación de los parámetros"
// @Failure    	500    		{object}  JsonError  				"Error interno del servidor"
// @Router     	/movies/{movieID}/roles					[get]
func (api *API) handleGetMovieRoles(w http.ResponseWriter, r *http.Request) {
	instanceGetWithIDRequestTemplate(w, r, api, api.RoleService.GetMovieRolesList)
}

// @Summary    	Lista los roles de una celebridad paginados
// @Description Obtiene la lista de roles asociados a una celebridad dada por ID, con paginación.
// @Tags       	Roles
// @Accept     	json
// @Produce    	json
// @Param		page   		query	 	int			true	"Número de página"
// @Param       rows   		query	 	int			true	"Cantidad de filas por página"
// @Param      	celebrityID path      	int   		true   	"ID de la celebridad"
// @Success    	200    		{array}   	service.Role		"Lista de roles exitosa"
// @Failure    	400    		{object}  	JsonError  			"Error de validación de los parámetros"
// @Failure    	500    		{object}  	JsonError  			"Error interno del servidor"
// @Router     	/celebrities/{celebrityID}/roles			[get]
func (api *API) handleGetCelebrityRoles(w http.ResponseWriter, r *http.Request) {
	instanceGetWithIDRequestTemplate(w, r, api, api.RoleService.GetCelebrityRolesList)
}
