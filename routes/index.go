package routes

import (
	"crawling/app"
	"github.com/gorilla/mux"
	"net/http"
)

func InitRoutes(SCHEME string, HOST string, PORT string) *mux.Router {
	router := mux.NewRouter()
	subRouter := router.
		Schemes(SCHEME).
		Host(HOST + ":" + PORT).
		PathPrefix("/api").
		Subrouter()

	subRouter.HandleFunc("/push/{date}", app.PushData).Methods(http.MethodPost)
	subRouter.HandleFunc("/get/{date}", app.GetData).Methods(http.MethodGet)

	return router
}
