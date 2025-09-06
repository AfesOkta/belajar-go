package app_test

import (
	"crud-api/internal/app"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Struct harus sama dengan yang di app.User
type Users struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func TestGetUsers(t *testing.T) {
	// Buat request dummy
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()

	// Jalankan handler langsung (public function dari app)
	app.GetUsers(rec, req)

	// Validasi status code
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	// Decode response JSON
	var got []User
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	// Pastikan data tidak kosong
	if len(got) == 0 {
		t.Errorf("expected users, got empty list")
	}

	// Contoh validasi user pertama
	if got[0].Name != "Afes" {
		t.Errorf("expected first user Afes, got %s", got[0].Name)
	}
}
