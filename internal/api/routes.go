package api

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	// CRUD Operations that the router will handle
	router.HandleFUNc("/execute", ExecuteHandler).Methods("POST")
	return router
}
