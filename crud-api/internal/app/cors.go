package app

import "net/http"

// Middleware untuk CORS
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Izinkan semua origin (untuk dev)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Header yang diizinkan
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// Method yang diizinkan
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Jika preflight request (OPTIONS), balas langsung
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// lanjut ke handler berikutnya
		next.ServeHTTP(w, r)
	})
}
