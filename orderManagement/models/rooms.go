package models

type Rooms struct {
	ID          string      `json:"id"`
	Type        string      `json:"type"`
	Observation Observation `json:"observation"`
	Status      string      `json:"status"`
	Evidence    Evidence    `json:"evidence"`
}
