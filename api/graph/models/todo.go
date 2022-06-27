package models

type Todo struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	UserID      int     `json:"userID"`
	StatusID    int     `json:"statusID"`
	PriorityID  int     `json:"priorityID"`
	FinishedAt  string  `json:"finishedAt"`
}
