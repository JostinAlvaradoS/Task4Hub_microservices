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
	err = assignEmployeesToActivities(order, activities)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Calcular el startTime y endTime de la orden
	orderStartTime, orderEndTime := calculateOrderTimes(activities)
	order.StartDate = orderStartTime.Format(time.RFC3339)
	order.EndDate = orderEndTime.Format(time.RFC3339)

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

func assignEmployeesToActivities(order models.Order, activities []models.Activity) error {
	employeeWorkload := make(map[string]time.Duration)
	for _, employee := range order.Employees {
		employeeWorkload[employee.ID] = 0
	}

	// Parsear el startDate de la orden
	startTime, err := time.Parse(time.RFC3339, order.StartDate)
	if err != nil {
		return fmt.Errorf("invalid start date format: %v", err)
	}

	for _, activity := range activities {
		assigned := false
		for _, employee := range order.Employees {
			if employeeWorkload[employee.ID]+time.Duration(activity.EstimatedTime)*time.Minute <= 8*time.Hour {
				activity.Employee = employee
				activityStartTime := startTime.Add(employeeWorkload[employee.ID])
				activity.StartDate = activityStartTime.Format(time.RFC3339)
				activityEndTime := activityStartTime.Add(time.Duration(activity.EstimatedTime) * time.Minute)
				activity.EndDate = activityEndTime.Format(time.RFC3339)
				employeeWorkload[employee.ID] += time.Duration(activity.EstimatedTime) * time.Minute
				assigned = true
				break
			}
		}
		if !assigned {
			return fmt.Errorf("not enough employees to cover all activities")
		}

		// Actualizar la actividad en Firestore
		docRef := firebase.Client.Collection("activity").Doc(activity.ID)
		_, err := docRef.Set(context.Background(), activity)
		if err != nil {
			return err
		}
	}

	return nil
}

func calculateOrderTimes(activities []models.Activity) (time.Time, time.Time) {
	var earliestStartTime, latestEndTime time.Time
	for i, activity := range activities {
		activityStartTime, err := time.Parse(time.RFC3339, activity.StartDate)
		if err != nil {
			continue
		}
		activityEndTime, err := time.Parse(time.RFC3339, activity.EndDate)
		if err != nil {
			continue
		}
		if i == 0 || activityStartTime.Before(earliestStartTime) {
			earliestStartTime = activityStartTime
		}
		if activityEndTime.After(latestEndTime) {
			latestEndTime = activityEndTime
		}
	}
	// Agregar 15 minutos de gracia
	return earliestStartTime, latestEndTime.Add(15 * time.Minute)
}
