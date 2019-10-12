package models

import "time"

type Sprints struct {
	SprintId   *int
	BoardId    string `json:"idBoard"`
	SprintName string
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	SprintDay  int       `json:"sprintDay"`
}
