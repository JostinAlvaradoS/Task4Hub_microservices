package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
	"task.com/companyManagement/firebase"
	"task.com/companyManagement/models"
)

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	companyID := params["companyId"]

	// Realizar una consulta en Firestore para buscar el stock por CompanyID
	iter := firebase.Client.Collection("stock").Where("CompanyID", "==", companyID).Documents(context.Background())
	var stock models.Stock
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			http.Error(w, "Error al obtener el stock", http.StatusInternalServerError)
			return
		}

		if err := doc.DataTo(&stock); err != nil {
			http.Error(w, "Error al decodificar el stock", http.StatusInternalServerError)
			return
		}
		// Solo necesitamos el primer documento que coincida
		break
	}

	// Si no se encuentra un documento, devolver un error
	if stock.CompanyID == "" {
		http.Error(w, "Stock no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stock)
}
