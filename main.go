// ENTRY POINT

package main

import (
	"log"
	"net/http"

	"runix/internal/api"
)

func main() {
	router := api.NewRouter()

	port := ":8080"

	log.Printf("Starting server on %s\n", port)

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Error starting server", err)
	}
}
