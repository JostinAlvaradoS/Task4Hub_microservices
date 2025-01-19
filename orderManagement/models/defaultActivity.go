package models

type RequiredDefaultStock struct {
	ID              string `json:"id"`
	CategoryID      string `json:"categoryId"`
	CategoryName    string `json:"categoryName"`
	SubcategoryID   string `json:"subcategoryId"`
	SubcategoryName string `json:"subcategoryName"`
	ProductID       string `json:"productId"`
	ProductName     string `json:"productName"`
	Quantity        int    `json:"quantity"`
}

type DefaultActivity struct {
	ID            string          `json:"id"`
	CompanyID     string          `json:"companyId"`
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	Type          string          `json:"type"`
	RequiredStock []RequiredStock `json:"requiredStock"`
}
