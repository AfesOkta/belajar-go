package main

import (
	"crud-api/internal/app"
	"fmt"
	"log"
	"net/http"
)

func main() {
	apiMux := http.NewServeMux()

	// Routes CRUD
	apiMux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			app.GetUsers(w, r)
		} else if r.Method == http.MethodPost {
			app.CreateUser(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	apiMux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			app.GetUser(w, r)
		case http.MethodPut:
			app.UpdateUser(w, r)
		case http.MethodDelete:
			app.DeleteUser(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
	apiMux.HandleFunc("/test-panic", app.TestPanic)

	// Router utama
	mux := http.NewServeMux()

	// Semua request dengan prefix /api diarahkan ke apiMux
	mux.Handle("/api/", http.StripPrefix("/api", apiMux))

	// Tambah middleware logging
	// Tambah middleware recover + logging
	// Gabungkan middleware: Recover dulu, lalu Logging, lalu CORS, lalu Auth
	finalHandler := app.ChainMiddlewares(
		mux,
		app.RecoverMiddleware,
		app.LoggingMiddleware,
		app.CORSMiddleware,
		app.AuthMiddleware,
	)

	loggedMux := app.LoggingMiddleware(finalHandler)
	safeMux := app.RecoverMiddleware(loggedMux)

	fmt.Println("🚀 Server running at http://localhost:8089")
	log.Fatal(http.ListenAndServe(":8089", safeMux))
}
