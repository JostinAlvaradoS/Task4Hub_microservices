package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

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
	assignedEmployees, err := assignEmployeesToActivities(order, activities)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Eliminar los empleados que no se asignaron a ninguna actividad de la orden
	order.Employees = filterAssignedEmployees(order.Employees, assignedEmployees)

	// Actualizar la orden en Firestore
	docRef := firebase.Client.Collection("order").Doc(order.ID)
	_, err = docRef.Set(context.Background(), order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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

func assignEmployeesToActivities(order models.Order, activities []models.Activity) (map[string]bool, error) {
	employeeWorkload := make(map[string]time.Duration)
	assignedEmployees := make(map[string]bool)
	for _, employee := range order.Employees {
		employeeWorkload[employee.ID] = 0
	}

	orderStartTime, err := time.Parse(time.RFC3339, order.StartDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date format: %v", err)
	}
	orderEndTime, err := time.Parse(time.RFC3339, order.EndDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date format: %v", err)
	}

	totalOrderDuration := orderEndTime.Sub(orderStartTime)
	totalActivityDuration := time.Duration(0)
	for _, activity := range activities {
		totalActivityDuration += time.Duration(activity.EstimatedTime) * time.Minute
	}

	if totalActivityDuration > totalOrderDuration*time.Duration(len(order.Employees)) {
		return nil, fmt.Errorf("not enough time to complete all activities with the available employees")
	}

	for _, activity := range activities {
		assigned := false
		for _, employee := range order.Employees {
			if employeeWorkload[employee.ID]+time.Duration(activity.EstimatedTime)*time.Minute <= totalOrderDuration {
				activity.Employee = employee
				employeeWorkload[employee.ID] += time.Duration(activity.EstimatedTime) * time.Minute
				assignedEmployees[employee.ID] = true
				assigned = true
				break
			}
		}
		if !assigned {
			return nil, fmt.Errorf("not enough employees to cover all activities")
		}

		// Actualizar la actividad en Firestore
		docRef := firebase.Client.Collection("activity").Doc(activity.ID)
		_, err := docRef.Set(context.Background(), activity)
		if err != nil {
			return nil, err
		}
	}

	return assignedEmployees, nil
}

func filterAssignedEmployees(employees []models.Employee, assignedEmployees map[string]bool) []models.Employee {
	var filteredEmployees []models.Employee
	for _, employee := range employees {
		if assignedEmployees[employee.ID] {
			filteredEmployees = append(filteredEmployees, employee)
		}
	}
	return filteredEmployees
}
