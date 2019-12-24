package models

type Todo struct {
	Model
	Name string `json:"name"`
	Done bool   `json:"done"`
}
