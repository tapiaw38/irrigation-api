package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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

// GetTurnByIDHandler is a function to get a Turn by ID.
func (tr *TurnRouter) GetTurnByIDHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred when trying to get turn in database", 400)
		return
	}

	ctx := r.Context()

	turn, err := tr.Storage.GetTurnByID(ctx, id)

	if err != nil {
		http.Error(w, "An error occurred when trying to get turn in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", turn)
	ResponseWithJson(w, response, http.StatusOK)
}

// UpdateTurnHandler is a function to update a Turn.
func (tr *TurnRouter) UpdateTurnHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred when trying to get turn in database", 400)
		return
	}

	var turn turn.Turn

	err := json.NewDecoder(r.Body).Decode(&turn)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter an turn "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	trs, err := tr.Storage.UpdateTurn(ctx, id, turn)

	if err != nil {
		http.Error(w, "An error occurred when trying to update turn in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", trs)
	ResponseWithJson(w, response, http.StatusOK)
}

// DeleteTurnHandler is a function to delete a Turn.
func (tr *TurnRouter) DeleteTurnHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred when trying to get turn in database", 400)
		return
	}

	ctx := r.Context()

	trs, err := tr.Storage.DeleteTurn(ctx, id)

	if err != nil {
		http.Error(w, "An error occurred when trying to delete turn in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", trs)
	ResponseWithJson(w, response, http.StatusOK)
}

// CreateTurnProductionHandler is a function to create a new TurnProduction.
func (tr *TurnRouter) CreateTurnProductionHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred when trying to get turn in database", 400)
		return
	}

	var turnProduction turn.TurnProduction

	err := json.NewDecoder(r.Body).Decode(&turnProduction)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter an turn "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	turn, err := tr.Storage.CreateTurnProduction(ctx, id, turnProduction)

	if err != nil {
		http.Error(w, "An error occurred when trying to create turn in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", turn)
	ResponseWithJson(w, response, http.StatusOK)
}

// DeleteTurnProductionHandler is a function to delete a TurnProduction.
func (tr *TurnRouter) DeleteTurnProductionHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred when trying to get turn in database", 400)
		return
	}

	var turnProduction turn.TurnProduction

	err := json.NewDecoder(r.Body).Decode(&turnProduction)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter an turn "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	turn, err := tr.Storage.DeleteTurnProduction(ctx, id, turnProduction)

	if err != nil {
		http.Error(w, "An error occurred when trying to delete turn in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", turn)
	ResponseWithJson(w, response, http.StatusOK)
}
