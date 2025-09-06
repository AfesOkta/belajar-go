package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // ambil query string ?name=Afes
	if name == "" {
		name = "World"
	}
	message := fmt.Sprintf("Hello, %s! Welcome to Golang REST API 🚀", name)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(message))
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "ok", "message": "pong"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/greet/")
	if path == "" {
		http.Error(w, "name not provided", http.StatusBadRequest)
		return
	}
	message := fmt.Sprintf("Nice to meet you, %s 👋", path)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(message))
}
