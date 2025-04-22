package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"task.com/companyManagement/firebase"
	"task.com/companyManagement/models"
)

type EditStockRequest struct {
	CompanyID     string `json:"companyId"`
	CategoryID    string `json:"categoryId,omitempty"`
	SubcategoryID string `json:"subcategoryId,omitempty"`
	ProductID     string `json:"productId,omitempty"`
	NewName       string `json:"newName"`
}

func EditStock(w http.ResponseWriter, r *http.Request) {
	var editRequest EditStockRequest
	if err := json.NewDecoder(r.Body).Decode(&editRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("EditStock request received: %+v", editRequest)

	// Obtener el documento del stock por el ID de la empresa
	iter := firebase.Client.Collection("stock").Where("CompanyID", "==", editRequest.CompanyID).Documents(context.Background())
	var stock models.Stock
	var docRef *firestore.DocumentRef
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			http.Error(w, "Error al buscar el stock", http.StatusInternalServerError)
			return
		}
		docRef = doc.Ref
		if err := doc.DataTo(&stock); err != nil {
			http.Error(w, "Error al decodificar el stock", http.StatusInternalServerError)
			return
		}
		break
	}

	if docRef == nil {
		http.Error(w, "Stock no encontrado", http.StatusNotFound)
		return
	}

	log.Printf("Stock found: %+v", stock)

	// Editar la categoría, subcategoría o producto
	updated := false
	for i, category := range stock.Categories {
		if category.CategoryID == editRequest.CategoryID {
			// Editar categoría
			if editRequest.SubcategoryID == "" && editRequest.ProductID == "" {
				stock.Categories[i].CategoryName = editRequest.NewName
				updated = true
				break
			}
			for j, subcategory := range category.Subcategories {
				if subcategory.SubcategoryID == editRequest.SubcategoryID {
					// Editar subcategoría
					if editRequest.ProductID == "" {
						stock.Categories[i].Subcategories[j].SubcategoryName = editRequest.NewName
						updated = true
						break
					}
					for k, product := range subcategory.Products {
						if product.ProductID == editRequest.ProductID {
							// Editar producto
							stock.Categories[i].Subcategories[j].Products[k].ProductName = editRequest.NewName
							updated = true
							break
						}
					}
				}
			}
		}
	}

	if !updated {
		http.Error(w, "Elemento no encontrado para editar", http.StatusNotFound)
		return
	}

	// Guardar el stock actualizado en Firestore
	_, err := docRef.Set(context.Background(), stock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Stock updated successfully: %+v", stock)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stock)
}
