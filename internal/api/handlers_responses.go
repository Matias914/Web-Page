package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func (api *API) handleErrorResponse(w http.ResponseWriter, status int, error error) {
	if status < 500 {
		content := JsonError{Mssg: error.Error()}
		api.handleJsonResponse(w, status, content)
	} else {
		content := JsonError{Mssg: "internal server error"}
		log.Println(error.Error())
		api.handleJsonResponse(w, status, content)
	}
}

func (api *API) handleJsonResponse(w http.ResponseWriter, status int, content any) {
	if status != http.StatusNoContent {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(status)
		if err := json.NewEncoder(w).Encode(content); err != nil {
			log.Printf("(handleJsonResponse) no se pudo escribir el contenido en formato json: %v", err)
			return
		}
	} else {
		w.WriteHeader(status)
	}
}
