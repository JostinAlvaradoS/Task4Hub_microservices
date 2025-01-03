package models

import "time"

type Invitation struct {
	ID          string    `json:"id"`
	Role        string    `json:"role"`
	Company     string    `json:"company"`
	ManagerId   string    `json:"manager"`
	ManagerName string    `json:"managerName"`
	ExpiresAt   time.Time `json:"expiresAt"`
}
