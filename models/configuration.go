package models

import "time"

// Location represents geographical location settings
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Zoom      int     `json:"zoom"`
}

// Configuration represents the system configuration settings
type Configuration struct {
	ID                 string    `json:"id,omitempty"`
	DefaultLocation    Location  `json:"default_location"`
	WateringHourFactor float64   `json:"watering_hour_factor"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
}

// ConfigurationRequest represents the request body for configuration
type ConfigurationRequest struct {
	DefaultLocation    Location `json:"default_location" validate:"required"`
	WateringHourFactor float64  `json:"watering_hour_factor" validate:"required,gt=0"`
}
