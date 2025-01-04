package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"task.com/usersManagement/firebase"
	"task.com/usersManagement/models"
)

func GetUsersByCompanyId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	companyId := params["companyId"]

	// Realizar una consulta en Firestore para buscar por el campo "companyId"
	iter := firebase.Client.Collection("user").Where("CompanyId", "==", companyId).Documents(context.Background())
	defer iter.Stop()

	var users []models.User

	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		var user models.User
		if err := doc.DataTo(&user); err != nil {
			continue
		}

		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
