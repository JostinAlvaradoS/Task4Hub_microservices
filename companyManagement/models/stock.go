package models

type Product struct {
	ProductID     string `json:"productId"`
	ProductName   string `json:"productName"`
	CurrentAmount int    `json:"currentAmount"`
	TargetAmount  int    `json:"targetAmount"`
}

type Subcategory struct {
	SubcategoryID   string    `json:"subcategoryId"`
	SubcategoryName string    `json:"subcategoryName"`
	Products        []Product `json:"products"`
}

type Category struct {
	CategoryID    string        `json:"categoryId"`
	CategoryName  string        `json:"categoryName"`
	Subcategories []Subcategory `json:"subcategories"`
}

type Stock struct {
	CompanyID  string     `json:"companyId"`
	Categories []Category `json:"categories"`
}

type RestockRequest struct {
	CompanyID     string `json:"companyId"`
	CategoryID    string `json:"categoryId"`
	SubcategoryID string `json:"subcategoryId"`
	ProductID     string `json:"productId"`
	Amount        int    `json:"amount"`
}
