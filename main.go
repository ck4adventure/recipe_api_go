package main

import (
	"fmt"
	"net/http"
	"recipe_app_go/internal/app_server"
)

func main() {

	s := app_server.NewAppServer(":8080")

	fmt.Println("starting server")
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("Server error:", err)
	}
}

func init() {
	fmt.Println("Initializing application...")
}
