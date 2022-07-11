package turn

import "time"

// Turn is the structure that holds the information about a Turn.
type Turn struct {
	ID        int64     `json:"id,omitempty"`
	StartDate string    `json:"start_date,omitempty"`
	EndDate   string    `json:"end_date,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
