package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Matias914/Web-Page/internal/service"
)

func (api *API) handleErrorResponse(w http.ResponseWriter, error error) {
	switch {
	case
		api.ValidationService.IsValidationError(error),
		errors.Is(error, service.ErrInvalidCategory),
		errors.Is(error, service.ErrInvalidCelebrity),
		errors.Is(error, service.ErrInvalidGenre),
		errors.Is(error, service.ErrInvalidMovie),
		errors.Is(error, service.ErrInvalidRating),
		errors.Is(error, service.ErrInvalidReview),
		errors.Is(error, service.ErrInvalidRole),
		errors.Is(error, service.ErrInvalidUser),
		errors.Is(error, service.ErrCelebrityReferenced),
		errors.Is(error, service.ErrGenreReferenced),
		errors.Is(error, service.ErrMovieReferenced),
		errors.Is(error, service.ErrUserReferenced):
		content := JsonError{Mssg: error.Error()}
		api.handleJsonResponse(w, http.StatusBadRequest, content)
	case
		errors.Is(error, service.ErrCategoryNotFound),
		errors.Is(error, service.ErrCelebrityNotFound),
		errors.Is(error, service.ErrGenreNotFound),
		errors.Is(error, service.ErrMovieNotFound),
		errors.Is(error, service.ErrRatingNotFound),
		errors.Is(error, service.ErrReviewNotFound),
		errors.Is(error, service.ErrRoleNotFound),
		errors.Is(error, service.ErrUserNotFound):
		content := JsonError{Mssg: error.Error()}
		api.handleJsonResponse(w, http.StatusNotFound, content)
	case
		errors.Is(error, service.ErrCategoryDuplicated),
		errors.Is(error, service.ErrCelebrityDuplicated),
		errors.Is(error, service.ErrGenreDuplicated),
		errors.Is(error, service.ErrMovieDuplicated),
		errors.Is(error, service.ErrRatingDuplicated),
		errors.Is(error, service.ErrReviewDuplicated),
		errors.Is(error, service.ErrRoleDuplicated),
		errors.Is(error, service.ErrUserDuplicated):
		content := JsonError{Mssg: error.Error()}
		api.handleJsonResponse(w, http.StatusConflict, content)
	default:
		content := JsonError{Mssg: "internal server error"}
		log.Println(error.Error())
		api.handleJsonResponse(w, http.StatusInternalServerError, content)
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
