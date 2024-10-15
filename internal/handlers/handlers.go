package handlers

import (
	"encoding/json"
	"net/http"

	"runix/internal/api"
	"runix/internal/executor"
	"runix/internal/utils"
)

// HANDLER HANDLES THE REQURESTS AND SENDS THEM TO THE EXECUTOR

type ExecuteRequest struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}

var limiter *api.RateLimiter

func ExecuteHandler(w http.ResponseWriter, r *http.Request) {
	var request ExecuteRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		utils.LogError(err)
		return
	}

	apiKey := r.Header.Get("X-API-KEY")
	if !limiter.Allow(apiKey) {
		http.Error(w, "rate limit exceeded.", http.StatusTooManyRequests)
	}

	result, err := executor.ExecuteCode(request.Code, request.Language)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		utils.LogError(err)
		return
	}
	utils.LogInfo(result)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
