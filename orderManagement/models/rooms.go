package models

type Rooms struct {
	ID          string      `json:"id"`
	Type        string      `json:"type"`
	Quantity    int         `json:"quantity"`
	Observation Observation `json:"observation"`
	Status      string      `json:"status"`
	Evidence    Evidence    `json:"evidence"`
}
