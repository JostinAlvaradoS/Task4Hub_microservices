package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"task.com/orderManagement/firebase"
	"task.com/orderManagement/models"
)

func CreateActivity(w http.ResponseWriter, r *http.Request) {
	var order models.Activity
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	docRef := firebase.Client.Collection("activity").NewDoc()
	order.ID = docRef.ID

	// Guardar la orden en Firestore
	_, err := docRef.Set(context.Background(), order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
