package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tapiaw38/irrigation-api/models"
	"github.com/tapiaw38/irrigation-api/repository"
	"github.com/tapiaw38/irrigation-api/server"
)

// CreateProducersHandler handles the request to get all producers
func CreateProducersHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var producers []models.Producer

		err := json.NewDecoder(r.Body).Decode(&producers)

		if err != nil {
			http.Error(w, "An error ocurred when trying to enter an producers "+err.Error(), 400)
			return
		}

		defer r.Body.Close()

		ctx := r.Context()

		producers, err = repository.CreateProducers(ctx, producers)

		if err != nil {
			http.Error(w, "An error occurred when trying to create producers in database "+err.Error(), 400)
			return
		}

		// send create producer websocket message
		var producerMessage = models.WebsocketMessage{
			Type:    "producer_created",
			Payload: producers,
		}
		s.Hub().Broadcast(producerMessage, nil)

		response := NewResponse(Message, "ok", producers)
		ResponseWithJson(w, response, http.StatusOK)
	}
}

// GetProducersHandler handles the request to get all producers
func GetProducersHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		producers, err := repository.GetProducers(ctx)

		if err != nil {
			http.Error(w, "An error occurred when trying to get producers in database "+err.Error(), 400)
			return
		}

		response := NewResponse(Message, "ok", producers)
		ResponseWithJson(w, response, http.StatusOK)
	}
}

// GetProducerByIdHandler handles the request to get a producer by id
func GetProducerByIDHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		id := mux.Vars(r)["id"]

		if id == "" {
			http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
			return
		}

		var producer *models.Producer = s.Redis().GetProducer(id)
		if producer == nil {
			producer, err := repository.GetProducerByID(ctx, id)
			if err != nil {
				http.Error(w, "An error occurred when trying to get producer in database "+err.Error(), 400)
				return
			}
			s.Redis().SetProducer(id, &producer)
			response := NewResponse(Message, "ok", producer)
			ResponseWithJson(w, response, http.StatusOK)
			return
		}
		response := NewResponse(Message, "ok", producer)
		ResponseWithJson(w, response, http.StatusOK)
	}
}

// UpdateProducerHandler handles the request to update a producer
func UpdateProducerHandler(w http.ResponseWriter, r *http.Request) {

	var pr models.Producer

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&pr)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a producer "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	pr, err = repository.UpdateProducer(ctx, id, pr)

	if err != nil {
		http.Error(w, "An error occurred when trying to update a producer in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", pr)
	ResponseWithJson(w, response, http.StatusOK)
}

// PartialUpdateProducerHandler handles the request to update a producer
func PartialUpdateProducerHandler(w http.ResponseWriter, r *http.Request) {

	var pr models.Producer

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&pr)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a producer "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	pr, err = repository.PartialUpdateProducer(ctx, id, pr)

	if err != nil {
		http.Error(w, "An error occurred when trying to update a producer in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", pr)
	ResponseWithJson(w, response, http.StatusOK)
}

// DeleteProducerHandler handles the request to delete a producer
func DeleteProducerHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	producer, err := repository.DeleteProducer(ctx, id)

	if err != nil {
		http.Error(w, "An error occurred when trying to delete a producer in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", producer)
	ResponseWithJson(w, response, http.StatusOK)
}
