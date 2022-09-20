package entity

type Task struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}
