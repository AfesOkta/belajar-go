package app

import "net/http"

// Middleware = fungsi yang menerima dan mengembalikan http.Handler
type Middleware func(http.Handler) http.Handler

// ChainMiddlewares menyatukan beberapa middleware jadi satu
func ChainMiddlewares(h http.Handler, middlewares ...Middleware) http.Handler {
	// urutan dibalik supaya middleware pertama di parameter
	// jalan paling luar
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}
