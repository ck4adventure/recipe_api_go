package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"recipe_app_go/internal/app_server"
	"syscall"
	"time"
)

func main() {
	s := app_server.NewAppServer(":8080")

	fmt.Println("Starting server on :8080")

	// Run server in a goroutine so that it doesn't block
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	// Listen for interrupt signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down server...")

	// Give active connections time to finish
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown: %v\n", err)
	}

	fmt.Println("Server exited gracefully")
}

// func init() {
// 	// do init things
// }
