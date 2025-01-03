package companyManagement

// package main

import (
	// "log"
	// "os"
	"net/http"

	"task.com/companyManagement/router"
)

// CloudFunctionEntryPoint es el punto de entrada para Google Cloud Functions
func CloudFunctionEntryPoint(w http.ResponseWriter, r *http.Request) {
	handler := router.NewHTTPHandler()
	handler.ServeHTTP(w, r)
}

// func main() {
// 	handler := router.NewHTTPHandler()
// 	// Obtener el puerto desde las variables de entorno (por defecto 8080)
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080"
// 	}
// 	// Iniciar el servidor HTTP
// 	log.Printf("Servidor escuchando en el puerto %s", port)
// 	if err := http.ListenAndServe(":"+port, handler); err != nil {
// 		log.Fatalf("Error al iniciar el servidor: %v", err)
// 	}
// }
