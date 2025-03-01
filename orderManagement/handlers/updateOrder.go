package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"task.com/orderManagement/firebase"
	"task.com/orderManagement/models"
)

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Verificar que el ID de la orden esté presente
	if order.ID == "" {
		http.Error(w, "Order ID is required", http.StatusBadRequest)
		return
	}

	// Obtener el documento de la orden en Firestore
	docRef := firebase.Client.Collection("order").Doc(order.ID)
	doc, err := docRef.Get(context.Background())
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	var existingOrder models.Order
	if err := doc.DataTo(&existingOrder); err != nil {
		http.Error(w, "Error decoding order", http.StatusInternalServerError)
		return
	}

	// Actualizar los campos de la orden solo si se envían en la solicitud
	if order.Date != "" {
		existingOrder.Date = order.Date
	}
	if order.Type != "" {
		existingOrder.Type = order.Type
	}
	if order.Location.Address != "" {
		existingOrder.Location = order.Location
	}
	if order.CompanyId != "" {
		existingOrder.CompanyId = order.CompanyId
	}
	if order.ManagerId != "" {
		existingOrder.ManagerId = order.ManagerId
	}
	if order.ManagerName != "" {
		existingOrder.ManagerName = order.ManagerName
	}
	if len(order.Rooms) > 0 {
		existingOrder.Rooms = order.Rooms
	}
	if len(order.Employees) > 0 {
		existingOrder.Employees = order.Employees
	}
	if order.StartDate != "" {
		existingOrder.StartDate = order.StartDate
	}
	if order.EndDate != "" {
		existingOrder.EndDate = order.EndDate
	}
	if order.Status != "" {
		existingOrder.Status = order.Status
	}
	if len(order.Schedule) > 0 {
		existingOrder.Schedule = order.Schedule
	}

	// Guardar la orden actualizada en Firestore
	_, err = docRef.Set(context.Background(), existingOrder)
	if err != nil {
		http.Error(w, "Error updating order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingOrder)
}
