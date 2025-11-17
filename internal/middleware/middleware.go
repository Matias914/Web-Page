package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/Matias914/Web-Page/internal/transport"
)

// WithLogging es una función que retorna un Handler que mide el tiempo, el path y el
// metodo de una solicitud por medio de un log.
func WithLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

// WithRecovery es una función que retorna un handler que intenta recuperar el servidor
// después de un error fatal.
func WithRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("el servidor entró en pánico: %v", err)

				// Crear una instancia de Application para usar su manejador de errores
				app := &transport.Application{}
				app.HandleServerError(w, r)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
