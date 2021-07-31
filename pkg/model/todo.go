package model

type Todo struct {
	Id          int64  `json:"id"`
	TaskName    string `json:"taskName"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Creator     string `json:"creator"`
}

var AllTasks []Todo
