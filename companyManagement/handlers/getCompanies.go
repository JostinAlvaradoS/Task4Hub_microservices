package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"task.com/companyManagement/firebase"
	"task.com/companyManagement/models"
)

func GetCompanyManagersAndEmployees(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Obtener todos los usuarios con rol de "supermanager"
	supermanagerIter := firebase.Client.Collection("user").Where("Role", "==", "supermanager").Documents(ctx)
	defer supermanagerIter.Stop()

	companyInfoMap := make(map[string]*models.CompanyInfo)

	for {
		doc, err := supermanagerIter.Next()
		if err != nil {
			break
		}

		var user models.User
		if err := doc.DataTo(&user); err != nil {
			continue
		}

		companyInfoMap[user.CompanyId] = &models.CompanyInfo{
			Company:      user.CompanyName,
			SuperManager: user.Name,
			Employees:    0, // Inicializar el contador de empleados
		}
	}

	// Obtener la información de las empresas
	companyIter := firebase.Client.Collection("company").Documents(ctx)
	defer companyIter.Stop()

	for {
		doc, err := companyIter.Next()
		if err != nil {
			break
		}

		var company models.Company
		if err := doc.DataTo(&company); err != nil {
			continue
		}

		if companyInfo, exists := companyInfoMap[company.ID]; exists {
			companyInfo.Employees = company.EmployeeCount
		}
	}

	// Convertir el mapa a una lista
	var companyInfoList []models.CompanyInfo
	for _, info := range companyInfoMap {
		companyInfoList = append(companyInfoList, *info)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(companyInfoList)
}
