package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"task.com/usersManagement/firebase"
	"task.com/usersManagement/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Crear un nuevo documento en Firestore y obtener su referencia
	docRef := firebase.Client.Collection("user").NewDoc()
	user.ID = docRef.ID

	// Guardar el usuario en Firestore usando el UID proporcionado
	_, err := docRef.Set(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
