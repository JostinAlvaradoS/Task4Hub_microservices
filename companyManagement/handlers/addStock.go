package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"task.com/companyManagement/firebase"
	"task.com/companyManagement/models"
)

func AddStock(w http.ResponseWriter, r *http.Request) {
	var newStock models.Stock
	if err := json.NewDecoder(r.Body).Decode(&newStock); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Referencia al documento
	docRef := firebase.Client.Collection("stock").Doc(newStock.CompanyID)

	// Ejecutar la transacción
	err := firebase.Client.RunTransaction(context.Background(), func(ctx context.Context, tx *firestore.Transaction) error {
		doc, err := tx.Get(docRef)
		if err != nil {
			if status.Code(err) == codes.NotFound {
				// Si el documento no existe, crear uno nuevo
				return tx.Set(docRef, newStock)
			}
			return err
		}

		// Documento existente
		var existingStock models.Stock
		if err := doc.DataTo(&existingStock); err != nil {
			return err
		}

		// Actualizar el documento en memoria
		existingStock = mergeStock(existingStock, newStock)

		// Escribir el documento modificado de vuelta a Firestore
		return tx.Set(docRef, existingStock)
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newStock)
}

// mergeStock combina el stock existente con el nuevo stock
func mergeStock(existing, new models.Stock) models.Stock {
	for _, newCategory := range new.Categories {
		foundCategory := false
		for i, existingCategory := range existing.Categories {
			if existingCategory.CategoryID == newCategory.CategoryID {
				foundCategory = true
				// Actualizar subcategorías si existen
				if len(newCategory.Subcategories) > 0 {
					existing.Categories[i] = mergeCategory(existingCategory, newCategory)
				}
				break
			}
		}
		if !foundCategory {
			// Agregar nueva categoría sin subcategorías si no existen
			if len(newCategory.Subcategories) == 0 {
				newCategory.Subcategories = nil
			}
			existing.Categories = append(existing.Categories, newCategory)
		}
	}
	return existing
}

// mergeCategory combina la categoría existente con la nueva categoría
func mergeCategory(existing, new models.Category) models.Category {
	for _, newSubcategory := range new.Subcategories {
		foundSubcategory := false
		for i, existingSubcategory := range existing.Subcategories {
			if existingSubcategory.SubcategoryID == newSubcategory.SubcategoryID {
				foundSubcategory = true
				// Actualizar productos si existen
				if len(newSubcategory.Products) > 0 {
					existing.Subcategories[i] = mergeSubcategory(existingSubcategory, newSubcategory)
				}
				break
			}
		}
		if !foundSubcategory {
			// Agregar nueva subcategoría sin productos si no existen
			if len(newSubcategory.Products) == 0 {
				newSubcategory.Products = nil
			}
			existing.Subcategories = append(existing.Subcategories, newSubcategory)
		}
	}
	return existing
}

// mergeSubcategory combina la subcategoría existente con la nueva subcategoría
func mergeSubcategory(existing, new models.Subcategory) models.Subcategory {
	for _, newProduct := range new.Products {
		foundProduct := false
		for i, existingProduct := range existing.Products {
			if existingProduct.ProductID == newProduct.ProductID {
				foundProduct = true
				// Actualizar cantidad del producto
				existing.Products[i].CurrentAmount += newProduct.CurrentAmount
				break
			}
		}
		if !foundProduct {
			// Agregar nuevo producto
			existing.Products = append(existing.Products, newProduct)
		}
	}
	return existing
}
