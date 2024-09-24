package api

import (
	"encoding/json"
	"net/http"

	"github.com/yossev/runix_core/internal/executor"
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
		//LogError(err)
		return
	}

	result, err := executor.ExecuteCode(request.Code, request.Language)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		//LogError(err)
		return
	}
	//LogInfo(result)
	json.NewEncoder(w).Encode(result)
}
