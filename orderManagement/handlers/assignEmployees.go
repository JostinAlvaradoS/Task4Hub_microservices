package handlers

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
	"task.com/orderManagement/firebase"
	"task.com/orderManagement/models"
)

func AssignEmployees(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["orderID"]

	// Obtener la orden por ID
	order, err := getOrderByID(orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Obtener todas las actividades relacionadas con la orden
	activities := getActivitiesByOrderID(orderID)

	// Asignar empleados a las actividades de manera equitativa
	assignEmployeesToActivities(order.Employees, activities)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Employees assigned to activities successfully"))
}

func getOrderByID(orderID string) (models.Order, error) {
	docRef := firebase.Client.Collection("order").Doc(orderID)
	doc, err := docRef.Get(context.Background())
	if err != nil {
		return models.Order{}, err
	}
	var order models.Order
	doc.DataTo(&order)
	return order, nil
}

func getActivitiesByOrderID(orderID string) []models.Activity {
	iter := firebase.Client.Collection("activity").Where("OrderID", "==", orderID).Documents(context.Background())
	var activities []models.Activity
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			continue
		}
		var activity models.Activity
		doc.DataTo(&activity)
		activities = append(activities, activity)
	}
	return activities
}

func assignEmployeesToActivities(employees []models.Employee, activities []models.Activity) {
	employeeIndex := 0
	for i := range activities {
		activities[i].Employee = employees[employeeIndex]
		employeeIndex = (employeeIndex + 1) % len(employees)

		// Actualizar la actividad en Firestore
		docRef := firebase.Client.Collection("activity").Doc(activities[i].ID)
		_, err := docRef.Set(context.Background(), activities[i])
		if err != nil {
			continue
		}
	}
}
