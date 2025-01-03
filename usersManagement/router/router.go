package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"task.com/usersManagement/firebase"
	"task.com/usersManagement/handlers" // Assuming handlers package exists
)

// NewHTTPHandler returns an HTTP handler that handles all the routes.
func NewHTTPHandler() http.Handler {
	firebase.InitFirebase()
	// firebase.InitFirebaseLocal()
	router := mux.NewRouter()
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{uid}", handlers.GetUser).Methods("GET")
	//Create invitations.
	router.HandleFunc("/createInvitation", handlers.InviteUser).Methods("POST")

	return router
}
