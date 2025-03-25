package main

import (
	"fmt"
	"net/http"
	"recipe_app_go/internal/handlers"
	"recipe_app_go/internal/middleware"

	"github.com/gorilla/mux"
)

func main() {
	routeHandler := mux.NewRouter()

	routeHandler.HandleFunc("/", handlers.Home)
	routeHandler.HandleFunc("/health", handlers.Health)

	routeHandler.Use(middleware.LoggingMiddleware)

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
