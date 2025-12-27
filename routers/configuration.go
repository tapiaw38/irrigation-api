package routers

import (
	"encoding/json"
	"net/http"

	"github.com/tapiaw38/irrigation-api/models"
	"github.com/tapiaw38/irrigation-api/repository"
)

// GetConfigurationHandler is a function to get the system configuration.
func GetConfigurationHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	config, err := repository.GetConfiguration(ctx)

	if err != nil {
		http.Error(w, "An error occurred when trying to get configuration from database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", config)
	ResponseWithJson(w, response, http.StatusOK)
}

// UpdateConfigurationHandler is a function to update the system configuration.
func UpdateConfigurationHandler(w http.ResponseWriter, r *http.Request) {
	var configReq models.ConfigurationRequest

	err := json.NewDecoder(r.Body).Decode(&configReq)

	if err != nil {
		http.Error(w, "An error occurred when trying to decode configuration "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	// First get the existing configuration to get the ID
	existingConfig, err := repository.GetConfiguration(ctx)
	if err != nil {
		http.Error(w, "An error occurred when trying to get configuration from database "+err.Error(), 400)
		return
	}

	// Update the configuration
	config := models.Configuration{
		ID:                 existingConfig.ID,
		DefaultLocation:    configReq.DefaultLocation,
		WateringHourFactor: configReq.WateringHourFactor,
	}

	updatedConfig, err := repository.UpdateConfiguration(ctx, &config)

	if err != nil {
		http.Error(w, "An error occurred when trying to update configuration in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", updatedConfig)
	ResponseWithJson(w, response, http.StatusOK)
}
