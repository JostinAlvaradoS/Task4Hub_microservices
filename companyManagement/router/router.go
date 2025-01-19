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
	// firebase.InitFirebaseLocal()
	router := mux.NewRouter()
	router.HandleFunc("/createCompany", handlers.CreateCompany).Methods("POST")
	router.HandleFunc("/companiesInfo", handlers.GetCompanyManagersAndEmployees).Methods("GET")
	//manmagement stock
	router.HandleFunc("/addStock", handlers.AddStock).Methods("POST")
	router.HandleFunc("/getStock/{companyId}", handlers.GetStock).Methods("GET")
	router.HandleFunc("/restock", handlers.Restock).Methods("POST")
	//get report
	router.HandleFunc("/report/{companyId}", handlers.GetReportManager).Methods("GET")
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "https://task4hub.com", "https://www.task4hub.com"}, // Cambia esto para restringir los orígenes permitidos
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}).Handler(router)

	return corsHandler
}
