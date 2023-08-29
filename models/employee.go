package models

type Employee struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Role string `json:"role"`
}
