package main

import (
	"fmt"
	"hello-api/internal/app"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/hello", app.HelloHandler)
	mux.HandleFunc("/ping", app.PingHandler)
	mux.HandleFunc("/greet/", app.GreetHandler) // pakai prefix "/greet/"

	fmt.Println("🚀 Server running at http://localhost:8089")

	log.Fatal(http.ListenAndServe(":8089", mux))
}
