package transport

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Matias914/Web-Page/internal/service"

	views "github.com/Matias914/Web-Page/internal/transport/views"
)

// handleIndexPage gestiona la solicitud HTTP para la página de inicio ("/").
// Renderiza el componente IndexPage utilizando la plantilla base.
func (app *Application) handleIndexPage(w http.ResponseWriter, r *http.Request) {
	component := views.IndexPage()

	if err := component.Render(r.Context(), w); err != nil {
		log.Printf("error rendering index page: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// handleCatalogPage gestiona la solicitud HTTP para la página del catálogo de películas ("/catalog").
// Llama al servicio para obtener la lista de películas y renderiza el componente CatalogPage.
func (app *Application) handleCatalogPage(w http.ResponseWriter, r *http.Request) {
	movies, err := app.MovieService.GetMoviesList(r.Context(), 1, 100) // Fetching all for now
	if err != nil {
		log.Printf("error getting movies list for catalog: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	component := views.CatalogPage(movies)
	err = component.Render(r.Context(), w)
	if err != nil {
		log.Printf("error rendering catalog page: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// handleControlPage gestiona la solicitud HTTP para el panel de administración ("/management").
// Determina la entidad (movies, genres, celebrities) a gestionar y renderiza la página completa.
func (app *Application) handleControlPage(w http.ResponseWriter, r *http.Request) {
	entityName := r.URL.Query().Get("entity")
	if entityName == "" {
		entityName = "movies"
	}

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}
	rows, _ := strconv.Atoi(r.URL.Query().Get("rows"))
	if rows < 1 {
		rows = 10
	}

	config, exists := service.EntityConfigs[entityName]
	if !exists {
		http.Error(w, "Entity not found", http.StatusNotFound)
		return
	}

	// Se llama al servicio correspondiente basado en 'entityName'.
	var items any
	var err error
	switch entityName {
	case "movies":
		items, err = app.MovieService.GetMoviesList(r.Context(), page, rows)
	case "genres":
		items, err = app.GenreService.GetGenresList(r.Context(), page, rows)
	case "celebrities":
		items, err = app.CelebrityService.GetCelebritiesList(r.Context(), page, rows)
	}

	if err != nil {
		log.Printf("error getting %s list: %v", entityName, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	component := views.ControlPage(config, items)
	if err := component.Render(r.Context(), w); err != nil {
		log.Printf("error rendering control page: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
