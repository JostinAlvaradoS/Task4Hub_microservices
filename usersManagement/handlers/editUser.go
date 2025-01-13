package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
	"task.com/usersManagement/firebase"
)

func EditUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtener la referencia del documento del usuario utilizando el ID del documento
	userRef := firebase.Client.Collection("user").Doc(id)

	// Actualizar los campos del usuario
	_, err := userRef.Update(context.Background(), convertToFirestoreUpdates(updates))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func convertToFirestoreUpdates(updates map[string]interface{}) []firestore.Update {
	var firestoreUpdates []firestore.Update
	for key, value := range updates {
		firestoreUpdates = append(firestoreUpdates, firestore.Update{
			Path:  key,
			Value: value,
		})
	}
	return firestoreUpdates
}
