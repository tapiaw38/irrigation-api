package routers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tapiaw38/irrigation-api/libs"
	"github.com/tapiaw38/irrigation-api/models"
	"github.com/tapiaw38/irrigation-api/repository"
)

// CreateProductionsHandler handles the request to get all productions
func CreateProductionHandler(w http.ResponseWriter, r *http.Request) {

	var productions []models.Production

	err := json.NewDecoder(r.Body).Decode(&productions)

	if err != nil {
		http.Error(w, "An error ocurred when trying to enter an productions "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	productions, err = repository.CreateProductions(ctx, productions)

	if err != nil {
		http.Error(w, "An error occurred when trying to create productions in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", productions)
	ResponseWithJson(w, response, http.StatusOK)
}

// GetProductionsHandler handles the request to get all productions
func GetProductionsHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	productions, err := repository.GetProductions(ctx)

	if err != nil {
		http.Error(w, "An error occurred when trying to get productions in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", productions)
	ResponseWithJson(w, response, http.StatusOK)
}

// GetProductionByIDHandler handles the request to get a production by id
func GetProductionByIDHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred when trying to get production by id", 400)
		return
	}

	ctx := r.Context()

	pds, err := repository.GetProductionsByID(ctx, id)

	if err != nil {
		http.Error(w, "An error occurred when trying to get production by id "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", pds)
	ResponseWithJson(w, response, http.StatusOK)
}

// UpdateProductionHandler handles the request to update a production
func UpdateProductionHandler(w http.ResponseWriter, r *http.Request) {

	var production models.Production

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

	pds, err := repository.UpdateProduction(ctx, id, production)

	if err != nil {
		http.Error(w, "An error occurred when trying to update a production in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", pds)
	ResponseWithJson(w, response, http.StatusOK)
}

// DeleteProductionHandler handles the request to delete a production
func DeleteProductionHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", 400)
		return
	}

	ctx := r.Context()

	pds, err := repository.DeleteProduction(ctx, id)

	if err != nil {
		http.Error(w, "An error occurred when trying to delete a production in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", pds)
	ResponseWithJson(w, response, http.StatusOK)
}

func UploadPictureHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	maxSize := int64(1024 * 1024 * 5) // 5MB

	err := r.ParseMultipartForm(maxSize)

	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Image too large. Max Size: %v", maxSize)
		return
	}

	file, fileHeader, err := r.FormFile("picture")

	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Could not get uploaded file")
		return
	}

	defer file.Close()

	// reused if we're uploading many files
	fileName, err := libs.S3.UploadFileToS3(file, fileHeader, id)

	if err != nil {
		fmt.Fprintf(w, "Could not upload file error"+err.Error(), 400)
		return
	}

	ctx := r.Context()

	fileUrl := libs.S3.GenerateUrl(fileName)

	var p models.Production

	p.Picture = fileUrl

	user, err := repository.PartialUpdateProduction(ctx, id, p)

	if err != nil {
		http.Error(w, "An error occurred when trying to update a user in database "+err.Error(), 400)
		return
	}

	response := NewResponse(Message, "ok", user)
	ResponseWithJson(w, response, http.StatusOK)

}
