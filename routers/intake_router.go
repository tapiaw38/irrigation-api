package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tapiaw38/irrigation-api/models/intake"
)

type IntakeRouter struct {
	Storage intake.Storage
}

// CreateIntakesHandler handles the request to get all intakes
func (ik *IntakeRouter) CreateIntakesHandler(w http.ResponseWriter, r *http.Request) {

	var intakes []intake.Intake

	err := json.NewDecoder(r.Body).Decode(&intakes)

	if err != nil {
		http.Error(w, "An error ocurred when trying to enter an intakes "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	intakes, err = ik.Storage.CreateIntakes(ctx, intakes)

	if err != nil {
		http.Error(w, "An error occurred when trying to create intakes in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", intakes)
	ResponseWithJson(w, response, http.StatusOK)
}

// GetIntakesHandler handles the request to get all intakes
func (ik *IntakeRouter) GetIntakesHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	intakes, err := ik.Storage.GetIntakes(ctx)

	if err != nil {
		http.Error(w, "An error occurred when trying to get intakes in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", intakes)
	ResponseWithJson(w, response, http.StatusOK)
}

// GetIntakeByIDHandler handles the request to get a intake by id
func (ik *IntakeRouter) GetIntakeByIDHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred when trying to get intake by id", 400)
		return
	}

	ctx := r.Context()

	intake, err := ik.Storage.GetIntakeByID(ctx, id)

	if err != nil {
		http.Error(w, "An error occurred when trying to get intakes in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", intake)
	ResponseWithJson(w, response, http.StatusOK)
}

// UpdateIntakeHandler handles the request to get a intake by id
func (ik *IntakeRouter) UpdateIntakeHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred when trying to get intakes in database", 400)
		return
	}

	var intake intake.Intake

	err := json.NewDecoder(r.Body).Decode(&intake)

	if err != nil {
		http.Error(w, "An error occurred when trying to get intakes in database "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	itk, err := ik.Storage.UpdateIntake(ctx, id, intake)

	if err != nil {
		http.Error(w, "An error occurred when trying to get intakes in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", itk)
	ResponseWithJson(w, response, http.StatusOK)
}

// DeleteIntakeHandler handles the request to get a intake by id
func (ik *IntakeRouter) DeleteIntakeHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred when trying to get intakes in database", 400)
		return
	}

	ctx := r.Context()

	intake, err := ik.Storage.DeleteIntake(ctx, id)

	if err != nil {
		http.Error(w, "An error occurred when trying to get intakes in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", intake)
	ResponseWithJson(w, response, http.StatusOK)
}
