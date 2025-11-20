package transport

import (
	"log"
	"net/http"
	"time"

	"github.com/Matias914/Web-Page/internal/service"
)

func (app *Application) handleCelebritiesCreate(w http.ResponseWriter, r *http.Request) {
	// Se parsea el formulario de la solicitud
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Se extraen y convierten los datos del formulario
	birthDate, _ := time.Parse("2006-01-02", r.FormValue("birth_date"))

	// Se crea la estructura CelebrityData con los datos recibidos
	celebrityData := service.CelebrityData{
		Name:      r.FormValue("name"),
		BirthDate: birthDate,
	}

	// Se valida la estructura CelebrityData
	if err := app.ValidationService.Validate(celebrityData); err != nil {
		http.Error(w, "Bad Request: Invalid data", http.StatusBadRequest)
		return
	}

	// Se llama al servicio para agregar la nueva celebridad a la base de datos
	_, err := app.CelebrityService.AddCelebrity(r.Context(), celebrityData)
	if err != nil {
		log.Printf("error creating celebrity: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Se usa StatusSeeOther (303) para forzar una solicitud GET posterior, previniendo reenvíos del formulario.
	http.Redirect(w, r, "/management?entity=celebrities", http.StatusSeeOther)
}
