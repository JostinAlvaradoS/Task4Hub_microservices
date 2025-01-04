package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"task.com/companyManagement/firebase"
	"task.com/companyManagement/models"
)

func AddStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock
	if err := json.NewDecoder(r.Body).Decode(&stock); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Guardar el stock en Firestore
	_, err := firebase.Client.Collection("stock").Doc(stock.CompanyID).Set(context.Background(), stock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(stock)
}
