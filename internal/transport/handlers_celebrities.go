package transport

import (
	"log"
	"net/http"
	"time"

	"github.com/Matias914/Web-Page/internal/service"
)

func (app *Application) handleCelebritiesCreate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	birthDate, _ := time.Parse("2006-01-02", r.FormValue("birth_date"))

	celebrityData := service.CelebrityData{
		Name:      r.FormValue("name"),
		BirthDate: birthDate,
	}
	if err := app.ValidationService.Validate(celebrityData); err != nil {
		http.Error(w, "Bad Request: Invalid data", http.StatusBadRequest)
		return
	}

	_, err := app.CelebrityService.AddCelebrity(r.Context(), celebrityData)
	if err != nil {
		log.Printf("error creating celebrity: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/management?entity=celebrities", http.StatusSeeOther)
}
