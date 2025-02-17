package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
	"task.com/orderManagement/firebase"
	"task.com/orderManagement/models"
)

func GetActivitiesPerEmployee(w http.ResponseWriter, r *http.Request) {
	// Obtener el companyId de la URL
	vars := mux.Vars(r)
	employeeId := vars["employeeId"]

	// Realizar una consulta en Firestore para buscar las órdenes del día actual
	// Realizar una consulta en Firestore para buscar las actividades predeterminadas por companyId
	iter := firebase.Client.Collection("Activity").Where("Employee.ID", "==", employeeId).Documents(context.Background())
	var users []models.User
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			http.Error(w, "Error al obtener las actividades por trabajador", http.StatusInternalServerError)
			return
		}
		var user models.User
		doc.DataTo(&user)
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
