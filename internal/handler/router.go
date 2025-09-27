package handler

import "net/http"

// GetRouter devuelve el multiplexer del Servidor
func (app *Application) GetRouter() *http.ServeMux {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./web/static/"))

	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", app.Home)

	return mux
}
