// ENTRY POINT

package main

import (
	"log"
	"net/http"

	"runix/internal/api"
	"runix/internal/handlers"
	"runix/internal/utils"

	"github.com/gorilla/mux"
)

func main() {
	router := api.NewRouter()

	port := ":8080"

	log.Printf("Starting server on %s\n", port)

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Error starting server", err)
		utils.LogError(err)

	}
}

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/execute", handlers.ExecuteHandler).Methods("POST")
}
