package models

type ScoreSprints struct {
	SprintName  string
	SprintPoint int
}

type ScoreSizes struct {
	SizeID    string  `json:"sizeId"`
	SizePoint float64 `json:"sizePoint"`
}

type Sizes struct {
	Size []string
}

// ScoreOfSprints ...
type ScoreOfSprint struct {
	Title     string    `json:"title"`
	StartDate string    `json:"startDate"`
	EndDate   string    `json:"endDate"`
	Data      []float64 `json:"data"`
	Name      []string  `json:"name"`
	Day       int       `json:"-"`
}

// DataScoreTotals
type DataScoreTotal struct {
	Data []float64 `json:"data"`
	Name []string  `json:"name"`
}
