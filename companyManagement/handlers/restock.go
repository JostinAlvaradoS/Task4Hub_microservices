package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"task.com/companyManagement/firebase"
	"task.com/companyManagement/models"
)

func Restock(w http.ResponseWriter, r *http.Request) {
	var restockRequest models.RestockRequest
	if err := json.NewDecoder(r.Body).Decode(&restockRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtener el documento del stock por el ID de la empresa
	docRef := firebase.Client.Collection("stock").Doc(restockRequest.CompanyID)
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

	// Actualizar la cantidad actual del producto
	updated := false
	for i, category := range stock.Categories {
		if category.CategoryID == restockRequest.CategoryID {
			for j, subcategory := range category.Subcategories {
				if subcategory.SubcategoryID == restockRequest.SubcategoryID {
					for k, product := range subcategory.Products {
						if product.ProductID == restockRequest.ProductID {
							stock.Categories[i].Subcategories[j].Products[k].CurrentAmount += restockRequest.Amount
							updated = true
							break
						}
					}
				}
			}
		}
	}

	if !updated {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}

	// Guardar el stock actualizado en Firestore
	_, err = docRef.Set(context.Background(), stock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stock)
}
