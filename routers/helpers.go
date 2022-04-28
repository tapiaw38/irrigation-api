package routers

import (
	"encoding/json"
	"net/http"
)

const (
	Error   = "error"
	Message = "message"
)

// Response is a response struct
type Response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Response    interface{} `json:"response"`
}

// NewResponse is a helper function to create a new response
func NewResponse(messageType string, message string, response interface{}) Response {
	return Response{
		MessageType: messageType,
		Message:     message,
		Response:    response,
	}
}

// ResponseWithJSON is a helper function to send a response with json
func ResponseWithJson(w http.ResponseWriter, response Response, statusCode int) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
