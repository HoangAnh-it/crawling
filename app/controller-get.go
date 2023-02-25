package app

import (
	"crawling/services"
	"crawling/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// [GET] /api/get/{id}
func GetDataById(response http.ResponseWriter, request *http.Request) {
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
	data := services.GetDataById(id)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(data)
}

// [GET] /api/get?date=...&md5=...&sha1=...&sha256=...&base64=...
func GetData(response http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			response.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(response).Encode(map[string]string{
				"message": err.(error).Error(),
			})
		}
	}()

	response.Header().Set("content-type", "application/json")
	query := utils.ExtractQueries(request)
	data := services.GetData(query)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(data)
}
