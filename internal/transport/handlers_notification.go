package transport

import (
	"log"
	"net/http"

	"github.com/Matias914/Web-Page/internal/transport/views"
)

// handleClearNotification devuelve un banner de notificación vacío para ocultarlo de la vista
func (app *Application) handleClearNotification(w http.ResponseWriter, r *http.Request) {
	component := views.ClearNotification()
	if err := component.Render(r.Context(), w); err != nil {
		log.Printf("error rendering clear notification: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
