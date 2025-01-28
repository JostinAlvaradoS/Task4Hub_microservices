package models

type RequiredDefaultStock struct {
	CategoryID      string `json:"categoryId"`
	CategoryName    string `json:"categoryName"`
	SubcategoryID   string `json:"subcategoryId"`
	SubcategoryName string `json:"subcategoryName"`
	ProductID       string `json:"productId"`
	ProductName     string `json:"productName"`
	Quantity        int    `json:"quantity"`
	Returnable      bool   `json:"returnable"`
}

type DefaultActivity struct {
	ID            string          `json:"id"`
	CompanyID     string          `json:"companyId"`
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	ActivityType  string          `json:"activityType"`
	RoomType      string          `json:"roomType"`
	RequiredStock []RequiredStock `json:"requiredStock"`
}
