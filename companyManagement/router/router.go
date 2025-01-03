package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"task.com/companyManagement/firebase"
	"task.com/companyManagement/handlers"
)

// NewHTTPHandler returns an HTTP handler that handles all the routes.
func NewHTTPHandler() http.Handler {
	firebase.InitFirebase()
	router := mux.NewRouter()
	router.HandleFunc("/createCompany", handlers.CreateCompany).Methods("POST")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Cambia esto para restringir los orígenes permitidos
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}).Handler(router)

	return corsHandler
}
