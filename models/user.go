package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User is the user model
type User struct {
	ID           uint      `json:"id,omitempty"`
	FirstName    string    `json:"first_name,omitempty"`
	LastName     string    `json:"last_name,omitempty"`
	Username     string    `json:"username,omitempty"`
	Email        string    `json:"email,omitempty"`
	Picture      string    `json:"picture,omitempty"`
	PhoneNumber  string    `json:"phone_number,omitempty"`
	Address      string    `json:"address,omitempty"`
	Password     string    `json:"password,omitempty"`
	PasswordHash string    `json:"-"`
	IsActive     bool      `json:"is_active,omitempty"`
	IsAdmin      bool      `json:"is_admin,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

// LoginResponse is the login response
type LoginResponse struct {
	User        interface{} `json:"user"`
	AccessToken string      `json:"access_token"`
}

// HashPassword hashes the password
func (u *User) HashPassword() error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(passwordHash)

	return nil
}

// PasswordMatch checks the password
func (u *User) PasswordMatch(password string) bool {
	passwordBytes := []byte(password)
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), passwordBytes)

	return err == nil
}
