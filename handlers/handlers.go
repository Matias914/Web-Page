package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// ---------------------------
//	     Error Handlers
// ---------------------------

// errorHandler se encarga de manejar los errores que puedan surgir y mostrando la
// página correspondiente. Necesita del escritor de respuesta y el código de error.
func errorHandler(w http.ResponseWriter, status int) {
	content, err := os.ReadFile(fmt.Sprintf("./errors/%d.html", status))
	if err != nil {
		log.Printf("No se pudo cargar el archivo '%d.html'", status)
		http.Error(w, "Error Interno del Servidor", http.StatusInternalServerError)
		return
	}
	// Setea los campos correspondientes del header. Al hacer WriteHeader la operación
	// se completa (y no puede sobreescribirse)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	// Escribe el contenido del archivo html en la respuesta y se completa el paquete
	if _, err = w.Write(content); err != nil {
		log.Printf("No se pudo hacer llegar la respuesta al cliente: %v", err)
		return
	}
}

// ---------------------------
//	      URL Handlers
// ---------------------------

// homeHandler o más bien Homelander, es una funcion que maneja las solicitudes
// que se dirigen al root o a URL desconocidas.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	t := newTemplateAdmin()
	if r.URL.Path != "/" {
		errorHandler(w, http.StatusNotFound)
		return
	}
	if ok := t.executeTemplate(w, "index", nil); !ok {
		errorHandler(w, http.StatusInternalServerError)
	}
}
