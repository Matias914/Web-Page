package transport

import (
	"log"
	"net/http"

	"github.com/Matias914/Web-Page/internal/transport/views"
	"github.com/a-h/templ"

	"github.com/Matias914/Web-Page/internal/service"
)

type Application struct {
	MovieService     *service.MovieService
	GenreService     *service.GenreService
	CelebrityService *service.CelebrityService

	ValidationService *service.ValidationService
}

func (app *Application) HandleServerError(w http.ResponseWriter, r *http.Request) {
	component := views.Error500Page()
	app.handleTemplRender(w, r, http.StatusInternalServerError, component)
}

func (app *Application) handleNotFound(w http.ResponseWriter, r *http.Request) {
	component := views.Error404Page()
	app.handleTemplRender(w, r, http.StatusNotFound, component)
}

// handleTemplRender es un método de conveniencia para renderizar componentes Templ.
// Se encarga de establecer las cabeceras HTTP necesarias y el código de estado.
func (app *Application) handleTemplRender(w http.ResponseWriter, r *http.Request, status int, component templ.Component) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)

	if err := component.Render(r.Context(), w); err != nil {
		log.Printf("(handleTemplRender) no se pudo escribir el componente en la respuesta: %v", err)
	}
}
