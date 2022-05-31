package intake

import "github.com/tapiaw38/irrigation-api/models/section"

// Intake is the struct that holds the intake data
type Intake struct {
	ID           int64   `json:"id,omitempty"`
	Section      int64   `json:"section,omitempty"`
	IntakeNumber string  `json:"intake_number,omitempty"`
	Name         string  `json:"name,omitempty"`
	Latitude     float64 `json:"latitude,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
	CreatedAt    string  `json:"created_at,omitempty"`
	UpdatedAt    string  `json:"updated_at,omitempty"`
}

// IntakeResponse is the struct that holds the intake data
type IntakeResponse struct {
	ID           int64           `json:"id,omitempty"`
	Section      section.Section `json:"section,omitempty"`
	IntakeNumber string          `json:"intake_number,omitempty"`
	Name         string          `json:"name,omitempty"`
	Latitude     float64         `json:"latitude,omitempty"`
	Longitude    float64         `json:"longitude,omitempty"`
	CreatedAt    string          `json:"created_at,omitempty"`
	UpdatedAt    string          `json:"updated_at,omitempty"`
}
