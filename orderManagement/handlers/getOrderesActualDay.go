package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"google.golang.org/api/iterator"
	"task.com/orderManagement/firebase"
	"task.com/orderManagement/models"
)

func GetOrdersActualDay(w http.ResponseWriter, r *http.Request) {
	// Obtener la fecha actual
	currentDate := time.Now().UTC().Format("2006-01-02")

	// Realizar una consulta en Firestore para buscar las órdenes del día actual
	iter := firebase.Client.Collection("order").Where("Date", "==", currentDate).Documents(context.Background())
	var orders []models.Order
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			http.Error(w, "Error al obtener las órdenes", http.StatusInternalServerError)
			return
		}
		var order models.Order
		doc.DataTo(&order)
		orders = append(orders, order)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
