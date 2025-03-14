package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"task.com/orderManagement/firebase"
	"task.com/orderManagement/models"
)

func CreateAirbnbOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Limpiar los campos de fecha y otros detalles específicos
	order.StartDate = ""
	order.EndDate = ""
	order.Status = "pending"

	docRef := firebase.Client.Collection("order").NewDoc()
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
