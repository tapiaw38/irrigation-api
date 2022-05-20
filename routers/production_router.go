package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tapiaw38/irrigation-api/models/production"
)

type ProductionRouter struct {
	Storage production.Storage
}

// CreateProductionsHandler handles the request to get all productions
func (pd *ProductionRouter) CreateProductionHandler(w http.ResponseWriter, r *http.Request) {

	var productions []production.Production

	err := json.NewDecoder(r.Body).Decode(&productions)

	if err != nil {
		http.Error(w, "An error ocurred when trying to enter an productions "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	productions, err = pd.Storage.CreateProductions(ctx, productions)

	if err != nil {
		http.Error(w, "An error occurred when trying to create productions in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", productions)
	ResponseWithJson(w, response, http.StatusOK)
}

// GetProductionsHandler handles the request to get all productions
func (pd *ProductionRouter) GetProductionsHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	productions, err := pd.Storage.GetProductions(ctx)

	if err != nil {
		http.Error(w, "An error occurred when trying to get productions in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", productions)
	ResponseWithJson(w, response, http.StatusOK)
}

// GetProductionByIDHandler handles the request to get a production by id
func (pd *ProductionRouter) GetProductionByIDHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred when trying to get production by id", 400)
		return
	}

	ctx := r.Context()

	pds, err := pd.Storage.GetProductionsByID(ctx, id)

	if err != nil {
		http.Error(w, "An error occurred when trying to get production by id "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", pds)
	ResponseWithJson(w, response, http.StatusOK)
}

// UpdateProductionHandler handles the request to update a production
func (pd *ProductionRouter) UpdateProductionHandler(w http.ResponseWriter, r *http.Request) {

	var production production.Production

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&production)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a production "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	pds, err := pd.Storage.UpdateProduction(ctx, id, production)

	if err != nil {
		http.Error(w, "An error occurred when trying to update a production in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", pds)
	ResponseWithJson(w, response, http.StatusOK)
}

// DeleteProductionHandler handles the request to delete a production
func (pd *ProductionRouter) DeleteProductionHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", 400)
		return
	}

	ctx := r.Context()

	pds, err := pd.Storage.DeleteProduction(ctx, id)

	if err != nil {
		http.Error(w, "An error occurred when trying to delete a production in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", pds)
	ResponseWithJson(w, response, http.StatusOK)
}
