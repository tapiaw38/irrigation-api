package models

import (
	"time"
)

// Section is the structure that holds the information about a Section.
type Section struct {
	ID            int64     `json:"id,omitempty"`
	SectionNumber string    `json:"section_number,omitempty"`
	Name          string    `json:"name,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}
