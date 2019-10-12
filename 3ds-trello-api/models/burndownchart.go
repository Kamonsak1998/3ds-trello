package models

type IdBoard struct {
	Token   string `json:"token"`
	IdBoard string `json:"IdBoard"`
}
type Data struct {
	TitleSprint    string    `json:"titleSprint"`
	StartDate      string    `json:"startDate"`
	EndDate        string    `json:"endDate"`
	DataActualBurn []float64 `json:"actualBurn"`
	DataIDealBurn  []float64 `json:"idealBurn""`
	DatePeriod     []string  `json:"datePeriod"`
}
type BurnDownChart struct {
	Data []Data `json:"burnDownChart"`
}
type DataScoreTotalSprint struct {
	Title     []string
	StartDate []string
	EndDate   []string
	Data      []float64
	DateStamp []string
}
