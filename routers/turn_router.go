package routers

import (
	"encoding/json"
	"net/http"

	"github.com/tapiaw38/irrigation-api/models/turn"
)

type TurnRouter struct {
	Storage turn.Storage
}

// CreateTurnHandler is a function to create a new Turn.
func (tr *TurnRouter) CreateTurnHandler(w http.ResponseWriter, r *http.Request) {

	var turn turn.Turn

	err := json.NewDecoder(r.Body).Decode(&turn)

	if err != nil {
		http.Error(w, "An error ocurred when trying to enter an turn "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	turn, err = tr.Storage.CreateTurn(ctx, turn)

	if err != nil {
		http.Error(w, "An error occurred when trying to create turn in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", turn)
	ResponseWithJson(w, response, http.StatusOK)
}

// GetTurnHandler is a function to get a Turn.
func (tr *TurnRouter) GetTurnsHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	turns, err := tr.Storage.GetTurns(ctx)

	if err != nil {
		http.Error(w, "An error occurred when trying to get turn in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", turns)
	ResponseWithJson(w, response, http.StatusOK)
}
