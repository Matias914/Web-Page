package handlers

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

var createOnce sync.Once
var loadOnce sync.Once
var instance *TemplateAdmin

// TemplateAdmin guarda la información de las plantillas a parsear y proporciona
// un metodo para ejecutarlas. Tiene un diseño Singleton.
type TemplateAdmin struct {
	templates map[string]*template.Template
}

// newTemplateAdmin genera una instancia del administrador de plantillas si no
// existe. De lo contrario, retorna la existente.
func newTemplateAdmin() *TemplateAdmin {
	createOnce.Do(func() {
		instance = &TemplateAdmin{
			templates: make(map[string]*template.Template),
		}
	})
	return instance
}

// loadTemplates parsea todas las plantillas del mapa miembro. Resulta en error
// si alguna no puede parsearse. Debe ser invocada con sync.
func (t *TemplateAdmin) loadTemplates() error {
	// Se pide al OS la lista de entradas de directorio
	files, err := os.ReadDir("./templates")
	if err != nil {
		return err
	}
	// Por cada entrada de directorio de la lista
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".html") {
			name := strings.TrimSuffix(f.Name(), ".html")
			// Se intenta parsear el archivo plantilla y agregar al mapa de templates
			tmpl, err := template.ParseFiles("./templates/" + f.Name())
			if err != nil {
				return err
			}
			t.templates[name] = tmpl
		} else {
			return errors.New("El template " + f.Name() + " no es un archivo HTML!")
		}
	}
	return nil
}

// executeTemplate recibe el escritor del paquete de respuesta, el nombre de la plantilla a
// ejecutar y los datos dinamicos que podrian necesitarse. Con esto, se carga la plantilla y
// se prepara el Content-Type de la respuesta. Retorna un código de OK si no hubo errores.
func (t *TemplateAdmin) executeTemplate(w http.ResponseWriter, name string, data interface{}) bool {
	// El primer hilo que llega, carga todas
	loadOnce.Do(func() {
		if err := t.loadTemplates(); err != nil {
			log.Printf("Algunos de los templates no se pudieron cargar: %v", err)
		}
	})
	// Se busca la plantilla requerida dentro del mapa
	tmpl, existe := t.templates[name]
	if !existe {
		log.Printf("El template '%s' no fue encontrado", name)
		return false
	}
	// Se ejecuta la plantilla
	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("No se pudo ejecutar el template '%s': %v", name, err)
		return false
	}
	// Se define el formato del contenido del paquete de respuesta
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return true
}
