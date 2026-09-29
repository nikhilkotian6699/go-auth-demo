package handlers

import (
	"encoding/json"
	"github.com/nikhilkotian6699/go-auth-demo/models"
	"github.com/nikhilkotian6699/go-auth-demo/services"
	"net/http"
	"strings"
)

// AuthHandler handles user registration and authentication requests
type AuthHandler struct {
	authService *services.AuthService
}

// NewAuthHandler creates a new AuthHandler instance
func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register handles user registration
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var input models.RegisterInput

	// Support both JSON payloads and Form submissions
	if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid JSON body", http.StatusBadRequest)
			return
		}
	} else {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}
		input.Username = r.FormValue("username")
		input.Email = r.FormValue("email")
		input.Password = r.FormValue("password")
	}

	user, err := h.authService.Register(input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered successfully",
		"user": map[string]string{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

// Login handles user authentication
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var input models.LoginInput

	// Support both JSON payloads and Form submissions
	if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid JSON body", http.StatusBadRequest)
			return
		}
	} else {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}
		input.Username = r.FormValue("username")
		input.Password = r.FormValue("password")
	}

	user, err := h.authService.Login(input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login successful",
		"user": map[string]string{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}
