package models

type TodoLabel struct {
	ID      string `json:"id"`
	TodoID  int    `json:"todoID"`
	LabelID int    `json:"labelID"`
}
