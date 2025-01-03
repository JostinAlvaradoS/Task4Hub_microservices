package models

type Company struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	ABN  string `json:"abn"`
	Logo string `json:"logo"`
}
