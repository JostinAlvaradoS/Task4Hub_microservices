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
	ActivityType  string          `json:"activityType"`
	StartDate     time.Time       `json:"startDate"`
	EndDate       time.Time       `json:"endDate"`
	Status        string          `json:"status"`
	RequiredStock []RequiredStock `json:"requiredStock"`
}

type RequiredStock struct {
	CategoryID      string `json:"categoryId"`
	CategoryName    string `json:"categoryName"`
	SubcategoryID   string `json:"subcategoryId"`
	SubcategoryName string `json:"subcategoryName"`
	ProductID       string `json:"productId"`
	ProductName     string `json:"productName"`
	Quantity        int    `json:"quantity"`
	Returnable      bool   `json:"returnable"`
}

type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
