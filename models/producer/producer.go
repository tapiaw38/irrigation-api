package producer

import "time"

type Producer struct {
	ID             uint      `json:"id,omitempty"`
	FirstName      string    `json:"first_name,omitempty"`
	LastName       string    `json:"last_name,omitempty"`
	DocumentNumber string    `json:"document_number,omitempty"`
	BirthDate      string    `json:"birth_date,omitempty"`
	PhoneNumber    string    `json:"phone_number,omitempty"`
	Address        string    `json:"address,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}
