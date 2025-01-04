package models

type CompanyInfo struct {
	Company      string `json:"company"`
	SuperManager string `json:"superManager"`
	Employees    int    `json:"employees"`
}
