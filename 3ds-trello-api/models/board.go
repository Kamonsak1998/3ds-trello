package models

type Boards struct {
	BoardId         string `json:"idBoard"`
	BoardName       string `json:"boardName"`
	ImageBackground string `gorm:"-" json:"imageBackground"`
	ColorBackground string `gorm:"-" json:"colorBackground"`
}
