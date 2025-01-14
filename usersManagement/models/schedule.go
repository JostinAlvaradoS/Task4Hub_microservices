package models

type Schedule struct {
	WorkDays []WorkDay `json:"workDays"`
}

type WorkDay struct {
	Day       string `json:"day"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
