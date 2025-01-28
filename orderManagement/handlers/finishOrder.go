package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
	"task.com/orderManagement/firebase"
	"task.com/orderManagement/models"
)

func FinishOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	orderID := params["id"]

	ctx := context.Background()

	// Obtener la orden
	orderDoc := firebase.Client.Collection("order").Doc(orderID)
	orderSnap, err := orderDoc.Get(ctx)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	var order models.Order
	if err := orderSnap.DataTo(&order); err != nil {
		http.Error(w, "Failed to parse order data", http.StatusInternalServerError)
		return
	}

	// Actualizar el estado de la orden a "finished"
	_, err = orderDoc.Update(ctx, []firestore.Update{
		{Path: "status", Value: "finished"},
	})
	if err != nil {
		http.Error(w, "Failed to update order status", http.StatusInternalServerError)
		return
	}

	// Devolver los productos retornables al stock
	for _, room := range order.Rooms {
		// Obtener las actividades para el cuarto actual
		activities := firebase.Client.Collection("activity").Where("RoomId", "==", room.ID).Documents(ctx)
		for {
			activitySnap, err := activities.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				http.Error(w, "Failed to get activities", http.StatusInternalServerError)
				return
			}

			var activity models.Activity
			if err := activitySnap.DataTo(&activity); err != nil {
				http.Error(w, "Failed to parse activity data", http.StatusInternalServerError)
				return
			}

			for _, stock := range activity.RequiredStock {
				if stock.Returnable {
					// Actualizar el stock de la empresa
					stockDoc := firebase.Client.Collection("company").Doc(order.CompanyId).Collection("stock").Doc(stock.ProductID)
					stockSnap, err := stockDoc.Get(ctx)
					if err != nil {
						http.Error(w, "Failed to get stock item", http.StatusInternalServerError)
						return
					}

					var product models.Product
					if err := stockSnap.DataTo(&product); err != nil {
						http.Error(w, "Failed to parse stock data", http.StatusInternalServerError)
						return
					}

					newAmount := product.CurrentAmount + stock.Quantity
					_, err = stockDoc.Update(ctx, []firestore.Update{
						{Path: "currentAmount", Value: newAmount},
					})
					if err != nil {
						http.Error(w, "Failed to update stock", http.StatusInternalServerError)
						return
					}
				}
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "Order finished and stock updated"})
}
