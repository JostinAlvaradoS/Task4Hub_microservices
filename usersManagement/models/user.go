package models

type User struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	Role        string   `json:"role"`
	CompanyId   string   `json:"companyId"`
	CompanyName string   `json:"companyName"`
	UID         string   `json:"uid"`
	Status      string   `json:"status"`
	Schedule    Schedule `json:"schedule"`
}

type Schedule struct {
	WorkDays []WorkDay `json:"workDays"`
}

type WorkDay struct {
	Day      string        `json:"day"`
	Schedule []ScheduleDay `json:"schedule"`
}

type ScheduleDay struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
