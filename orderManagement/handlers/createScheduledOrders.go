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
		Rooms       []models.Rooms         `json:"rooms"`
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
				startTime, err := time.Parse("15:04", scheduleEntry.StartTime)
				if err != nil {
					http.Error(w, "Invalid start time format", http.StatusBadRequest)
					return
				}
				endTime, err := time.Parse("15:04", scheduleEntry.EndTime)
				if err != nil {
					http.Error(w, "Invalid end time format", http.StatusBadRequest)
					return
				}

				startDateTime := time.Date(date.Year(), date.Month(), date.Day(), startTime.Hour(), startTime.Minute(), 0, 0, time.UTC)
				endDateTime := time.Date(date.Year(), date.Month(), date.Day(), endTime.Hour(), endTime.Minute(), 0, 0, time.UTC)

				newOrder := models.Order{
					CompanyId:   orderTemplate.CompanyId,
					Location:    orderTemplate.Location,
					Type:        orderTemplate.Type,
					ManagerId:   orderTemplate.ManagerId,
					ManagerName: orderTemplate.ManagerName,
					Rooms:       orderTemplate.Rooms,
					Schedule:    orderTemplate.Schedule,
					Date:        date.Format("2006-01-02"),
					StartDate:   startDateTime.Format(time.RFC3339),
					EndDate:     endDateTime.Format(time.RFC3339),
					Status:      "Pending",
					Scheduled:   true,
				}

				// Guardar la nueva orden en Firestore
				docRef := firebase.Client.Collection("order").NewDoc()
				newOrder.ID = docRef.ID
				_, err = docRef.Set(context.Background(), newOrder)
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
