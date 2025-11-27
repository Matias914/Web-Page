package transport

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/Matias914/Web-Page/internal/service"
	"github.com/Matias914/Web-Page/internal/transport/views"
	"github.com/go-chi/chi/v5"
)

// handleCelebritiesCreate gestiona la solicitud HTTP POST para crear una nueva celebridad.
func (app *Application) handleCelebritiesCreate(w http.ResponseWriter, r *http.Request) {
	// Se parsea el formulario de la solicitud
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// Se crea la estructura CelebrityData con los datos recibidos
	birthDate, _ := time.Parse("2006-01-02", r.FormValue("birth_date"))
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
		if errors.Is(err, service.ErrCelebrityDuplicated) {
			views.Notification("Error: La celebridad ya existe.", true).Render(r.Context(), w)
			config := service.EntityConfigs["celebrities"]
			items, _ := app.CelebrityService.GetCelebritiesList(r.Context(), 1, 100)
			views.ManagementFormAndList(config, items).Render(r.Context(), w)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Se re-renderiza la lista de celebridades
	config := service.EntityConfigs["celebrities"]
	items, _ := app.CelebrityService.GetCelebritiesList(r.Context(), 1, 100)
	views.Notification("Celebridad creada con exito", false).Render(r.Context(), w)
	views.ManagementFormAndList(config, items).Render(r.Context(), w)
}

// handleCelebrityDelete gestiona la solicitud HTTP DELETE para eliminar una celebridad.
func (app *Application) handleCelebrityDelete(w http.ResponseWriter, r *http.Request) {
	// Se obtiene el ID de la celebridad desde la URL
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	// Se llama al servicio para eliminar la celebridad de la base de datos
	err = app.CelebrityService.DeleteCelebrity(r.Context(), id)
	if err != nil {
		component := views.Notification("Error al eliminar la celebridad", true)
		component.Render(r.Context(), w)
		return
	}
	// Se envía una notificación de éxito
	component := views.Notification("Celebridad eliminada con exito", false)
	component.Render(r.Context(), w)
}

// handleCelebritiesUpdate gestiona la solicitud HTTP PUT para actualizar una celebridad existente.
func (app *Application) handleCelebritiesUpdate(w http.ResponseWriter, r *http.Request) {
	// Se parsea el formulario de la solicitud
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// Se obtiene el ID de la celebridad desde la URL
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	// Se crea la estructura CelebrityData con los datos recibidos
	birthDate, _ := time.Parse("2006-01-02", r.FormValue("birth_date"))
	celebrityData := service.CelebrityData{
		Name:      r.FormValue("name"),
		BirthDate: birthDate,
	}
	// Se valida la estructura CelebrityData
	if err := app.ValidationService.Validate(celebrityData); err != nil {
		http.Error(w, "Bad Request: Invalid data", http.StatusBadRequest)
		return
	}

	// Se llama al servicio para actualizar la celebridad en la base de datos
	_, err = app.CelebrityService.UpdateCelebrity(r.Context(), id, celebrityData)
	if err != nil {
		if errors.Is(err, service.ErrCelebrityDuplicated) {
			views.Notification("Error: Ya existe una celebridad con ese nombre y fecha de nacimiento.", true).Render(r.Context(), w)

			// Se re-renderiza la lista de celebridades
			config := service.EntityConfigs["celebrities"]
			items, _ := app.CelebrityService.GetCelebritiesList(r.Context(), 1, 100)
			views.ManagementFormAndList(config, items).Render(r.Context(), w)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Se re-renderiza la lista de celebridades
	config := service.EntityConfigs["celebrities"]
	items, _ := app.CelebrityService.GetCelebritiesList(r.Context(), 1, 100)
	views.Notification("Celebridad actualizada con exito", false).Render(r.Context(), w)
	views.ManagementFormAndList(config, items).Render(r.Context(), w)
}

// handleCelebrityEditForm gestiona la solicitud HTTP GET para mostrar el formulario de edición de una celebridad.
func (app *Application) handleCelebrityEditForm(w http.ResponseWriter, r *http.Request) {
	// Se obtiene el ID de la celebridad desde la URL
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	// Se obtiene la celebridad desde el servicio
	celebrity, err := app.CelebrityService.GetCelebrity(r.Context(), id)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	// Se renderiza el formulario de edición con los datos de la celebridad
	config := service.EntityConfigs["celebrities"]
	component := views.CelebrityEditForm(config, celebrity)
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
