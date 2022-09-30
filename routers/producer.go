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
			http.Error(w, "An error ocurred when trying to enter an producers "+err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		ctx := r.Context()
		producers, err = repository.CreateProducers(ctx, producers)
		if err != nil {
			http.Error(w, "An error occurred when trying to create producers in database "+err.Error(), http.StatusBadRequest)
			return
		}
		// send producer message created websocket
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
		// get producers cache
		/*
			var producers *[]models.Producer = s.Redis().GetProducers("producers")
			if producers == nil {
				producers, err := repository.GetProducers(ctx)
				if err != nil {
					http.Error(w, "An error occurred when trying to get producers in database "+err.Error(), http.StatusBadRequest)
					return
				}
				s.Redis().SetProducers("producers", &producers)
				response := NewResponse(Message, "ok", producers)
				ResponseWithJson(w, response, http.StatusOK)
				return
			}
			response := NewResponse(Message, "ok", producers)
			ResponseWithJson(w, response, http.StatusOK)
		*/

		producers, err := repository.GetProducers(ctx)
		if err != nil {
			http.Error(w, "An error occurred when trying to get producers in database "+err.Error(), http.StatusBadRequest)
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
		// get producer cache
		/*
			var producer *models.Producer = s.Redis().GetProducer(id)
			if producer == nil {
				producer, err := repository.GetProducerByID(ctx, id)
				if err != nil {
					http.Error(w, "An error occurred when trying to get producer in database "+err.Error(), http.StatusBadRequest)
					return
				}
				s.Redis().SetProducer(id, &producer)
				response := NewResponse(Message, "ok", producer)
				ResponseWithJson(w, response, http.StatusOK)
				return
			}
			response := NewResponse(Message, "ok", producer)
			ResponseWithJson(w, response, http.StatusOK)
		*/

		producer, err := repository.GetProducerByID(ctx, id)
		if err != nil {
			http.Error(w, "An error occurred when trying to get producer in database "+err.Error(), http.StatusBadRequest)
			return
		}
		response := NewResponse(Message, "ok", producer)
		ResponseWithJson(w, response, http.StatusOK)
	}
}

// UpdateProducerHandler handles the request to update a producer
func UpdateProducerHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		if id == "" {
			http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
			return
		}
		var producer models.Producer
		err := json.NewDecoder(r.Body).Decode(&producer)
		if err != nil {
			http.Error(w, "An error occurred when trying to enter a producer "+err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		ctx := r.Context()
		producer, err = repository.UpdateProducer(ctx, id, producer)
		if err != nil {
			http.Error(w, "An error occurred when trying to update a producer in database "+err.Error(), http.StatusBadRequest)
			return
		}
		response := NewResponse(Message, "ok", producer)
		ResponseWithJson(w, response, http.StatusOK)
	}
}

// PartialUpdateProducerHandler handles the request to update a producer
func PartialUpdateProducerHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		if id == "" {
			http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
			return
		}
		var producer models.Producer
		err := json.NewDecoder(r.Body).Decode(&producer)
		if err != nil {
			http.Error(w, "An error occurred when trying to enter a producer "+err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		ctx := r.Context()
		producer, err = repository.PartialUpdateProducer(ctx, id, producer)
		if err != nil {
			http.Error(w, "An error occurred when trying to update a producer in database "+err.Error(), http.StatusBadRequest)
			return
		}
		response := NewResponse(Message, "ok", producer)
		ResponseWithJson(w, response, http.StatusOK)
	}
}

// DeleteProducerHandler handles the request to delete a producer
func DeleteProducerHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		if id == "" {
			http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
			return
		}
		ctx := r.Context()
		producer, err := repository.DeleteProducer(ctx, id)
		if err != nil {
			http.Error(w, "An error occurred when trying to delete a producer in database "+err.Error(), http.StatusBadRequest)
			return
		}
		response := NewResponse(Message, "ok", producer)
		ResponseWithJson(w, response, http.StatusOK)
	}
}
