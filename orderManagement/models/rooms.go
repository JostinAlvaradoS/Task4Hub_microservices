package models

type Observation struct {
	Description string `json:"description"`
	Image       string `json:"image"`
}

type Evidence struct {
	ImageBefore string `json:"imageBefore"`
	ImageAfter  string `json:"imageAfter"`
}

type Locker struct {
	Image    string `json:"image"`
	Password string `json:"password"`
}

type Rooms struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Observation Observation `json:"observation"`
	Status      string      `json:"status"`
	Evidence    Evidence    `json:"evidence"`
	Locker      Locker      `json:"locker"`
	BedCount    int         `json:"bedCount"`
	FloorType   string      `json:"floorType"`
}
