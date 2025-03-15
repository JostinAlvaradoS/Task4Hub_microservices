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

func GetScheduledOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]

	iter := firebase.Client.Collection("order").Where("CompanyId", "==", companyId).Where("Scheduled", "==", true).Documents(context.Background())
	var orders []models.Order
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			http.Error(w, "Error fetching orders", http.StatusInternalServerError)
			return
		}
		var order models.Order
		doc.DataTo(&order)
		orders = append(orders, order)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}
