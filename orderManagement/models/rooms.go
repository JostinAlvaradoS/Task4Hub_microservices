package models

type Observation struct {
	Description string `json:"description"`
	Image       string `json:"image"`
}

type Evidence struct {
	ImageBefore string `json:"imageBefore"`
	ImageAfter  string `json:"imageAfter"`
}

type Rooms struct {
	ID          string      `json:"id"`
	Type        string      `json:"type"`
	Observation Observation `json:"observation"`
	Status      string      `json:"status"`
	Evidence    Evidence    `json:"evidence"`
}
