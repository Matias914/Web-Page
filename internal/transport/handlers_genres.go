package transport

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Matias914/Web-Page/internal/service"
	"github.com/Matias914/Web-Page/internal/transport/views"
	"github.com/go-chi/chi/v5"
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
		if errors.Is(err, service.ErrGenreDuplicated) {
			views.Notification("Error: El genero ya existe.", true).Render(r.Context(), w)

			// Se re-renderiza la lista de géneros
			config := service.EntityConfigs["genres"]
			items, _ := app.GenreService.GetGenresList(r.Context(), 1, 100)
			views.ManagementFormAndList(config, items).Render(r.Context(), w)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Se re-renderiza la lista de géneros y se envía una notificación de éxito
	config := service.EntityConfigs["genres"]
	items, _ := app.GenreService.GetGenresList(r.Context(), 1, 100)
	views.Notification("Genero creado con exito", false).Render(r.Context(), w)
	views.ManagementFormAndList(config, items).Render(r.Context(), w)
}

// handleGenreDelete gestiona la solicitud HTTP DELETE para eliminar un género existente.
func (app *Application) handleGenreDelete(w http.ResponseWriter, r *http.Request) {
	// Se obtiene el ID del género desde la URL
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Se llama al servicio para eliminar el género de la base de datos
	err = app.GenreService.DeleteGenre(r.Context(), id)
	if err != nil {
		component := views.Notification("Error al eliminar el género", true)
		component.Render(r.Context(), w)
		return
	}

	// Se envía una notificación de éxito
	component := views.Notification("Genero eliminado con exito", false)
	component.Render(r.Context(), w)
}

// handleGenresUpdate gestiona la solicitud HTTP PUT para actualizar un género existente.
func (app *Application) handleGenresUpdate(w http.ResponseWriter, r *http.Request) {
	// Se parsea el formulario de la solicitud
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Se obtiene el ID del género desde la URL
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
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

	// Se llama al servicio para actualizar el género en la base de datos
	_, err = app.GenreService.UpdateGenre(r.Context(), id, genreData)
	if err != nil {
		if errors.Is(err, service.ErrGenreDuplicated) {
			views.Notification("Error: Ya existe un genero con ese nombre.", true).Render(r.Context(), w)
			config := service.EntityConfigs["genres"]
			items, _ := app.GenreService.GetGenresList(r.Context(), 1, 100)
			views.ManagementFormAndList(config, items).Render(r.Context(), w)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Se re-renderiza la lista de géneros y se envía una notificación de éxito
	config := service.EntityConfigs["genres"]
	items, _ := app.GenreService.GetGenresList(r.Context(), 1, 100)
	views.Notification("Genero actualizado con exito", false).Render(r.Context(), w)
	views.ManagementFormAndList(config, items).Render(r.Context(), w)
}

// handleGenreEditForm gestiona la solicitud HTTP GET para mostrar el formulario de edición de un género.
func (app *Application) handleGenreEditForm(w http.ResponseWriter, r *http.Request) {
	// Se obtiene el ID del género desde la URL
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Se obtiene el género desde el servicio
	genre, err := app.GenreService.GetGenre(r.Context(), id)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// Se crea el componente del formulario de edición y se renderiza
	config := service.EntityConfigs["genres"]
	component := views.GenreEditForm(config, genre)
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
