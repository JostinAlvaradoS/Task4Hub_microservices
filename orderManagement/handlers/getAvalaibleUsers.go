package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
	"task.com/orderManagement/firebase"
	"task.com/orderManagement/models"
)

func GetAvailableUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]
	startDateStr := vars["startDate"]
	endDateStr := vars["endDate"]

	startDate, err := time.Parse(time.RFC3339, startDateStr)
	if err != nil {
		http.Error(w, "Invalid startDate format", http.StatusBadRequest)
		fmt.Println("Invalid startDate format")
		return
	}

	endDate, err := time.Parse(time.RFC3339, endDateStr)
	if err != nil {
		http.Error(w, "Invalid endDate format", http.StatusBadRequest)
		fmt.Println("Invalid endDate format")
		return
	}

	// Obtener todos los usuarios de la compañía
	users := getUsersByCompanyId(companyId)

	// Filtrar usuarios disponibles
	availableUsers := filterAvailableUsers(users, startDate, endDate)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(availableUsers)
}

func getUsersByCompanyId(companyId string) []models.User {
	iter := firebase.Client.Collection("user").Where("CompanyId", "==", companyId).Documents(context.Background())
	var users []models.User
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Printf("Error fetching user: %v\n", err)
			continue
		}
		var user models.User
		doc.DataTo(&user)
		users = append(users, user)
	}
	return users
}

func filterAvailableUsers(users []models.User, startDate, endDate time.Time) []models.User {
	var availableUsers []models.User
	for _, user := range users {
		if isUserAvailable(user, startDate, endDate) {
			availableUsers = append(availableUsers, user)
		}
	}
	return availableUsers
}

func isUserAvailable(user models.User, startDate, endDate time.Time) bool {
	// Obtener todas las actividades del usuario
	iter := firebase.Client.Collection("activity").Where("Employee.ID", "==", user.ID).Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Printf("Error fetching activity for user %s: %v\n", user.ID, err)
			continue
		}
		var activity models.Activity
		doc.DataTo(&activity)

		// Verificar si la actividad se solapa con el intervalo de tiempo especificado
		if activity.StartDate.Before(endDate.Add(-1*time.Hour)) && activity.EndDate.After(startDate.Add(1*time.Hour)) {
			return false
		}
	}

	// Verificar la disponibilidad según el horario de trabajo del usuario
	isAvailable := isWithinWorkHours(user.Schedule, startDate, endDate)
	return isAvailable
}

func isWithinWorkHours(schedule models.Schedule, startDate, endDate time.Time) bool {
	dayOfWeek := startDate.Weekday().String()
	for _, workDay := range schedule.WorkDays {
		if workDay.Day == dayOfWeek {
			for _, scheduleDay := range workDay.Schedule {
				startTime, _ := time.Parse("15:04", scheduleDay.StartTime)
				endTime, _ := time.Parse("15:04", scheduleDay.EndTime)

				workStart := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), startTime.Hour(), startTime.Minute(), 0, 0, startDate.Location())
				workEnd := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), endTime.Hour(), endTime.Minute(), 0, 0, startDate.Location())

				// Verificar si el intervalo de tiempo especificado está completamente dentro de algún rango de horario de trabajo
				if (startDate.After(workStart) || startDate.Equal(workStart)) && (endDate.Before(workEnd) || endDate.Equal(workEnd)) {
					return true
				}
			}
		}
	}
	return false
}
