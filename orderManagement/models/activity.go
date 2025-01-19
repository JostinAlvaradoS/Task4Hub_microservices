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
	CategoryName  string `json:"categoryName"`
	SubcategoryID string `json:"subcategoryId"`
	Subcategory   string `json:"subcategory"`
	ProductID     string `json:"productId"`
	ProductName   string `json:"productName"`
	Quantity      int    `json:"quantity"`
}
