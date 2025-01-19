package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"task.com/orderManagement/firebase"
	"task.com/orderManagement/models"
)

func CreateActivity(w http.ResponseWriter, r *http.Request) {
	var activity models.Activity
	if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Crear la actividad en Firestore
	docRef := firebase.Client.Collection("activities").NewDoc()
	activity.ID = docRef.ID

	_, err := docRef.Set(context.Background(), activity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Restar el stock necesario de la empresa
	err = subtractStock(activity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(activity)
}

func subtractStock(activity models.Activity) error {
	// Obtener el documento del stock por el ID de la empresa usando un where

	iter := firebase.Client.Collection("stock").Where("CompanyID", "==", activity.CompanyID).Documents(context.Background())
	doc, err := iter.Next()
	if err != nil {
		return err
	}
	if !doc.Exists() {
		return fmt.Errorf("stock no encontrado para la compañía %s", activity.CompanyID)
	}

	var stock models.Stock
	if err := doc.DataTo(&stock); err != nil {
		return err
	}

	// Restar el stock requerido por la actividad del stock de la empresa
	for _, requiredStock := range activity.RequiredStock {
		for i, category := range stock.Categories {
			if category.CategoryID == requiredStock.CategoryID {
				for j, subcategory := range category.Subcategories {
					if subcategory.SubcategoryID == requiredStock.SubcategoryID {
						for k, product := range subcategory.Products {
							if product.ProductID == requiredStock.ProductID {
								if product.CurrentAmount < requiredStock.Quantity {
									return fmt.Errorf("stock insuficiente para el producto %s", product.ProductID)
								}
								stock.Categories[i].Subcategories[j].Products[k].CurrentAmount -= requiredStock.Quantity
							}
						}
					}
				}
			}
		}
	}

	// Actualizar el stock de la empresa en Firestore
	_, err = doc.Ref.Set(context.Background(), stock)
	return err
}
