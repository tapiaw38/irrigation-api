package routers

import (
	"encoding/json"
	"net/http"

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
