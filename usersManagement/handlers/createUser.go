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

	// Agregar el usuario a la colección y obtener el ID generado por Firebase
	docRef, _, err := firebase.Client.Collection("user").Add(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Actualizar el campo ID del usuario con el ID generado por Firebase
	user.ID = docRef.ID

	// Guardar el usuario actualizado en Firestore
	_, err = docRef.Set(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
