package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

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

func (api *API) handleSingleIdentifierParsing(r *http.Request) (int, error) {
	id, err := strconv.Atoi(chi.URLParam(r, "id1"))
	if err != nil || id < 1 {
		return -1, errors.New("invalid path identifier")
	}
	return id, nil
}

func (api *API) handleDoubleIdentifierParsing(r *http.Request) (int, int, error) {
	id1, err := api.handleSingleIdentifierParsing(r)
	if err != nil {
		return -1, -1, err
	}
	id2, err := strconv.Atoi(chi.URLParam(r, "id2"))
	if err != nil || id2 < 1 {
		return id1, -1, errors.New("invalid path identifier")
	}
	return id1, id2, nil
}

func (api *API) handleTripleIdentifierParsing(r *http.Request) (int, int, int, error) {
	id1, id2, err := api.handleDoubleIdentifierParsing(r)
	if err != nil {
		return id1, id2, -1, err
	}
	id3, err := strconv.Atoi(chi.URLParam(r, "id3"))
	if err != nil || id3 < 1 {
		return id1, id2, -1, errors.New("invalid path identifier")
	}
	return id1, id2, id3, nil
}
