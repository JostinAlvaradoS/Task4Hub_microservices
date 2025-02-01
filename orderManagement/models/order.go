package models

type Order struct {
	ID          string     `json:"id"`
	Date        string     `json:"date"`
	Type        string     `json:"type"`
	Location    Location   `json:"location"`
	CompanyId   string     `json:"companyId"`
	ManagerId   string     `json:"managerId"`
	ManagerName string     `json:"managerName"`
	Rooms       []Rooms    `json:"rooms"`
	Employees   []Employee `json:"employees"`
	StartDate   string     `json:"startDate"`
	EndDate     string     `json:"endDate"`
	Status      string     `json:"status"`
}

type Location struct {
	Address   string `json:"address"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}
