package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"task.com/orderManagement/firebase"
	"task.com/orderManagement/models"
)

func GetDefaultActivities(w http.ResponseWriter, r *http.Request) {

	// Realizar una consulta en Firestore para buscar por el campo "companyId"
	iter := firebase.Client.Collection("defaultActivity").Documents(context.Background())
	defer iter.Stop()

	var activities []models.DefaultActivity

	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		var user models.DefaultActivity
		if err := doc.DataTo(&user); err != nil {
			continue
		}

		activities = append(activities, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activities)
}
