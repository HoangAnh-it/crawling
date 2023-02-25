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

	subRouter.HandleFunc("/push", app.PushData).Methods(http.MethodPost)
	subRouter.HandleFunc("/push/raw-string", app.PushRawString).Methods(http.MethodPost)

	subRouter.HandleFunc("/get/{id}", app.GetDataById).Methods(http.MethodGet)
	subRouter.HandleFunc("/get", app.GetData).Methods(http.MethodGet)

	subRouter.HandleFunc("/update/{id}", app.UpdateById).Methods(http.MethodPatch)

	subRouter.HandleFunc("/delete/{id}", app.DeleteById).Methods(http.MethodDelete)
	subRouter.HandleFunc("/delete", app.DeleteData).Methods(http.MethodDelete)

	return router
}
