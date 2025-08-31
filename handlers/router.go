package handlers

import "net/http"

// GetRouter devuelve el multiplexer del Servidor
func GetRouter() *http.ServeMux {
	mux := http.NewServeMux()
	// TODO: crear un file server personalizado para mostrar paginas de error
	// Crea un servidor de archivos estáticos
	fs := http.FileServer(http.Dir("./static/"))
	// Agrega al routing al FileServer
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	// Agrega al routing los handlers
	mux.HandleFunc("/", homeHandler)
	// Retorna el handler correspondiente
	return mux
}
