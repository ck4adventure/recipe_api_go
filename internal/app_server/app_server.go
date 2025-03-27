package app_server

import (
	"net/http"
	"recipe_app_go/internal/handlers"
	"recipe_app_go/internal/middleware"

	"github.com/gorilla/mux"
)

func MuxRouter() *mux.Router {
	routeHandler := mux.NewRouter()

	routeHandler.HandleFunc("/", handlers.Home)
	routeHandler.HandleFunc("/health", handlers.Health)
	routeHandler.Use(middleware.LoggingMiddleware)

	return routeHandler
}

func NewAppServer(addr string) *http.Server {
	// TODO check addr to be proper before using
	return &http.Server{
		Addr:    addr,
		Handler: MuxRouter(),
	}
}
