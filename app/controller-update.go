package app

import (
	"crawling/model"
	"crawling/services"
	"crawling/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// [PATCH] /api/update/{id}
func UpdateById(response http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			response.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(response).Encode(map[string]string{
				"message": err.(error).Error(),
			})
		}
	}()

	response.Header().Set("content-type", "application/json")
	id := mux.Vars(request)["id"]
	var updatedData utils.DataModel
	json.NewDecoder(request.Body).Decode(&updatedData)
	services.UpdateById(id, updatedData)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(map[string]string{
		"message": "Updated successfully. ID: " + id,
	})
}

// [PATCH] /api/get?date=...&md5=...&sha1=...&sha256=...&base64=...
func UpdateData(response http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			response.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(response).Encode(map[string]string{
				"message": err.(error).Error(),
			})
		}
	}()

	response.Header().Set("content-type", "application/json")
	var data model.Data
	json.NewDecoder(request.Body).Decode(&data)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(data)
}
