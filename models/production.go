package models

// Coordinate represents a geographic coordinate
type Coordinate struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Production is the struct that holds the production data
type Production struct {
	ID                         int64   `json:"id,omitempty"`
	Producer                   int64   `json:"producer,omitempty"`
	LoteNumber                 string  `json:"lote_number,omitempty"`
	Entry                      string  `json:"entry,omitempty"`
	Name                       string  `json:"name,omitempty"`
	ProductionType             string  `json:"production_type,omitempty"`
	Area                       float64 `json:"area,omitempty"`
	CultivatedArea             float64 `json:"cultivated_area,omitempty"`
	AreaCoordinates            string  `json:"area_coordinates,omitempty"` // JSON string
	CultivatedAreaCoordinates  string  `json:"cultivated_area_coordinates,omitempty"` // JSON string
	Picture                    string  `json:"picture,omitempty"`
	CadastralRegistration      string  `json:"cadastral_registration,omitempty"`
	District                   string  `json:"district,omitempty"`
	CreatedAt                  string  `json:"created_at,omitempty"`
	UpdatedAt                  string  `json:"updated_at,omitempty"`
}

// ProductionResponse is the struct that holds the production data
type ProductionResponse struct {
	ID                        int64    `json:"id,omitempty"`
	Producer                  Producer `json:"producer,omitempty"`
	LoteNumber                string   `json:"lote_number,omitempty"`
	Entry                     string   `json:"entry,omitempty"`
	Name                      string   `json:"name,omitempty"`
	ProductionType            string   `json:"production_type,omitempty"`
	Area                      float64  `json:"area,omitempty"`
	CultivatedArea            float64  `json:"cultivated_area,omitempty"`
	AreaCoordinates           string   `json:"area_coordinates,omitempty"`
	CultivatedAreaCoordinates string   `json:"cultivated_area_coordinates,omitempty"`
	Picture                   string   `json:"picture,omitempty"`
	CadastralRegistration     string   `json:"cadastral_registration,omitempty"`
	District                  string   `json:"district,omitempty"`
	CreatedAt                 string   `json:"created_at,omitempty"`
	UpdatedAt                 string   `json:"updated_at,omitempty"`
}

// ProductionIntakeResponse is the struct that holds the production intake data
type ProductionIntakeResponse struct {
	ProductionResponse
	WateringOrder int64 `json:"watering_order,omitempty"`
}

// ProductionTurnResponse is the struct that holds the production turn data
type ProductionTurnResponse struct {
	ProductionResponse
	IntakeID      int64   `json:"intake_id,omitempty"`
	IntakeNumber  string  `json:"intake_number,omitempty"`
	WateringOrder int64   `json:"watering_order,omitempty"`
	WateringHour  float64 `json:"watering_hour,omitempty"`
}
