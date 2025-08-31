package middleware

import (
	"log"
	"net/http"
	"time"
)

// WithLogging es una función que retorna un Handler que mide el tiempo, el path y el
// metodo de una solicitud por medio de un log.
func WithLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// Se invoca al siguiente handler
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

// WithRecovery es una función que retorna un handler que intenta recuperar el servidor
// después de un error fatal.
func WithRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Al tener defer se ejecuta luego de la función (haya habido error o no)
		// Si no hubo error, recover() retorna nil. Si hubo error, depende de si
		// puede recuperarse o no. Si no lo hace, se loguea.
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recuperado de pánico: %v", err)
				http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
