package app

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{Id: 1, Name: "Afes"},
	{Id: 2, Name: "Retno"},
	{Id: 3, Name: "Naya"},
	{Id: 4, Name: "Aditya"},
}

// GET /users -> ambil semua user
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GET /users/{id} -> ambil user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	for _, u := range users {
		if u.Id == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(u)
			return
		}
	}
	http.Error(w, "user not found", http.StatusNotFound)
}

// POST /users -> tambah user baru
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	newUser.Id = len(users) + 1
	users = append(users, newUser)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

// PUT /users/{id} -> update user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	var updateUser User
	if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	for i, u := range users {
		if u.Id == id {
			users[i].Name = updateUser.Name
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}
	http.Error(w, "user not found", http.StatusNotFound)
}

// DELETE /users/{id} -> hapus user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	for i, u := range users {
		if u.Id == id {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "user not found", http.StatusNotFound)
}

// TestPanic -> endpoint untuk simulasi panic
func TestPanic(w http.ResponseWriter, r *http.Request) {
	// Contoh panic: akses slice out of range
	users := []string{"Afes", "Retno"}
	_ = users[5] // ❌ ini akan panic

	w.Write([]byte("seharusnya tidak sampai sini"))
}
