package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"task.com/companyManagement/firebase"
	"task.com/companyManagement/handlers"
)

// NewHTTPHandler returns an HTTP handler that handles all the routes.
func NewHTTPHandler() http.Handler {
	firebase.InitFirebase()
	router := mux.NewRouter()
	router.HandleFunc("/createCompany", handlers.CreateCompany).Methods("POST")

	return router
}
