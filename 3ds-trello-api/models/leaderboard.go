package models

type LeaderBoard struct {
	MemberNames  string  `json:"name"`
	MemberPoints float64 `json:"point"`
	AvatarsHash  string  `json:"avatar"`
}
type LeaderBoards struct {
	Leader []LeaderBoard `json:"board"`
}
