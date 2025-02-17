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

func GetAirbnbOrders(w http.ResponseWriter, r *http.Request) {
	// Obtener el companyId de la URL
	vars := mux.Vars(r)
	companyId := vars["companyId"]

	// Realizar una consulta en Firestore para buscar las órdenes del día actual y de la empresa específica
	iter := firebase.Client.Collection("order").Where("CompanyId", "==", companyId).Where("Type", "==", "airbnb").Documents(context.Background())
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
