package models

import "time"

// Producer is the structure that holds the information about a Producer.
type Producer struct {
	ID             int64     `json:"id,omitempty"`
	FirstName      string    `json:"first_name,omitempty"`
	LastName       string    `json:"last_name,omitempty"`
	DocumentNumber string    `json:"document_number,omitempty"`
	BirthDate      string    `json:"birth_date,omitempty"`
	PhoneNumber    string    `json:"phone_number,omitempty"`
	Address        string    `json:"address,omitempty"`
	IsActive       bool      `json:"is_active,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}
