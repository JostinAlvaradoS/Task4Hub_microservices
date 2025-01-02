package router

import (
	"net/http"

	"task.com/usersManagement/handlers" // Assuming handlers package exists

	"github.com/gorilla/mux"
)

// NewHTTPHandler returns an HTTP handler that handles all the routes.
func NewHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{uid}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	return router
}

// This InitRouter function is for local testing
func InitRouter() {
	http.ListenAndServe(":8080", NewHTTPHandler())
}
