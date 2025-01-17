package models

import "time"

type Activity struct {
	ID        string    `json:"id"`
	OrderID   string    `json:"orderId"`
	RoomId    string    `json:"roomId"`
	Employee  Employee  `json:"employee"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Status    string    `json:"status"`
}
