package turn

import (
	"time"

	"github.com/tapiaw38/irrigation-api/models/production"
)

// Turn is the structure that holds the information about a Turn.
type Turn struct {
	ID        int64     `json:"id,omitempty"`
	StartDate string    `json:"start_date,omitempty"`
	TurnHours float64   `json:"turn_hours,omitempty"`
	EndDate   string    `json:"end_date,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// TunrResponse is the structure that holds the information about a Turn.
type TurnResponse struct {
	ID          int64                               `json:"id,omitempty"`
	StartDate   string                              `json:"start_date,omitempty"`
	TurnHours   float64                             `json:"turn_hours,omitempty"`
	EndDate     string                              `json:"end_date,omitempty"`
	Productions []production.ProductionTurnResponse `json:"productions,omitempty"`
	CreatedAt   time.Time                           `json:"created_at,omitempty"`
	UpdatedAt   time.Time                           `json:"updated_at,omitempty"`
}

// TurnProduction is the structure that holds the information about a TurnProduction.
type TurnProduction struct {
	TurnID       string `json:"turn_id,omitempty"`
	ProductionID string `json:"production_id,omitempty"`
}
