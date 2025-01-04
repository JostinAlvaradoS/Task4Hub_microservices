package models

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	CompanyId   string `json:"companyId"`
	CompanyName string `json:"companyName"`
	UID         string `json:"uid"`
	Status      string `json:"status"`
}
