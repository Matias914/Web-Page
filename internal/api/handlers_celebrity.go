package api

import (
	"net/http"

	"github.com/Matias914/Web-Page/internal/service"
)

// @Summary		Lista celebridades paginadas
// @Description	Obtiene una lista de celebridades con paginación.
// @Tags		Celebrities
// @Accept		json
// @Produce		json
// @Param		page	query		int				true	"Número de página"
// @Param       rows	query		int				true	"Cantidad de filas por página"
// @Success     200		{array}		service.Celebrity		"Lista de celebridads exitosa"
// @Failure		400		{object}	JsonError	 			"Error de validación de los parámetros"
// @Failure		500		{object}	JsonError	 			"Error interno del servidor"
// @Router		/celebrities 		[get]
func (api *API) handleGetCelebrities(w http.ResponseWriter, r *http.Request) {
	instanceGetRequestTemplate(w, r, api, api.CelebrityService.GetCelebritiesList)
}

// @Summary		Agrega una celebridad
// @Description	Agrega una nueva celebridad al catálogo.
// @Tags		Celebrities
// @Accept		json
// @Produce		json
// @Param		celebrity	body		service.CelebrityData  true		"Datos de la celebridad"
// @Success     201			{object}	service.Celebrity				"Celebridad agregada exitosamente"
// @Failure		400			{object}	JsonError 						"Error de validación de los parámetros"
// @Failure		409			{object}	JsonError 						"Error de duplicacion de recursos"
// @Failure		500			{object}	JsonError 						"Error interno del servidor"
// @Router		/celebrities 			[post]
func (api *API) handlePostCelebrity(w http.ResponseWriter, r *http.Request) {
	instancePostRequestTemplate(w, r, api, service.CelebrityData{}, api.CelebrityService.AddCelebrity)
}

// @Summary		Obtiene una celebridad
// @Description	Obtiene toda la información asociada a una celebridad dada por ID.
// @Tags		Celebrities
// @Accept		json
// @Produce		json
// @Param		ID		path		int		true		"ID de la celebridad a obtener"
// @Success     200		{object}	service.Celebrity	"Celebridad obtenida exitosamente"
// @Failure     400		{object}	JsonError			"Error de validación de los parámetros"
// @Failure		404		{object}	JsonError	 		"Error por celebridad no encontrada"
// @Failure		500		{object}	JsonError	 		"Error interno del servidor"
// @Router		/celebrities/{ID} 	[get]
func (api *API) handleGetCeleberity(w http.ResponseWriter, r *http.Request) {
	instanceGetIDRequestTemplate(w, r, api, api.CelebrityService.GetCelebrity)
}

// @Summary		Actualiza los datos de una celebridad
// @Description	Actualiza toda la información de una celebridad dado por ID. Todos los campos son requeridos.
// @Tags		Celebrities
// @Accept		json
// @Produce		json
// @Param		ID			path		int		      			true	"ID de la celebridad a actualizar"
// @Param		celebrity	body		service.CelebrityData 	true	"Datos actualizables de la celebridad"
// @Success     200			{object}	service.Celebrity				"Celebridad actualizada exitosamente"
// @Failure		400			{object}	JsonError	 					"Error de validación de los parámetros"
// @Failure		404			{object}	JsonError	 					"Error por celebridad no encontrada"
// @Failure		409			{object}	JsonError	 					"Error de duplicación de recursos"
// @Failure		500			{object}	JsonError	 					"Error interno del servidor"
// @Router		/celebrities/{ID} 		[put]
func (api *API) handlePutCelebrity(w http.ResponseWriter, r *http.Request) {
	instancePutIDRequestTemplate(w, r, api, service.CelebrityData{}, api.CelebrityService.UpdateCelebrity)
}

// @Summary		Borra una celebridad
// @Description	Borra una celebridad dada por ID permanentemente de la base de datos.
// @Tags		Celebrities
// @Accept		json
// @Produce		json
// @Param		ID		path   int	true		"ID de la celebridad a eliminar"
// @Success     204								"Celebridad eliminada exitosamente"
// @Failure		400		{object}	JsonError 	"Error de validación de los parámetros"
// @Failure		404		{object}	JsonError 	"Error por celebridad no encontrada"
// @Failure		500		{object}	JsonError 	"Error interno del servidor"
// @Router		/celebrities/{ID} 	[delete]
func (api *API) handleDeleteCelebrity(w http.ResponseWriter, r *http.Request) {
	instanceDeleteIDRequestTemplate(w, r, api, api.CelebrityService.DeleteCelebrity)
}
