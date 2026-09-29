package models

import "time"

// User represents an application user
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

// RegisterInput represents the input required for registration
type RegisterInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginInput represents the input required for user login
type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
