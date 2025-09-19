package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWithRecovery(t *testing.T) {
	// Handler que siempre hace panic
	panicHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("test WithRecovery")
	})
	// Lo envolvemos con el middleware
	handler := WithRecovery(panicHandler)
	// Request falso
	req := httptest.NewRequest(http.MethodGet, "/panic", nil)
	rec := httptest.NewRecorder()
	// Ejecutamos el handler
	handler.ServeHTTP(rec, req)
	// Verificamos que respondió con 500
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("esperado status 500, obtuve %d", rec.Code)
	}
	// Verificamos que se devolvió el mensaje de error esperado
	expected := "Error interno del servidor\n"
	if rec.Body.String() != expected {
		t.Errorf("esperado '%s', obtuve '%s'", expected, rec.Body.String())
	}
}
