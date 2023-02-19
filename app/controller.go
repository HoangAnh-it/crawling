package app

import (
	"crawling/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// [POST] /api/push?ext
func PushData(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	date := mux.Vars(request)["date"]
	typeFile := request.URL.Query().Get("ext")
	err := services.InsertData(date, typeFile)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(err.Error()))
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode("Push data successfully")
}

// [GET] /api/get?ext
func GetData(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	date := mux.Vars(request)["date"]
	typeFile := request.URL.Query().Get("ext")
	data, err := services.GetData(date, typeFile)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(err.Error()))
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(data)
}
