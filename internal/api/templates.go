package api

import (
	"context"
	"encoding/json"
	"net/http"
)

// instanceGetRequestTemplate consiste en una plantilla que busca una lista de estructuras con
// un servicio. Su invocación instancia una GET request que busca leer una colección de datos
// paginados.
func instanceGetRequestTemplate[OUTPUT any](
	w http.ResponseWriter,
	r *http.Request,
	api *API, service func(context.Context, int, int) ([]OUTPUT, error),
) {
	page, rows, err := api.handlePageAndRowsParsing(r)
	if err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	result, err := service(r.Context(), page, rows)
	if err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	api.handleJsonResponse(w, http.StatusOK, result)
}

// instanceGetWithIDRequestTemplate consiste en una plantilla que busca una lista de estructuras
// con un servicio. Su invocación instancia una GET request con un path ID que busca leer una
// colección de datos paginados.
func instanceGetWithIDRequestTemplate[OUTPUT any](
	w http.ResponseWriter,
	r *http.Request,
	api *API, service func(context.Context, int, int, int) ([]OUTPUT, error),
) {
	id, err := api.handleSingleIdentifierParsing(r)
	if err != nil {
		api.handleErrorResponse(w, err)
	}
	page, rows, err := api.handlePageAndRowsParsing(r)
	if err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	result, err := service(r.Context(), id, page, rows)
	if err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	api.handleJsonResponse(w, http.StatusOK, result)
}

// instancePostRequestTemplate consiste en una plantilla que inserta una estructuras con un servicio.
// Su invocación instancia una POST request que busca agregar un elemento a una colección de datos.
func instancePostRequestTemplate[INPUT any, OUTPUT any](
	w http.ResponseWriter,
	r *http.Request,
	api *API,
	input INPUT,
	service func(context.Context, INPUT) (OUTPUT, error),
) {
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	if err := api.ValidationService.Validate(input); err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	result, err := service(r.Context(), input)
	if err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	api.handleJsonResponse(w, http.StatusCreated, result)
}

// instancePostWithIDRequestTemplate consiste en una plantilla que inserta una estructuras con un servicio.
// Su invocación instancia una POST request con un path ID que busca agregar un elemento a una colección de
// datos.
func instancePostWithIDRequestTemplate[INPUT any, OUTPUT any](
	w http.ResponseWriter,
	r *http.Request,
	api *API,
	input INPUT,
	service func(context.Context, int, INPUT) (OUTPUT, error),
) {
	id, err := api.handleSingleIdentifierParsing(r)
	if err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	if err := api.ValidationService.Validate(input); err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	result, err := service(r.Context(), id, input)
	if err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	api.handleJsonResponse(w, http.StatusCreated, result)
}

// instanceGetIDRequestTemplate consiste en una plantilla que busca una estructura con un servicio.
// Su invocación instancia una GET request con un path ID que busca un elemento de una colección de
// datos.
func instanceGetIDRequestTemplate[OUTPUT any](
	w http.ResponseWriter,
	r *http.Request,
	api *API,
	service func(context.Context, int) (OUTPUT, error),
) {
	id, err := api.handleSingleIdentifierParsing(r)
	if err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	result, err := service(r.Context(), id)
	if err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	api.handleJsonResponse(w, http.StatusOK, result)
}

// instanceGetDoubleIDRequestTemplate consiste en una plantilla que busca una estructura con un servicio.
// Su invocación instancia una GET request con dos path ID que busca un elemento de una colección de
// datos. El orden de los ID es el mismo que el de los parámetros de invoacación.
func instanceGetDoubleIDRequestTemplate[OUTPUT any](
	w http.ResponseWriter,
	r *http.Request,
	api *API,
	service func(context.Context, int, int) (OUTPUT, error),
) {
	id1, id2, err := api.handleDoubleIdentifierParsing(r)
	if err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	result, err := service(r.Context(), id1, id2)
	if err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	api.handleJsonResponse(w, http.StatusOK, result)
}

// instancePutIDRequestTemplate consiste en una plantilla que actualiza una estructura con un servicio.
// Su invocación instancia un PUT request con un path ID que busca agregar un elemento a una colección
// de datos.
func instancePutIDRequestTemplate[INPUT any, OUTPUT any](
	w http.ResponseWriter,
	r *http.Request,
	api *API,
	input INPUT,
	service func(context.Context, int, INPUT) (OUTPUT, error),
) {
	id, err := api.handleSingleIdentifierParsing(r)
	if err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	if err := api.ValidationService.Validate(input); err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	result, err := service(r.Context(), id, input)
	if err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	api.handleJsonResponse(w, http.StatusOK, result)
}

// instancePutDoubleIDRequestTemplate consiste en una plantilla que actualiza una estructura con un servicio.
// Su invocación instancia un PUT request con dos path ID que busca agregar un elemento a una colección de
// datos. El orden de los ID es el mismo que el de los parámetros de invoacación.
func instancePutDoubleIDRequestTemplate[INPUT any, OUTPUT any](
	w http.ResponseWriter,
	r *http.Request,
	api *API,
	input INPUT,
	service func(context.Context, int, int, INPUT) (OUTPUT, error),
) {
	id1, id2, err := api.handleDoubleIdentifierParsing(r)
	if err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	if err := api.ValidationService.Validate(input); err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	result, err := service(r.Context(), id1, id2, input)
	if err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	api.handleJsonResponse(w, http.StatusOK, result)
}

// instanceDeleteIDRequestTemplate consiste en una plantilla que elimina una estructura con un servicio.
// Su invocación instancia un DELETE request con un path ID que busca eliminar un elemento a una colección
// de datos.
func instanceDeleteIDRequestTemplate(
	w http.ResponseWriter,
	r *http.Request,
	api *API,
	service func(context.Context, int) error,
) {
	id, err := api.handleSingleIdentifierParsing(r)
	if err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	err = service(r.Context(), id)
	if err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	api.handleJsonResponse(w, http.StatusNoContent, nil)
}

// instanceDeleteDoubleIDRequestTemplate consiste en una plantilla que elimina una estructura con un servicio.
// Su invocación instancia un DELETE request con dos path ID que busca eliminar un elemento a una colección
// de datos. El orden de los ID es el mismo que el de los parámetros de invoacación.
func instanceDeleteDoubleIDRequestTemplate(
	w http.ResponseWriter,
	r *http.Request,
	api *API,
	service func(context.Context, int, int) error,
) {
	id1, id2, err := api.handleDoubleIdentifierParsing(r)
	if err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	err = service(r.Context(), id1, id2)
	if err != nil {
		api.handleErrorResponse(w, err)
		return
	}
	api.handleJsonResponse(w, http.StatusNoContent, nil)
}
