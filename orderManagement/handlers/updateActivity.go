package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"task.com/orderManagement/firebase"
	"task.com/orderManagement/models"
)

func UpdateActivity(w http.ResponseWriter, r *http.Request) {
	var activity models.Activity
	if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Verificar que el ID de la actividad esté presente
	if activity.ID == "" {
		http.Error(w, "Activity ID is required", http.StatusBadRequest)
		return
	}

	// Obtener el documento de la actividad en Firestore
	docRef := firebase.Client.Collection("activitiy").Doc(activity.ID)
	doc, err := docRef.Get(context.Background())
	if err != nil {
		http.Error(w, "Activity not found", http.StatusNotFound)
		return
	}

	var existingActivity models.Activity
	if err := doc.DataTo(&existingActivity); err != nil {
		http.Error(w, "Error decoding activity", http.StatusInternalServerError)
		return
	}

	// Actualizar los campos de la actividad solo si se envían en la solicitud
	if activity.Name != "" {
		existingActivity.Name = activity.Name
	}
	if activity.Description != "" {
		existingActivity.Description = activity.Description
	}
	if activity.OrderID != "" {
		existingActivity.OrderID = activity.OrderID
	}
	if activity.RoomId != "" {
		existingActivity.RoomId = activity.RoomId
	}
	if activity.Employee.ID != "" {
		existingActivity.Employee = activity.Employee
	}
	if activity.ActivityType != "" {
		existingActivity.ActivityType = activity.ActivityType
	}
	if !activity.StartDate.IsZero() {
		existingActivity.StartDate = activity.StartDate
	}
	if !activity.EndDate.IsZero() {
		existingActivity.EndDate = activity.EndDate
	}
	if activity.Status != "" {
		existingActivity.Status = activity.Status
	}
	if len(activity.RequiredStock) > 0 {
		existingActivity.RequiredStock = activity.RequiredStock
	}

	// Guardar la actividad actualizada en Firestore
	_, err = docRef.Set(context.Background(), existingActivity)
	if err != nil {
		http.Error(w, "Error updating activity", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingActivity)
}
