package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"task.com/orderManagement/firebase"
	"task.com/orderManagement/handlers" // Assuming handlers package exists
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
	//Create order
	router.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "https://task4hub.com"}, // Cambia esto para restringir los orígenes permitidos
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}).Handler(router)

	return corsHandler
}
