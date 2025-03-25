package handlers

import (
	"encoding/json"
	"net/http"
	"recipe_app_go/internal/models"
)

func respondWithJSON(w http.ResponseWriter, status int, payload models.APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func Home(w http.ResponseWriter, r *http.Request) {
	response := models.APIResponse{
		Success: true,
		Data:    map[string]string{"message": "hi"},
	}
	respondWithJSON(w, http.StatusOK, response)
}

func Health(w http.ResponseWriter, r *http.Request) {
	response := models.APIResponse{
		Success: true,
		Data:    map[string]string{"status": "ok"},
	}
	respondWithJSON(w, http.StatusOK, response)
}
