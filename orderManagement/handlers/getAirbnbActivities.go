package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
	"task.com/orderManagement/firebase"
	"task.com/orderManagement/models"
)

func GetAirbnbActivities(w http.ResponseWriter, r *http.Request) {
	// Obtener el companyId de la URL
	vars := mux.Vars(r)
	companyId := vars["orderId"]

	// Realizar una consulta en Firestore para buscar las actividades pendientes de Airbnb de la empresa específica
	iter := firebase.Client.Collection("activitiesAirbnbPending").Where("OrderID", "==", companyId).Documents(context.Background())
	var activities []models.Activity
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			http.Error(w, "Error al obtener las actividades", http.StatusInternalServerError)
			return
		}
		var activity models.Activity
		doc.DataTo(&activity)
		activities = append(activities, activity)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activities)
}
