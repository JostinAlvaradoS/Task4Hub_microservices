package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"task.com/companyManagement/firebase"
	"task.com/companyManagement/models"
)

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	companyID := params["companyId"]

	// Obtener el documento del stock por el ID de la empresa
	docRef := firebase.Client.Collection("stock").Doc(companyID)
	doc, err := docRef.Get(context.Background())
	if err != nil {
		http.Error(w, "Stock no encontrado", http.StatusNotFound)
		return
	}

	var stock models.Stock
	if err := doc.DataTo(&stock); err != nil {
		http.Error(w, "Error al decodificar el stock", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stock)
}
