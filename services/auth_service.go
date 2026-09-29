package services

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"github.com/nikhilkotian6699/go-auth-demo/models"
	"github.com/nikhilkotian6699/go-auth-demo/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrInvalidCredentials indicates incorrect username or password
	ErrInvalidCredentials = errors.New("invalid username or password")
	// ErrEmptyFields indicates required registration or login fields are missing
	ErrEmptyFields = errors.New("all fields are required")
	// ErrShortPassword indicates password is too short
	ErrShortPassword = errors.New("password must be at least 6 characters long")
)

// AuthService handles authentication business logic
type AuthService struct {
	userRepo repository.UserRepository
}

// NewAuthService creates a new AuthService
func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

// Register validates input, hashes password, and creates a new user
func (s *AuthService) Register(input models.RegisterInput) (*models.User, error) {
	username := strings.TrimSpace(input.Username)
	email := strings.TrimSpace(strings.ToLower(input.Email))
	password := input.Password

	if username == "" || email == "" || password == "" {
		return nil, ErrEmptyFields
	}

	if len(password) < 6 {
		return nil, ErrShortPassword
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	idBytes := make([]byte, 16)
	if _, err := rand.Read(idBytes); err != nil {
		return nil, err
	}
	id := hex.EncodeToString(idBytes)

	user := &models.User{
		ID:        id,
		Username:  username,
		Email:     email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// Login verifies user credentials and returns the authenticated user
func (s *AuthService) Login(input models.LoginInput) (*models.User, error) {
	username := strings.TrimSpace(input.Username)
	password := input.Password

	if username == "" || password == "" {
		return nil, ErrEmptyFields
	}

	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}
