package models

type WebsocketMessage struct {
	Type    string      `json:"type,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}
