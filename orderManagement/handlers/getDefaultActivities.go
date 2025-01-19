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

func GetDefaultActivities(w http.ResponseWriter, r *http.Request) {
	// Obtener el companyId de la URL
	vars := mux.Vars(r)
	companyId := vars["companyId"]

	// Realizar una consulta en Firestore para buscar las actividades predeterminadas por companyId
	iter := firebase.Client.Collection("defaultActivity").Where("CompanyID", "==", companyId).Documents(context.Background())
	var activities []models.DefaultActivity
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			http.Error(w, "Error al obtener las actividades predeterminadas", http.StatusInternalServerError)
			return
		}
		var activity models.DefaultActivity
		doc.DataTo(&activity)
		activities = append(activities, activity)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activities)
}
