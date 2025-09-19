package handler

import "net/http"

// TODO: testear si es posible capturar el 404 del file server (tipo wrapper) y reesribirlo como
// TODO: un nuevo paquete con header de status correcto. Si se puede, el middleware es posible. Si hay
// TODO: overflow de headers, no

// GetRouter devuelve el multiplexer del Servidor
func (app *Application) GetRouter() *http.ServeMux {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./web/static/"))

	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", app.Home)

	return mux
}
