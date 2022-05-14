package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tapiaw38/irrigation-api/models/producer"
)

type ProducerRouter struct {
	Storage producer.Storage
}

// CreateProducersHandler handles the request to get all producers
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

// UpdateProducerHandler handles the request to update a producer
func (pr *ProducerRouter) UpdateProducerHandler(w http.ResponseWriter, r *http.Request) {

	var producer producer.Producer

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&producer)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a producer "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	producer, err = pr.Storage.UpdateProducer(ctx, id, producer)

	if err != nil {
		http.Error(w, "An error occurred when trying to update a producer in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", producer)
	ResponseWithJson(w, response, http.StatusOK)
}

// PartialUpdateProducerHandler handles the request to update a producer
func (pr *ProducerRouter) PartialUpdateProducerHandler(w http.ResponseWriter, r *http.Request) {

	var producer producer.Producer

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&producer)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a producer "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	producer, err = pr.Storage.PartialUpdateProducer(ctx, id, producer)

	if err != nil {
		http.Error(w, "An error occurred when trying to update a producer in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", producer)
	ResponseWithJson(w, response, http.StatusOK)
}
