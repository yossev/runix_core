package api

import (
	"runix/internal/handlers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	SetupRoutes(r)
	return r
}

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/execute", handlers.ExecuteHandler).Methods("POST")

}
