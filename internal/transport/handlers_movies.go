package transport

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Matias914/Web-Page/internal/service"
)

// handleMoviesCreate gestiona la solicitud HTTP POST para crear una nueva película.
// Se encarga de parsear los datos del formulario, validar la entrada, y delegar
// la lógica de persistencia a la capa de servicio.
func (app *Application) handleMoviesCreate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("error parsing form: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	duration, _ := strconv.Atoi(r.FormValue("duration_minutes"))
	releaseDate, _ := time.Parse("2006-01-02", r.FormValue("released_at"))

	movieData := service.MovieData{
		Title:           r.FormValue("title"),
		Synopsis:        r.FormValue("synopsis"),
		ReleasedAt:      releaseDate,
		DurationMinutes: duration,
		PosterUrl:       r.FormValue("poster_url"),
	}

	if err := app.ValidationService.Validate(movieData); err != nil {
		log.Printf("validation error: %v", err)
		http.Error(w, "Bad Request: Invalid data", http.StatusBadRequest)
		return
	}

	_, err := app.MovieService.AddMovie(r.Context(), movieData)
	if err != nil {
		log.Printf("error creating movie: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Se usa StatusSeeOther (303) para forzar una solicitud GET posterior, previniendo reenvíos del formulario.
	http.Redirect(w, r, "/management?entity=movies", http.StatusSeeOther)
}
