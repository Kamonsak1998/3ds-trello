package models

type Auth struct {
	BoardID  string    `json:"idBoard"`
	MemberID string    `json:"idMember"`
	Token    string    `json:"token"`
	Points   []float64 `json:"Points"`
}
