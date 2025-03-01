package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"task.com/orderManagement/firebase"
	"task.com/orderManagement/models"
)

func CreateScheduledOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	startDateStr := vars["startDate"]
	endDateStr := vars["endDate"]

	// Parsear las fechas en el formato aaaa-mm-dd
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		http.Error(w, "Invalid startDate format", http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		http.Error(w, "Invalid endDate format", http.StatusBadRequest)
		return
	}

	var orderTemplate struct {
		CompanyId   string                 `json:"companyId"`
		Location    models.Location        `json:"location"`
		Type        string                 `json:"type"`
		ManagerId   string                 `json:"managerId"`
		ManagerName string                 `json:"managerName"`
		Schedule    []models.ScheduleEntry `json:"schedule"`
	}
	if err := json.NewDecoder(r.Body).Decode(&orderTemplate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Crear órdenes programadas
	var createdOrders []models.Order
	for date := startDate; !date.After(endDate); date = date.AddDate(0, 0, 1) {
		dayOfWeek := date.Weekday().String()
		for _, scheduleEntry := range orderTemplate.Schedule {
			if scheduleEntry.Day == dayOfWeek {
				newOrder := models.Order{
					CompanyId:   orderTemplate.CompanyId,
					Location:    orderTemplate.Location,
					Type:        orderTemplate.Type,
					ManagerId:   orderTemplate.ManagerId,
					ManagerName: orderTemplate.ManagerName,
					Date:        date.Format("2006-01-02"),
					StartDate:   date.Format(time.RFC3339),
					EndDate:     date.Add(8 * time.Hour).Format(time.RFC3339), // Ejemplo: duración de 8 horas
					Status:      "Scheduled",
				}

				// Guardar la nueva orden en Firestore
				docRef := firebase.Client.Collection("order").NewDoc()
				newOrder.ID = docRef.ID
				_, err := docRef.Set(context.Background(), newOrder)
				if err != nil {
					http.Error(w, "Error creating order", http.StatusInternalServerError)
					return
				}

				createdOrders = append(createdOrders, newOrder)
			}
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdOrders)
}
