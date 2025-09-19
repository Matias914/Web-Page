package handler

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strings"
)

// Renderer guarda la información de las plantillas a parsear y proporciona
// un metodo para ejecutarlas.
type Renderer struct {
	templates map[string]*template.Template
}

// NewRenderer genera una instancia del administrador de plantillas y carga
// todas las plantillas una vez.
func NewRenderer(templatesPath string) (*Renderer, error) {
	r := &Renderer{
		templates: make(map[string]*template.Template),
	}
	if err := r.loadTemplates(templatesPath); err != nil {
		return nil, err
	}
	return r, nil
}

// loadTemplates parsea todas las plantillas del mapa miembro. Resulta en error
// si alguna no puede parsearse.
func (t *Renderer) loadTemplates(templatesPath string) error {
	files, err := os.ReadDir(templatesPath)
	if err != nil {
		return err
	}
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".html") {
			name := strings.TrimSuffix(f.Name(), ".html")
			tmpl, err := template.ParseFiles(templatesPath + f.Name())
			if err != nil {
				return err
			}
			t.templates[name] = tmpl
		} else {
			return fmt.Errorf("el template '%s' no es un archivo '.html'", f.Name())
		}
	}
	return nil
}

// executeTemplate recibe el escritor del paquete de respuesta, el nombre de la plantilla a
// ejecutar y los datos dinamicos que podrian necesitarse. Con esto, se busca la plantilla y
// se prepara el Content-Type de la respuesta. Retorna un código de OK si no hubo errores.
func (t *Renderer) renderToBuffer(name string, data interface{}) (*bytes.Buffer, error) {
	tmpl, existe := t.templates[name]
	if !existe {
		err := fmt.Errorf("el template '%s' no fue encontrado", name)
		return nil, err
	}
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, data); err != nil {
		return nil, err
	}
	return buf, nil
}
