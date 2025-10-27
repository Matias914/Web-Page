package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// handlePageAndRowsParsing esta funcion parsea los campos de página y filas de la URL,
// pero solo funciona si los mismos son llamados "page" y "rows".
func (api *API) handlePageAndRowsParsing(r *http.Request) (int, int, error) {
	pageStr := r.URL.Query().Get("page")
	rowsStr := r.URL.Query().Get("rows")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		return -1, -1, errors.New("invalid page number")
	}
	rows, err := strconv.Atoi(rowsStr)
	if err != nil || rows < 1 {
		return -1, -1, errors.New("invalid row number")
	}
	return page, rows, nil
}

// handleSingleIdentifierParsing esta funcion parsea el campo de identificador,
// pero solo funciona si el mismo es llamado "id1".
func (api *API) handleSingleIdentifierParsing(r *http.Request) (int, error) {
	id, err := strconv.Atoi(chi.URLParam(r, "id1"))
	if err != nil || id < 1 {
		return -1, errors.New("invalid first path identifier")
	}
	return id, nil
}

// handleDoubleIdentifierParsing esta funcion parsea el campo de identificador,
// pero solo funciona si los mismos son llamados "id1" e "id2".
func (api *API) handleDoubleIdentifierParsing(r *http.Request) (int, int, error) {
	id1, err := api.handleSingleIdentifierParsing(r)
	if err != nil {
		return -1, -1, err
	}
	id2, err := strconv.Atoi(chi.URLParam(r, "id2"))
	if err != nil || id2 < 1 {
		return id1, -1, errors.New("invalid second path identifier")
	}
	return id1, id2, nil
}
