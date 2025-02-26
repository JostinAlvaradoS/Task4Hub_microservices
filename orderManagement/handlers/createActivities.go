package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"task.com/orderManagement/firebase"
	"task.com/orderManagement/models"
)

func CreateActivities(w http.ResponseWriter, r *http.Request) {
	var activities []models.Activity
	if err := json.NewDecoder(r.Body).Decode(&activities); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Iniciar una transacción
	ctx := context.Background()
	batch := firebase.Client.Batch()

	for _, activity := range activities {
		if activity.StartDate.IsZero() {
			activity.StartDate = time.Now().UTC()
		}

		// Establecer EndDate en una hora predeterminada si no se proporciona
		if activity.EndDate.IsZero() {
			activity.EndDate = activity.StartDate.Add(1 * time.Hour) // Por ejemplo, 1 hora después de StartDate
		}

		// Crear la actividad en Firestore
		docRef := firebase.Client.Collection("activity").NewDoc()
		activity.ID = docRef.ID

		batch.Set(docRef, activity)
	}

	// Ejecutar la transacción
	_, err := batch.Commit(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(activities)
}
