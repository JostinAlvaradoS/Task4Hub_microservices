package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"task.com/orderManagement/firebase"
	"task.com/orderManagement/models"
)

func CreateDefaultActivity(w http.ResponseWriter, r *http.Request) {
	var defaultActivity models.DefaultActivity
	if err := json.NewDecoder(r.Body).Decode(&defaultActivity); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Crear la actividad predeterminada en Firestore
	docRef := firebase.Client.Collection("defaultActivity").NewDoc()
	defaultActivity.ID = docRef.ID

	_, err := docRef.Set(context.Background(), defaultActivity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(defaultActivity)
}
