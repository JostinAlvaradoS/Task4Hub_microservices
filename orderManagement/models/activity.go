package models

import "time"

type Activity struct {
	ID           string    `json:"id"`
	RoomId       string    `json:"roomId"`
	EmployeeId   string    `json:"employeeId"`
	EmployeeName string    `json:"employeeName"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	Status       string    `json:"status"`
}
