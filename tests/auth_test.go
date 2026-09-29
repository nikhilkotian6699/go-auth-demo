package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nikhilkotian6699/go-auth-demo/handlers"
	"github.com/nikhilkotian6699/go-auth-demo/models"
	"github.com/nikhilkotian6699/go-auth-demo/repository"
	"github.com/nikhilkotian6699/go-auth-demo/services"
)

func setupAuthService() (*services.AuthService, repository.UserRepository) {
	repo := repository.NewInMemoryUserRepository()
	authService := services.NewAuthService(repo)
	return authService, repo
}

func TestUserRegistration(t *testing.T) {
	authService, _ := setupAuthService()

	t.Run("successful registration", func(t *testing.T) {
		input := models.RegisterInput{
			Username: "testuser",
			Email:    "test@example.com",
			Password: "securepassword123",
		}

		user, err := authService.Register(input)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}

		if user.Username != input.Username {
			t.Errorf("expected username %s, got %s", input.Username, user.Username)
		}

		if user.Email != input.Email {
			t.Errorf("expected email %s, got %s", input.Email, user.Email)
		}

		if user.Password == input.Password {
			t.Error("expected password to be hashed, but was stored as plaintext")
		}
	})

	t.Run("duplicate registration fails", func(t *testing.T) {
		input := models.RegisterInput{
			Username: "testuser",
			Email:    "test2@example.com",
			Password: "password123",
		}

		_, err := authService.Register(input)
		if err == nil {
			t.Fatal("expected error for duplicate username, got nil")
		}
	})

	t.Run("short password fails", func(t *testing.T) {
		input := models.RegisterInput{
			Username: "shortpassuser",
			Email:    "short@example.com",
			Password: "123",
		}

		_, err := authService.Register(input)
		if err != services.ErrShortPassword {
			t.Fatalf("expected ErrShortPassword, got: %v", err)
		}
	})
}

func TestUserLogin(t *testing.T) {
	authService, _ := setupAuthService()

	// Register user first
	_, err := authService.Register(models.RegisterInput{
		Username: "loginuser",
		Email:    "login@example.com",
		Password: "correctpassword",
	})
	if err != nil {
		t.Fatalf("failed to register user for login test: %v", err)
	}

	t.Run("valid credentials", func(t *testing.T) {
		user, err := authService.Login(models.LoginInput{
			Username: "loginuser",
			Password: "correctpassword",
		})
		if err != nil {
			t.Fatalf("expected successful login, got: %v", err)
		}

		if user.Username != "loginuser" {
			t.Errorf("expected username loginuser, got: %s", user.Username)
		}
	})

	t.Run("invalid password", func(t *testing.T) {
		_, err := authService.Login(models.LoginInput{
			Username: "loginuser",
			Password: "wrongpassword",
		})
		if err != services.ErrInvalidCredentials {
			t.Fatalf("expected ErrInvalidCredentials, got: %v", err)
		}
	})

	t.Run("non-existent user", func(t *testing.T) {
		_, err := authService.Login(models.LoginInput{
			Username: "nonexistent",
			Password: "correctpassword",
		})
		if err != services.ErrInvalidCredentials {
			t.Fatalf("expected ErrInvalidCredentials, got: %v", err)
		}
	})
}

func TestAuthHandlersHTTP(t *testing.T) {
	authService, _ := setupAuthService()
	authHandler := handlers.NewAuthHandler(authService)

	t.Run("POST /register HTTP endpoint", func(t *testing.T) {
		payload, _ := json.Marshal(models.RegisterInput{
			Username: "apiuser",
			Email:    "api@example.com",
			Password: "password123",
		})

		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		authHandler.Register(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("expected status %d, got %d", http.StatusCreated, w.Code)
		}
	})

	t.Run("POST /login HTTP endpoint", func(t *testing.T) {
		payload, _ := json.Marshal(models.LoginInput{
			Username: "apiuser",
			Password: "password123",
		})

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		authHandler.Login(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
		}
	})
}
