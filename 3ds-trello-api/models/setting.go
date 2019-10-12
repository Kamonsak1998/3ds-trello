package models

type Setting struct {
	SprintDate Sprints `json:"sprintDate"`
	ScoreSize  Auth    `json:"scoreSize"`
	SelectList []Lists `json:"selectList"`
}

type SettingRes struct {
	Date      interface{} `json:"date"`
	ScoreSize interface{} `json:"scoreSize"`
	List      interface{} `json:"lists"`
}
