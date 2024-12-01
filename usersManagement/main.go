package main

import (
    "log"
    "net/http"
    "os"
    "github.com/gorilla/mux"
    "usersManagement/handlers"
	"usersManagement/firebase"
)

func main() {
    // Inicializar Firebase
    firebase.InitFirebase()

    // Configurar el router
    r := mux.NewRouter()
    r.HandleFunc("/", homeHandler).Methods("GET")
    r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
    r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

    // Iniciar el servidor
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Printf("Listening on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Microservicio de gestión de usuarios"))
}