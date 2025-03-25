package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// API Response will be the struct for all responses, ensuring uniformity
type APIResponse struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func respondWithJSON(w http.ResponseWriter, status int, payload APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusTeapot, APIResponse{Success: true, Data: "hi"})
}

func main() {
	routeHandler := mux.NewRouter()

	routeHandler.HandleFunc("/", greetHandler)

	s := &http.Server{
		Addr:    ":8080",
		Handler: routeHandler,
	}
	fmt.Println("starting server")
	s.ListenAndServe()

}

func init() {
	fmt.Println("Initializing application...")
}
