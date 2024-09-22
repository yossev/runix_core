package api

import (
	"encoding/json"
	"net/http"
)

// HANDLER HANDLES THE REQURESTS AND SENDS THEM TO THE EXECUTOR

type ExecuteRequest struct {
	Code     string `json:"code"`
	Language string `json:"langauge`
}

func ExecuteHandler(w http.ResponseWriter, r *http.Request) {
	var request ExecuteRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	result, err := executor.ExecuteCode(request.Code, request.Language)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}
