package api

import (
	"net/http"

	"github.com/Matias914/Web-Page/internal/service"
)

// @Summary		Lista usuarios paginados
// @Description	Obtiene una lista de usuarios con paginación.
// @Tags		Users
// @Accept		json
// @Produce		json
// @Param		page	query		int				true	"Número de página"
// @Param       rows	query		int				true	"Cantidad de filas por página"
// @Success     200		{array}		service.User			"Lista de usuarios exitosa"
// @Failure		400		{object}	JsonError	 			"Error de validación de los parámetros"
// @Failure		500		{object}	JsonError	 			"Error interno del servidor"
// @Router		/users [get]
func (api *API) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	instanceGetRequestTemplate(w, r, api, api.UserService.GetUsersList)
}

// @Summary		Agrega un usuario
// @Description	Agrega un nuevo usuario a la base de datos.
// @Tags		Users
// @Accept		json
// @Produce		json
// @Param		user	body		service.UserData  true	"Datos del usuario"
// @Success     201		{object}	service.User			"Usuario agregado exitosamente"
// @Failure		400		{object}	JsonError 				"Error de validación de los parámetros"
// @Failure		409		{object}	JsonError 				"Error de duplicacion de recursos"
// @Failure		500		{object}	JsonError 				"Error interno del servidor"
// @Router		/users [post]
func (api *API) handlePostUser(w http.ResponseWriter, r *http.Request) {
	instancePostRequestTemplate(w, r, api, service.UserData{}, api.UserService.AddUser)
}

// @Summary		Obtiene un usuario
// @Description	Obtiene toda la información asociada a un usuario dado por ID.
// @Tags		Users
// @Accept		json
// @Produce		json
// @Param		ID		path		int		true	"ID del usuario a obtener"
// @Success     200		{object}	service.User	"Usuario obtenida exitosamente"
// @Failure     400		{object}	JsonError		"Error de validación de los parámetros"
// @Failure		404		{object}	JsonError	 	"Error por usuario no encontrado"
// @Failure		500		{object}	JsonError	 	"Error interno del servidor"
// @Router		/users/{ID} 		[get]
func (api *API) handleGetUser(w http.ResponseWriter, r *http.Request) {
	instanceGetIDRequestTemplate(w, r, api, api.UserService.GetUser)
}

// @Summary		Actualiza los datos de un usuario
// @Description	Actualiza toda la información de un usuario dado por ID. Todos los campos son requeridos.
// @Tags		Users
// @Accept		json
// @Produce		json
// @Param		ID		path		int		      				true	"ID del usuario a actualizar"
// @Param		user	body		service.UpdatableUserData   true	"Datos actualizables del usuario"
// @Success     200		{object}	service.User						"Usuario actualizado exitosamente"
// @Failure		400		{object}	JsonError	 						"Error de validación de los parámetros"
// @Failure		404		{object}	JsonError	 						"Error por usuario no encontrado"
// @Failure		409		{object}	JsonError	 						"Error de duplicación de recursos"
// @Failure		500		{object}	JsonError	 						"Error interno del servidor"
// @Router		/users/{ID} 		[put]
func (api *API) handlePutUser(w http.ResponseWriter, r *http.Request) {
	instancePutIDRequestTemplate(w, r, api, service.UpdatableUserData{}, api.UserService.UpdateUser)
}

// @Summary		Borra un usuario
// @Description	Borra un usuario dado por ID permanentemente de la base de datos.
// @Tags		Users
// @Accept		json
// @Produce		json
// @Param		ID		path   int	true		"ID del usuario a eliminar"
// @Success     204								"Usuario borrado exitosamente"
// @Failure		400		{object}	JsonError 	"Error de validación de los parámetros"
// @Failure		404		{object}	JsonError 	"Error por usuario no encontrado"
// @Failure		500		{object}	JsonError 	"Error interno del servidor"
// @Router		/users/{ID} 		[delete]
func (api *API) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	instanceDeleteIDRequestTemplate(w, r, api, api.UserService.DeleteUser)
}
