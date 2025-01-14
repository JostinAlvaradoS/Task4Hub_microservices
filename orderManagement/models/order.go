package models

import "time"

type Order struct {
	ID          string    `json:"id"`
	Date        time.Time `json:"date"`
	Location    Location  `json:"location"`
	CompanyId   string    `json:"companyId"`
	ManagerId   string    `json:"managerId"`
	ManagerName string    `json:"managerName"`
	Rooms       Rooms     `json:"rooms"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	Status      string    `json:"status"`
}
