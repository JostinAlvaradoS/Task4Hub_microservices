package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/api/iterator"
	"task.com/orderManagement/firebase"
	"task.com/orderManagement/models"
)

func CreateActivities(w http.ResponseWriter, r *http.Request) {
	var activities []models.Activity
	if err := json.NewDecoder(r.Body).Decode(&activities); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Iniciar una transacción
	ctx := context.Background()
	batch := firebase.Client.Batch()

	for _, activity := range activities {
		// Crear la actividad en Firestore
		docRef := firebase.Client.Collection("activity").NewDoc()
		activity.ID = docRef.ID

		batch.Set(docRef, activity)

		// Restar el stock necesario de la empresa
		err := subtractStock(ctx, activity)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Ejecutar la transacción
	_, err := batch.Commit(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(activities)
}

func subtractStock(ctx context.Context, activity models.Activity) error {
	// Obtener el documento del stock por el ID de la empresa usando un where
	iter := firebase.Client.Collection("stock").Where("CompanyID", "==", activity.CompanyID).Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
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
										return fmt.Errorf("insufficient stock for product %s", product.ProductID)
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
		_, err = doc.Ref.Set(ctx, stock)
		if err != nil {
			return err
		}
	}

	return nil
}
