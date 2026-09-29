package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nikhilkotian6699/go-auth-demo/handlers"
	"github.com/nikhilkotian6699/go-auth-demo/repository"
	"github.com/nikhilkotian6699/go-auth-demo/routes"
	"github.com/nikhilkotian6699/go-auth-demo/services"
)

func main() {
	// Initialize in-memory user repository
	userRepo := repository.NewInMemoryUserRepository()

	// Initialize authentication service
	authService := services.NewAuthService(userRepo)

	// Initialize HTTP handlers
	authHandler := handlers.NewAuthHandler(authService)

	// Set up router and auth routes
	mux := http.NewServeMux()
	routes.SetupRoutes(mux, authHandler)

	port := ":8080"
	fmt.Printf("Auth service running on http://localhost%s\n", port)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server error: %v\n", err)
	}
}
