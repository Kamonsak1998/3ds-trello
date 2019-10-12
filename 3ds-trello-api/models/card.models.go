package models

import "time"

type Cards struct {
	CardID      string
	BoardID     string
	MemberID    string
	CardName    string
	CardTime    time.Time
	CardPoint   float64
	SprintName  string
}

type Actions struct {
	CardID  string
	Date    time.Time
}