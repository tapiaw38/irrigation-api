package routers

import (
	"encoding/json"
	"net/http"

	"github.com/tapiaw38/irrigation-api/models/producer"
)

type ProducerRouter struct {
	Storage producer.Storage
}

func (pr *ProducerRouter) CreateProducersHandler(w http.ResponseWriter, r *http.Request) {

	var producers []producer.Producer

	err := json.NewDecoder(r.Body).Decode(&producers)

	if err != nil {
		http.Error(w, "An error ocurred when trying to enter an producers "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	producers, err = pr.Storage.CreateProducers(ctx, producers)

	if err != nil {
		http.Error(w, "An error occurred when trying to create producers in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", producers)
	ResponseWithJson(w, response, http.StatusOK)
}

// GetProducersHandler handles the request to get all producers
func (pr *ProducerRouter) GetProducersHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	producers, err := pr.Storage.GetProducers(ctx)

	if err != nil {
		http.Error(w, "An error occurred when trying to get producers in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", producers)
	ResponseWithJson(w, response, http.StatusOK)
}
