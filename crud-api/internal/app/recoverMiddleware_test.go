package app

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// Handler dummy yang sengaja panic
func panicHandler(w http.ResponseWriter, r *http.Request) {
	panic("test panic")
}

func TestRecoverMiddleware(t *testing.T) {
	// Matikan log sementara supaya output test bersih
	log.SetOutput(io.Discard)
	defer log.SetOutput(os.Stderr)

	// Handler yang sengaja panic
	panicHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("test panic")
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	// Bungkus handler dengan recoverMiddleware
	handler := RecoverMiddleware(panicHandler)
	handler.ServeHTTP(rr, req)

	// Pastikan response tidak crash dan status code sesuai
	if rr.Code != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, rr.Code)
	}
}
