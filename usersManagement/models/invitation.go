package models

import "time"

type Invitation struct {
	ID          string    `json:"id"`
	Role        string    `json:"role"`
	CompanyId   string    `json:"companyId"`
	CompanyName string    `json:"companyName"`
	ManagerId   string    `json:"managerId"`
	ManagerName string    `json:"managerName"`
	ExpiresAt   time.Time `json:"expiresAt"`
}
