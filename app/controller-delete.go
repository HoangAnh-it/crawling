package app

import (
	"crawling/services"
	"crawling/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// [DELETE] /api/delete/{id}
func DeleteById(response http.ResponseWriter, request *http.Request) {
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
	services.DeleteById(id)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(map[string]interface{}{
		"message": "Delete successfully. ID: " + id,
	})
}

// [DELETE] /api/delete?date=...&md5=...&sha1=...&sha256=...&base64=...
func DeleteData(response http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			response.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(response).Encode(map[string]string{
				"message": err.(error).Error(),
			})
		}
	}()

	response.Header().Set("content-type", "application/json")
	dataModel := utils.ExtractQueries(request)
	services.DeleteData(dataModel)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(map[string]interface{}{
		"message": "Delete successfully.",
	})
}
