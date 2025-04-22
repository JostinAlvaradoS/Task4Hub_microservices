package handlers

import (
	"context"
	"log"
	"net/http"

	"task.com/companyManagement/firebase"
	"task.com/companyManagement/models"
)

func InsertBaseData(w http.ResponseWriter, r *http.Request) {
	// Insertar datos base en stockBase
	if err := insertStockBase(); err != nil {
		http.Error(w, "Error al insertar stockBase: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Insertar datos base en activitiesBase
	if err := insertActivitiesBase(); err != nil {
		http.Error(w, "Error al insertar activitiesBase: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Datos base insertados correctamente."))
}

func insertStockBase() error {
	stockBase := models.Stock{
		Categories: []models.Category{
			{
				CategoryID:   "cat001",
				CategoryName: "Cleaning Supplies",
				Subcategories: []models.Subcategory{
					{
						SubcategoryID:   "subcat001",
						SubcategoryName: "Detergents",
						Products: []models.Product{
							{
								ProductID:     "prod001",
								ProductName:   "Floor Cleaner",
								CurrentAmount: 100,
								TargetAmount:  200,
							},
							{
								ProductID:     "prod002",
								ProductName:   "Glass Cleaner",
								CurrentAmount: 50,
								TargetAmount:  100,
							},
						},
					},
					{
						SubcategoryID:   "subcat002",
						SubcategoryName: "Disinfectants",
						Products: []models.Product{
							{
								ProductID:     "prod003",
								ProductName:   "Bathroom Disinfectant",
								CurrentAmount: 80,
								TargetAmount:  150,
							},
							{
								ProductID:     "prod004",
								ProductName:   "Surface Disinfectant",
								CurrentAmount: 60,
								TargetAmount:  120,
							},
						},
					},
				},
			},
			{
				CategoryID:   "cat002",
				CategoryName: "Bedroom Supplies",
				Subcategories: []models.Subcategory{
					{
						SubcategoryID:   "subcat003",
						SubcategoryName: "Bedding",
						Products: []models.Product{
							{
								ProductID:     "prod005",
								ProductName:   "Bed Sheets",
								CurrentAmount: 200,
								TargetAmount:  300,
							},
							{
								ProductID:     "prod006",
								ProductName:   "Pillow Covers",
								CurrentAmount: 150,
								TargetAmount:  250,
							},
						},
					},
				},
			},
		},
	}

	// Guardar en Firestore
	docRef := firebase.Client.Collection("stockBase").NewDoc()
	_, err := docRef.Set(context.Background(), stockBase)
	if err != nil {
		log.Printf("Error al insertar stockBase: %v", err)
		return err
	}
	log.Println("Stock base insertado correctamente.")
	return nil
}

func insertActivitiesBase() error {
	activitiesBase := []models.DefaultActivity{
		// Bathroom Activities
		{
			ID:           "act001",
			Name:         "Clean Bathroom Floor",
			Description:  "Thoroughly clean the bathroom floor.",
			ActivityType: "Cleaning",
			RoomType:     "Bathroom",
			RequiredStock: []models.RequiredStock{
				{
					CategoryID:      "cat001",
					CategoryName:    "Cleaning Supplies",
					SubcategoryID:   "subcat001",
					SubcategoryName: "Detergents",
					ProductID:       "prod001",
					ProductName:     "Floor Cleaner",
					Quantity:        1,
					Returnable:      false,
				},
			},
			EstimatedTime: 30,
		},
		{
			ID:           "act002",
			Name:         "Disinfect Bathroom Surfaces",
			Description:  "Disinfect all bathroom surfaces.",
			ActivityType: "Disinfection",
			RoomType:     "Bathroom",
			RequiredStock: []models.RequiredStock{
				{
					CategoryID:      "cat001",
					CategoryName:    "Cleaning Supplies",
					SubcategoryID:   "subcat002",
					SubcategoryName: "Disinfectants",
					ProductID:       "prod003",
					ProductName:     "Bathroom Disinfectant",
					Quantity:        1,
					Returnable:      false,
				},
			},
			EstimatedTime: 20,
		},
		// Bedroom Activities
		{
			ID:           "act003",
			Name:         "Change Bed Sheets",
			Description:  "Replace the bed sheets with clean ones.",
			ActivityType: "Maintenance",
			RoomType:     "Bedroom",
			RequiredStock: []models.RequiredStock{
				{
					CategoryID:      "cat002",
					CategoryName:    "Bedroom Supplies",
					SubcategoryID:   "subcat003",
					SubcategoryName: "Bedding",
					ProductID:       "prod005",
					ProductName:     "Bed Sheets",
					Quantity:        1,
					Returnable:      false,
				},
			},
			EstimatedTime: 15,
		},
		{
			ID:            "act004",
			Name:          "Vacuum Bedroom Floor",
			Description:   "Vacuum the entire bedroom floor.",
			ActivityType:  "Cleaning",
			RoomType:      "Bedroom",
			RequiredStock: []models.RequiredStock{},
			EstimatedTime: 25,
		},
		// Living Room Activities
		{
			ID:            "act005",
			Name:          "Dust Furniture",
			Description:   "Dust all furniture in the living room.",
			ActivityType:  "Cleaning",
			RoomType:      "Living Room",
			RequiredStock: []models.RequiredStock{},
			EstimatedTime: 20,
		},
		{
			ID:           "act006",
			Name:         "Mop Living Room Floor",
			Description:  "Mop the floor in the living room.",
			ActivityType: "Cleaning",
			RoomType:     "Living Room",
			RequiredStock: []models.RequiredStock{
				{
					CategoryID:      "cat001",
					CategoryName:    "Cleaning Supplies",
					SubcategoryID:   "subcat001",
					SubcategoryName: "Detergents",
					ProductID:       "prod001",
					ProductName:     "Floor Cleaner",
					Quantity:        1,
					Returnable:      false,
				},
			},
			EstimatedTime: 30,
		},
	}

	// Guardar cada actividad en Firestore
	for _, activity := range activitiesBase {
		docRef := firebase.Client.Collection("activitiesBase").NewDoc()
		_, err := docRef.Set(context.Background(), activity)
		if err != nil {
			log.Printf("Error al insertar activityBase: %v", err)
			return err
		}
	}
	log.Println("Activities base insertadas correctamente.")
	return nil
}
