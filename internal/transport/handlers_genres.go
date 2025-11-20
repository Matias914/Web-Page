package transport

import (
	"log"
	"net/http"

	"github.com/Matias914/Web-Page/internal/service"
)

// handleGenresCreate gestiona la solicitud HTTP POST para crear un nuevo género.
func (app *Application) handleGenresCreate(w http.ResponseWriter, r *http.Request) {
	// Se parsea el formulario de la solicitud
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Se crea la estructura GenreData con los datos recibidos
	genreData := service.GenreData{
		Name: r.FormValue("name"),
	}

	// Se valida la estructura GenreData
	if err := app.ValidationService.Validate(genreData); err != nil {
		http.Error(w, "Bad Request: Invalid data", http.StatusBadRequest)
		return
	}

	// Se llama al servicio para agregar el nuevo género a la base de datos
	_, err := app.GenreService.AddGenre(r.Context(), genreData)
	if err != nil {
		log.Printf("error creating genre: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Se usa StatusSeeOther (303) para forzar una solicitud GET posterior, previniendo reenvíos del formulario.
	http.Redirect(w, r, "/management?entity=genres", http.StatusSeeOther)
}
