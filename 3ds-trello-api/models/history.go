package models

// History ....
type History struct {
	ScoreOfSprints []ScoreOfSprint `json:"scoreOfSprint"`
	ScoreTotal     DataScoreTotal  `json: "scoreTotal"`
}
