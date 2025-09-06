package app

import (
	"log"
	"net/http"
	"time"
)

// Middleware untuk logging request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("➡️  %s %s", r.Method, r.URL.Path)

		// jalankan handler berikutnya
		next.ServeHTTP(w, r)

		log.Printf("✅ selesai %s %s (%s)", r.Method, r.URL.Path, time.Since(start))
	})
}
