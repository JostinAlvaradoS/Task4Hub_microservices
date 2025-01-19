package models

import "time"

type Activity struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	CompanyID     string          `json:"companyId"`
	Description   string          `json:"description"`
	OrderID       string          `json:"orderId"`
	RoomId        string          `json:"roomId"`
	Employee      Employee        `json:"employee"`
	StartDate     time.Time       `json:"startDate"`
	EndDate       time.Time       `json:"endDate"`
	Status        string          `json:"status"`
	RequiredStock []RequiredStock `json:"requiredStock"`
}

type RequiredStock struct {
	ID            string `json:"id"`
	CategoryID    string `json:"categoryId"`
	SubcategoryID string `json:"subcategoryId"`
	ProductID     string `json:"productId"`
	Quantity      int    `json:"quantity"`
}
