package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tapiaw38/irrigation-api/models/section"
)

type SectionRouter struct {
	Storage section.Storage
}

// CreateSectionHandler is a function to create a section
func (ss *SectionRouter) CreateSectionsHandler(w http.ResponseWriter, r *http.Request) {

	var sections []section.Section

	err := json.NewDecoder(r.Body).Decode(&sections)

	if err != nil {
		http.Error(w, "An error ocurred when trying to enter an section "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	sections, err = ss.Storage.CreateSections(ctx, sections)

	if err != nil {
		http.Error(w, "An error occurred when trying to create section in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", sections)
	ResponseWithJson(w, response, http.StatusOK)
}

// GetSectionsHandler is a function to get all sections
func (ss *SectionRouter) GetSectionsHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	sections, err := ss.Storage.GetSections(ctx)

	if err != nil {
		http.Error(w, "An error occurred when trying to get sections in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", sections)
	ResponseWithJson(w, response, http.StatusOK)
}

// GetSectionByIDHandler is a function to get a section by id
func (ss *SectionRouter) GetSectionByIDHandler(w http.ResponseWriter, r *http.Request) {

	var section section.Section

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&section)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a section "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	section, err = ss.Storage.GetSectionByID(ctx, id)

	if err != nil {
		http.Error(w, "An error occurred when trying to get section in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", section)
	ResponseWithJson(w, response, http.StatusOK)
}

// UpdateSectionHandler is a function to update a section
func (ss *SectionRouter) UpdateSectionHandler(w http.ResponseWriter, r *http.Request) {

	var section section.Section

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&section)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a section "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	section, err = ss.Storage.UpdateSection(ctx, id, section)

	if err != nil {
		http.Error(w, "An error occurred when trying to update section in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", section)
	ResponseWithJson(w, response, http.StatusOK)
}

// DeleteSectionHandler is a function to delete a section
func (ss *SectionRouter) DeleteSectionHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	section, err := ss.Storage.DeleteSection(ctx, id)

	if err != nil {
		http.Error(w, "An error occurred when trying to delete section in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", section)
	ResponseWithJson(w, response, http.StatusOK)
}
