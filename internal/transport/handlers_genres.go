package transport

import (
	"log"
	"net/http"

	"github.com/Matias914/Web-Page/internal/service"
)

// handleGenresCreate gestiona la solicitud HTTP POST para crear un nuevo género.
func (app *Application) handleGenresCreate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	genreData := service.GenreData{
		Name: r.FormValue("name"),
	}
	if err := app.ValidationService.Validate(genreData); err != nil {
		http.Error(w, "Bad Request: Invalid data", http.StatusBadRequest)
		return
	}

	_, err := app.GenreService.AddGenre(r.Context(), genreData)
	if err != nil {
		log.Printf("error creating genre: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/management?entity=genres", http.StatusSeeOther)
}
