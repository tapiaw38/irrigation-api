package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tapiaw38/irrigation-api/models"
	"github.com/tapiaw38/irrigation-api/repository"
)

// CreateIntakesHandler handles the request to get all intakes
func CreateIntakesHandler(w http.ResponseWriter, r *http.Request) {

	var intakes []models.Intake

	err := json.NewDecoder(r.Body).Decode(&intakes)

	if err != nil {
		http.Error(w, "An error ocurred when trying to enter an intakes "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	intakes, err = repository.CreateIntakes(ctx, intakes)

	if err != nil {
		http.Error(w, "An error occurred when trying to create intakes in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", intakes)
	ResponseWithJson(w, response, http.StatusOK)
}

// GetIntakesHandler handles the request to get all intakes
func GetIntakesHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	intakes, err := repository.GetIntakes(ctx)

	if err != nil {
		http.Error(w, "An error occurred when trying to get intakes in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", intakes)
	ResponseWithJson(w, response, http.StatusOK)
}

// GetIntakeByIDHandler handles the request to get a intake by id
func GetIntakeByIDHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred when trying to get intake by id", 400)
		return
	}

	ctx := r.Context()

	intake, err := repository.GetIntakeByID(ctx, id)

	if err != nil {
		http.Error(w, "An error occurred when trying to get intakes in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", intake)
	ResponseWithJson(w, response, http.StatusOK)
}

// UpdateIntakeHandler handles the request to get a intake by id
func UpdateIntakeHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred when trying to get intakes in database", 400)
		return
	}

	var intake models.Intake

	err := json.NewDecoder(r.Body).Decode(&intake)

	if err != nil {
		http.Error(w, "An error occurred when trying to get intakes in database "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	itk, err := repository.UpdateIntake(ctx, id, intake)

	if err != nil {
		http.Error(w, "An error occurred when trying to get intakes in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", itk)
	ResponseWithJson(w, response, http.StatusOK)
}

// DeleteIntakeHandler handles the request to get a intake by id
func DeleteIntakeHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred when trying to get intakes in database", 400)
		return
	}

	ctx := r.Context()

	intake, err := repository.DeleteIntake(ctx, id)

	if err != nil {
		http.Error(w, "An error occurred when trying to get intakes in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", intake)
	ResponseWithJson(w, response, http.StatusOK)
}

// CreateIntakeProductionHandler handles the request to get a intake by id
func CreateIntakeProductionHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred when trying to get intakes in database", 400)
		return
	}

	var intakeProduction models.IntakeProduction

	err := json.NewDecoder(r.Body).Decode(&intakeProduction)

	if err != nil {
		http.Error(w, "An error occurred when trying to get intakes in database "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	intake, err := repository.CreateIntakeProduction(ctx, id, intakeProduction)

	if err != nil {
		http.Error(w, "An error occurred when trying to get intakes in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", intake)
	ResponseWithJson(w, response, http.StatusOK)
}

// UpdateIntakeProductionHandler handles the request to get a intake by id
func UpdateIntakeProductionHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred when trying to get intakes in database", 400)
		return
	}

	var intakeProduction models.IntakeProduction

	err := json.NewDecoder(r.Body).Decode(&intakeProduction)

	if err != nil {
		http.Error(w, "An error occurred when trying to get intakes in database "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	intake, err := repository.UpdateIntakeProduction(ctx, id, intakeProduction)

	if err != nil {
		http.Error(w, "An error occurred when trying to get intakes in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", intake)
	ResponseWithJson(w, response, http.StatusOK)
}

// DeleteIntakeProductionHandler handles the request to get a intake by id
func DeleteIntakeProductionHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred when trying to get intakes in database", 400)
		return
	}

	var intakeProduction models.IntakeProduction

	err := json.NewDecoder(r.Body).Decode(&intakeProduction)

	if err != nil {
		http.Error(w, "An error occurred when trying to get intakes in database "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	intake, err := repository.DeleteIntakeProduction(ctx, id, intakeProduction)

	if err != nil {
		http.Error(w, "An error occurred when trying to get intakes in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", intake)
	ResponseWithJson(w, response, http.StatusOK)
}
