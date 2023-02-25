package app

import (
	"crawling/services"
	"crawling/utils"
	"encoding/json"
	"net/http"
)

// [POST] /api/push/raw-string
func PushRawString(response http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			response.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(response).Encode(map[string]string{
				"message": err.(error).Error(),
			})
		}
	}()

	response.Header().Set("content-type", "application/json")

	var data utils.RawStringDataModel
	json.NewDecoder(request.Body).Decode(&data)
	id := services.PushRawString(data)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(map[string]string{
		"message": "Push successfully.",
		"id":      id,
	})
}

// [POST] /api/push
func PushData(response http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			response.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(response).Encode(map[string]string{
				"message": err.(error).Error(),
			})
		}
	}()

	response.Header().Set("content-type", "application/json")
	var data utils.DataModel
	json.NewDecoder(request.Body).Decode(&data)
	ids := services.PushData(data)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(map[string]string{
		"message": "Push successfully.",
		"ids":     ids,
	})
}
