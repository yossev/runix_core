package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"runix/internal/executor"
	"runix/internal/utils"
)

// HANDLER HANDLES THE REQURESTS AND SENDS THEM TO THE EXECUTOR

type ExecuteRequest struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}

func ExecuteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received Request") //Debugging
	var request ExecuteRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		utils.LogError(err)
		return
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
