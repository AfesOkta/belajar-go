package app_test

import (
	"bytes"
	"crud-api/internal/app"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

// Struct sama dengan app.User
type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// helper buat decode response JSON
func decodeBody(t *testing.T, rec *httptest.ResponseRecorder, v any) {
	t.Helper()
	if err := json.Unmarshal(rec.Body.Bytes(), v); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
}

func TestCRUDUser(t *testing.T) {
	// 1️⃣ Create User
	newUser := User{Name: "Testing User"}
	body, _ := json.Marshal(newUser)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	app.CreateUser(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("CreateUser: expected 200, got %d", rec.Code)
	}

	var created User
	decodeBody(t, rec, &created)

	if created.Id == 0 {
		t.Errorf("expected valid ID, got %d", created.Id)
	}
	if created.Name != newUser.Name {
		t.Errorf("expected name %s, got %s", newUser.Name, created.Name)
	}

	createdID := created.Id

	// 2️⃣ Get User by ID
	req = httptest.NewRequest(http.MethodGet, "/users/"+strconv.Itoa(createdID), nil)
	rec = httptest.NewRecorder()
	app.GetUser(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("GetUser: expected 200, got %d", rec.Code)
	}

	var fetched User
	decodeBody(t, rec, &fetched)

	if fetched.Id != createdID {
		t.Errorf("expected ID %d, got %d", createdID, fetched.Id)
	}

	// 3️⃣ Update User
	updatedUser := User{Name: "Updated User"}
	body, _ = json.Marshal(updatedUser)

	req = httptest.NewRequest(http.MethodPut, "/users/"+strconv.Itoa(createdID), bytes.NewReader(body))
	rec = httptest.NewRecorder()
	app.UpdateUser(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("UpdateUser: expected 200, got %d", rec.Code)
	}

	var updated User
	decodeBody(t, rec, &updated)

	if updated.Name != "Updated User" {
		t.Errorf("expected name Updated User, got %s", updated.Name)
	}

	// 4️⃣ Delete User
	req = httptest.NewRequest(http.MethodDelete, "/users/"+strconv.Itoa(createdID), nil)
	rec = httptest.NewRecorder()
	app.DeleteUser(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("DeleteUser: expected 204, got %d", rec.Code)
	}

	// 5️⃣ Get User lagi → harus 404
	req = httptest.NewRequest(http.MethodGet, "/users/"+strconv.Itoa(createdID), nil)
	rec = httptest.NewRecorder()
	app.GetUser(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("GetUser after delete: expected 404, got %d", rec.Code)
	}
}
