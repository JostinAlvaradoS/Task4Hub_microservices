package orderManagement

//package main

import (
	//"log"
	"net/http"
	//"os"

	"task.com/orderManagement/router"
)

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
