package models

type Lists struct {
	ListID       string `json:"listId"`
	BoardID      string `json:"-"`
	ListPriority int    `json:"-"`
	ListName     string `json:"listName"`
}
