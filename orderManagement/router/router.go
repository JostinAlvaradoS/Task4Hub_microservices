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
	//firebase.InitFirebaseLocal()
	router := mux.NewRouter()
	//Create invitations.
	router.HandleFunc("/createInvitation", handlers.InviteUser).Methods("POST")
	//Create order
	router.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")
	//Finish order
	router.HandleFunc("/finishOrder/{id}", handlers.FinishOrder).Methods("POST")
	//Create activity
	router.HandleFunc("/defaultActivities/{companyId}", handlers.GetDefaultActivities).Methods("GET")
	router.HandleFunc("/createActivity", handlers.CreateActivities).Methods("POST")
	router.HandleFunc("/createDefaultActivity", handlers.CreateDefaultActivity).Methods("POST")
	//edit activity
	router.HandleFunc("/editActivity/{id}", handlers.UpdateActivity).Methods("POST")
	//assign employees
	router.HandleFunc("/assignEmployees/{orderID}", handlers.AssignEmployees).Methods("POST")
	// Get available users
	router.HandleFunc("/getAvailableUsers/{companyId}/{startDate}/{endDate}", handlers.GetAvailableUsers).Methods("GET")
	// Get orders by present day
	router.HandleFunc("/getDiariesOrders/{companyId}", handlers.GetDiariesOrders).Methods("GET")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "https://task4hub.com", "https://www.task4hub.com"}, // Cambia esto para restringir los orígenes permitidos
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}).Handler(router)

	return corsHandler
}
