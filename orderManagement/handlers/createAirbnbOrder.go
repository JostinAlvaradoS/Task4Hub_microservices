package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"task.com/orderManagement/firebase"
	"task.com/orderManagement/models"
)

type CreateAirbnbOrderRequest struct {
	Order      models.Order      `json:"order"`
	Activities []models.Activity `json:"activities"`
}

func CreateAirbnbOrder(w http.ResponseWriter, r *http.Request) {
	var request CreateAirbnbOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := request.Order
	order.Status = "pending"

	docRef := firebase.Client.Collection("orderAirbnbPending").NewDoc()
	order.ID = docRef.ID

	// Guardar la orden en Firestore
	_, err := docRef.Set(context.Background(), order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Crear actividades pendientes para la orden
	for _, activity := range request.Activities {
		activity.OrderID = order.ID
		activity.CompanyID = order.CompanyId
		activity.Status = "pending"
		activity.ActivityType = "airbnb"

		activityDocRef := firebase.Client.Collection("activitiesAirbnbPending").NewDoc()
		activity.ID = activityDocRef.ID

		_, err := activityDocRef.Set(context.Background(), activity)
		if err != nil {
			http.Error(w, "Error creating activity", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
