package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"task.com/usersManagement/firebase"
	"task.com/usersManagement/handlers"
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
	//receptar todos los usuarios de una empresa
	router.HandleFunc("/users/company/{companyId}", handlers.GetUsersByCompanyId).Methods("GET")
	//editar usuario
	router.HandleFunc("/editUser/{uid}", handlers.EditUser).Methods("POST")
	//verificar invitacion
	router.HandleFunc("/verifyInvitation/{id}", handlers.VerifyInvitation).Methods("GET")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "https://task4hub.com"}, // Cambia esto para restringir los orígenes permitidos
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}).Handler(router)

	return corsHandler
}
