package app

import (
	"log"
	"net/http"
	"runtime/debug"
)

// Middleware untuk menangkap panic agar server tidak crash
func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("⛔ panic recovered: %v\n%s", err, debug.Stack())
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		// jalankan handler berikutnya
		next.ServeHTTP(w, r)
	})
}
