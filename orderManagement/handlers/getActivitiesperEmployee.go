package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
	"task.com/orderManagement/firebase"
	"task.com/orderManagement/models"
)

func GetActivitiesPerEmployee(w http.ResponseWriter, r *http.Request) {
	// Obtener el employeeId de la URL
	vars := mux.Vars(r)
	employeeId := vars["employeeID"]

	// Obtener la fecha de hoy
	today := time.Now().Format("2006-01-02")

	// Realizar una consulta en Firestore para buscar las actividades por employeeId, fecha de hoy y estado "pending"
	iter := firebase.Client.Collection("activity").Where("Employee.ID", "==", employeeId).Where("Date", "==", today).Where("Status", "==", "pending").Documents(context.Background())
	var activities []models.Activity
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			http.Error(w, "Error al obtener las actividades por trabajador", http.StatusInternalServerError)
			return
		}
		var activity models.Activity
		doc.DataTo(&activity)
		activities = append(activities, activity)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activities)
}
