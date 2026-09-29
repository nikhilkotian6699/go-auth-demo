package repository

import (
	"errors"
	"sync"

	"github.com/nikhilkotian6699/go-auth-demo/models"
)

var (
	// ErrUserNotFound is returned when a user cannot be found
	ErrUserNotFound = errors.New("user not found")
	// ErrUserAlreadyExists is returned when attempting to register a username or email that is taken
	ErrUserAlreadyExists = errors.New("user with this username or email already exists")
)

// UserRepository defines the persistence contract for users
type UserRepository interface {
	Create(user *models.User) error
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByID(id string) (*models.User, error)
}

// InMemoryUserRepository provides an in-memory thread-safe implementation of UserRepository
type InMemoryUserRepository struct {
	mu    sync.RWMutex
	users map[string]*models.User
}

// NewInMemoryUserRepository initializes a new InMemoryUserRepository
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*models.User),
	}
}

// Create persists a user to the in-memory store
func (r *InMemoryUserRepository) Create(user *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, u := range r.users {
		if u.Username == user.Username || u.Email == user.Email {
			return ErrUserAlreadyExists
		}
	}

	r.users[user.ID] = user
	return nil
}

// FindByUsername retrieves a user by username
func (r *InMemoryUserRepository) FindByUsername(username string) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, u := range r.users {
		if u.Username == username {
			return u, nil
		}
	}
	return nil, ErrUserNotFound
}

// FindByEmail retrieves a user by email
func (r *InMemoryUserRepository) FindByEmail(email string) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, ErrUserNotFound
}

// FindByID retrieves a user by ID
func (r *InMemoryUserRepository) FindByID(id string) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, ErrUserNotFound
	}
	return user, nil
}
