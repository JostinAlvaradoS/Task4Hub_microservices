package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"google.golang.org/api/iterator"
	"task.com/companyManagement/firebase"
	"task.com/companyManagement/models"
)

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	var company models.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Crear la compañía en Firestore
	docRef := firebase.Client.Collection("company").NewDoc()
	company.ID = docRef.ID

	_, err := docRef.Set(context.Background(), company)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Crear stock y actividades predeterminadas para la compañía
	if err := createCompanyStock(company.ID); err != nil {
		http.Error(w, "Error al crear el stock de la compañía: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := createCompanyDefaultActivities(company.ID); err != nil {
		http.Error(w, "Error al crear las actividades predeterminadas de la compañía: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(company)
}

func createCompanyStock(companyID string) error {
	// Obtener el stock base desde Firestore
	iter := firebase.Client.Collection("stockBase").Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Error al obtener stockBase: %v", err)
			return err
		}

		var stockBase models.Stock
		doc.DataTo(&stockBase)

		// Asignar el companyID al stock
		stockBase.CompanyID = companyID

		// Guardar el stock en la colección "stock"
		docRef := firebase.Client.Collection("stock").NewDoc()
		_, err = docRef.Set(context.Background(), stockBase)
		if err != nil {
			log.Printf("Error al guardar el stock de la compañía: %v", err)
			return err
		}
	}

	log.Println("Stock creado correctamente para la compañía:", companyID)
	return nil
}

func createCompanyDefaultActivities(companyID string) error {
	// Obtener las actividades base desde Firestore
	iter := firebase.Client.Collection("activitiesBase").Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Error al obtener activitiesBase: %v", err)
			return err
		}

		var activityBase models.DefaultActivity
		doc.DataTo(&activityBase)

		// Asignar el companyID a la actividad
		activityBase.CompanyID = companyID

		// Guardar la actividad en la colección "defaultActivities"
		docRef := firebase.Client.Collection("defaultActivity").NewDoc()
		_, err = docRef.Set(context.Background(), activityBase)
		if err != nil {
			log.Printf("Error al guardar la actividad predeterminada de la compañía: %v", err)
			return err
		}
	}

	log.Println("Actividades predeterminadas creadas correctamente para la compañía:", companyID)
	return nil
}
