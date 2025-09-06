package app

import (
	"net/http"
)

// Middleware sederhana untuk cek header Authorization
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		// contoh validasi token (misalnya harus "Bearer mysecrettoken")
		if token != "Bearer mysecrettoken" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		// lanjut ke handler berikutnya
		next.ServeHTTP(w, r)
	})
}
