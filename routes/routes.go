package routes

import (
	"encoding/json"
	"net/http"

	"github.com/nikhilkotian6699/go-auth-demo/handlers"
)

// SetupRoutes registers authentication API endpoints
func SetupRoutes(mux *http.ServeMux, authHandler *handlers.AuthHandler) {
	// Root health check endpoint
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "running",
			"service": "Go Auth Service",
		})
	})

	// Auth routes
	mux.HandleFunc("/register", authHandler.Register)
	mux.HandleFunc("/login", authHandler.Login)
}
