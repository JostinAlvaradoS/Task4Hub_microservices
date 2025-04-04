package models

type Warehouse struct {
	Name      string `json:"name"`
	Location  string `json:"location"`
	AcessCode string `json:"accessCode"`
	Type      string `json:"type"`
}

type Company struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	ABN           string    `json:"abn"`
	Logo          string    `json:"logo"`
	EmployeeCount int       `json:"employeeCount"`
	TypeCompany   string    `json:"typeCompany"`
	Warehouse     Warehouse `json:"warehouse"`
}
