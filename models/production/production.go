package production

import "github.com/tapiaw38/irrigation-api/models/producer"

// Production is the struct that holds the production data
type Production struct {
	ID             uint              `json:"id,omitempty"`
	Producer       producer.Producer `json:"producer,omitempty"`
	LoteNumber     string            `json:"lote_number,omitempty"`
	Entry          string            `json:"entry,omitempty"`
	Name           string            `json:"name,omitempty"`
	ProductionType string            `json:"production_type,omitempty"`
	Latitude       float64           `json:"latitude,omitempty"`
	Longitude      float64           `json:"longitude,omitempty"`
	Picture        string            `json:"picture,omitempty"`
	CreatedAt      string            `json:"created_at,omitempty"`
	UpdatedAt      string            `json:"updated_at,omitempty"`
}
