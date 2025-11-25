package transport

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Matias914/Web-Page/internal/service"
	"github.com/Matias914/Web-Page/internal/transport/views"
	"github.com/go-chi/chi/v5"
)

// handleMoviesCreate gestiona la solicitud HTTP POST para crear una nueva película.
// Se encarga de parsear los datos del formulario, validar la entrada, y delegar
// la lógica de persistencia a la capa de servicio.
func (app *Application) handleMoviesCreate(w http.ResponseWriter, r *http.Request) {
	// Se parsea el formulario de la solicitud
	if err := r.ParseForm(); err != nil {
		log.Printf("error parsing form: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Se extraen y convierten los datos del formulario
	duration, _ := strconv.Atoi(r.FormValue("duration_minutes"))
	releaseDate, _ := time.Parse("2006-01-02", r.FormValue("released_at"))

	// Se crea la estructura MovieData con los datos recibidos
	movieData := service.MovieData{
		Title:           r.FormValue("title"),
		Synopsis:        r.FormValue("synopsis"),
		ReleasedAt:      releaseDate,
		DurationMinutes: duration,
		PosterUrl:       r.FormValue("poster_url"),
	}

	// Se valida la estructura MovieData
	if err := app.ValidationService.Validate(movieData); err != nil {
		log.Printf("validation error: %v", err)
		http.Error(w, "Bad Request: Invalid data", http.StatusBadRequest)
		return
	}

	// Se llama al servicio para agregar la nueva película a la base de datos
	_, err := app.MovieService.AddMovie(r.Context(), movieData)
	if err != nil {
		if errors.Is(err, service.ErrMovieDuplicated) {
			views.Notification("Error: La pelicula ya existe.", true).Render(r.Context(), w)

			// Se re-renderiza la lista de películas
			config := service.EntityConfigs["movies"]
			items, _ := app.MovieService.GetMoviesList(r.Context(), 1, 100)
			views.ManagementFormAndList(config, items).Render(r.Context(), w)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Se re-renderiza la lista de películas y se envía una notificación de éxito
	config := service.EntityConfigs["movies"]
	items, _ := app.MovieService.GetMoviesList(r.Context(), 1, 100)

	views.Notification("Pelicula creada con exito", false).Render(r.Context(), w)
	views.ManagementFormAndList(config, items).Render(r.Context(), w)
}

// handleMovieDelete gestiona la solicitud HTTP DELETE para eliminar una película existente
func (app *Application) handleMovieDelete(w http.ResponseWriter, r *http.Request) {
	// Se obtiene el ID de la película desde los parámetros de la URL
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("error parsing id: %v", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Se llama al servicio para eliminar la película de la base de datos
	err = app.MovieService.DeleteMovie(r.Context(), id)
	if err != nil {
		// En caso de error, también se puede notificar al usuario
		log.Printf("error deleting movie: %v", err)
		component := views.Notification("Error al eliminar la película", true)
		component.Render(r.Context(), w)
		return
	}

	// Se renderiza el componente de notificación para mostrar el éxito
	component := views.Notification("Pelicula eliminada con exito", false)
	component.Render(r.Context(), w)
}

// handleMoviesUpdate gestiona la solicitud HTTP PUT para actualizar una película existente
func (app *Application) handleMoviesUpdate(w http.ResponseWriter, r *http.Request) {
	// Se parsea el formulario de la solicitud
	if err := r.ParseForm(); err != nil {
		log.Printf("error parsing form: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Se obtiene el ID de la película desde los parámetros de la URL
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("invalid movie id: %v", err)
		http.Error(w, "Invalid Movie ID", http.StatusBadRequest)
		return
	}

	// Se construye la estructura MovieData con los datos del formulario
	duration, _ := strconv.Atoi(r.FormValue("duration_minutes"))
	releaseDate, _ := time.Parse("2006-01-02", r.FormValue("released_at"))

	// Se crea la estructura MovieData con los datos recibidos
	movieData := service.MovieData{
		Title:           r.FormValue("title"),
		Synopsis:        r.FormValue("synopsis"),
		ReleasedAt:      releaseDate,
		DurationMinutes: duration,
		PosterUrl:       r.FormValue("poster_url"),
	}

	// Se valida la estructura MovieData
	if err := app.ValidationService.Validate(movieData); err != nil {
		log.Printf("validation error: %v", err)
		http.Error(w, "Bad Request: invalid movie data", http.StatusBadRequest)
		return
	}

	// Se llama al servicio para actualizar la película en la base de datos
	_, err = app.MovieService.UpdateMovie(r.Context(), id, movieData)
	if err != nil {
		if errors.Is(err, service.ErrMovieDuplicated) {
			// Se notifica al usuario sobre el error de duplicación usando OOB swap.
			views.Notification("Error: Ya existe una pelicula con ese titulo y fecha de lanzamiento.", true).Render(r.Context(), w)
			// Se recarga el formulario y la lista para que no desaparezcan.
			config := service.EntityConfigs["movies"]
			items, _ := app.MovieService.GetMoviesList(r.Context(), 1, 100)
			views.ManagementFormAndList(config, items).Render(r.Context(), w)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Se re-renderiza la lista de películas y se envía una notificación de éxito
	config := service.EntityConfigs["movies"]
	items, _ := app.MovieService.GetMoviesList(r.Context(), 1, 100)
	views.Notification("Pelicula actualizada con exito", false).Render(r.Context(), w)
	views.ManagementFormAndList(config, items).Render(r.Context(), w)
}

// handleMovieEditForm gestiona la solicitud HTTP GET para mostrar el formulario de edición de una película.
func (app *Application) handleMovieEditForm(w http.ResponseWriter, r *http.Request) {
	// Se obtiene el ID de la película desde los parámetros de la URL
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	movie, err := app.MovieService.GetMovie(r.Context(), id)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// Se obtiene la configuración de la entidad película
	config := service.EntityConfigs["movies"] // o como lo obtengas
	component := views.MovieEditForm(config, movie)

	// Se genera un render del componente del formulario de edición
	if err := component.Render(r.Context(), w); err != nil {
		log.Printf("error rendering edit form: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
