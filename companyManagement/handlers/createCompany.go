package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"task.com/companyManagement/firebase"
	"task.com/companyManagement/models"
)

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	var company models.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	docRef := firebase.Client.Collection("company").NewDoc()
	company.ID = docRef.ID

	_, err := docRef.Set(context.Background(), company)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(company)
}
