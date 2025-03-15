package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
	"task.com/companyManagement/firebase"
	"task.com/companyManagement/models"
)

type CompletedActivity struct {
	EmployeeID   string `json:"employeeId"`
	EmployeeName string `json:"employeeName"`
	Count        int    `json:"count"`
}

type Report struct {
	ActiveUsersCount      int                 `json:"activeUsersCount"`
	CompletedActivities   []CompletedActivity `json:"completedActivities"`
	UncompletedActivities int                 `json:"uncompletedActivities"`
}

func GetReportManager(w http.ResponseWriter, r *http.Request) {
	// Obtener el companyId y la fecha de la URL
	vars := mux.Vars(r)
	companyId := vars["companyId"]
	dateStr := vars["date"]

	// Parsear la fecha en el formato aaaa-mm-dd
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}
	formattedDate := date.Format("2006-01-02")

	// Obtener todos los usuarios con Status == active
	usersIter := firebase.Client.Collection("user").Where("CompanyId", "==", companyId).Where("Status", "==", "active").Documents(context.Background())
	activeUsersCount := 0
	for {
		_, err := usersIter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			http.Error(w, "Error al obtener los usuarios activos", http.StatusInternalServerError)
			return
		}
		activeUsersCount++
	}

	// Obtener todas las actividades completadas de la fecha especificada
	completedActivitiesIter := firebase.Client.Collection("activity").Where("CompanyID", "==", companyId).Where("Status", "==", "finished").Where("Date", "==", formattedDate).Documents(context.Background())
	completedActivitiesMap := make(map[string]*CompletedActivity)
	for {
		doc, err := completedActivitiesIter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			http.Error(w, "Error al obtener las actividades completadas", http.StatusInternalServerError)
			return
		}
		var activity models.Activity
		doc.DataTo(&activity)
		employeeID := activity.Employee.ID
		if _, exists := completedActivitiesMap[employeeID]; !exists {
			completedActivitiesMap[employeeID] = &CompletedActivity{
				EmployeeID:   employeeID,
				EmployeeName: activity.Employee.Name,
				Count:        0,
			}
		}
		completedActivitiesMap[employeeID].Count++
	}

	// Convertir el mapa a una lista
	var completedActivities []CompletedActivity
	for _, ca := range completedActivitiesMap {
		completedActivities = append(completedActivities, *ca)
	}

	// Obtener todas las actividades no completadas de la fecha especificada
	uncompletedActivitiesIter := firebase.Client.Collection("activity").Where("CompanyID", "==", companyId).Where("Status", "==", "pending").Where("Date", "==", formattedDate).Documents(context.Background())
	uncompletedActivitiesCount := 0
	for {
		_, err := uncompletedActivitiesIter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			http.Error(w, "Error al obtener las actividades no completadas", http.StatusInternalServerError)
			return
		}
		uncompletedActivitiesCount++
	}

	// Crear el reporte
	report := Report{
		ActiveUsersCount:      activeUsersCount,
		CompletedActivities:   completedActivities,
		UncompletedActivities: uncompletedActivitiesCount,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
