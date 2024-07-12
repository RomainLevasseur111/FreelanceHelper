package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/api/register", register)
    c := cors.AllowAll()

	// Wrap your handler with CORS
	handler := c.Handler(mux)

    log.Println("API server is running on port 8040")

    log.Fatal(http.ListenAndServe(":8040", handler))
}
