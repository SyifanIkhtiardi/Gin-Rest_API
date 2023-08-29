package models

// Schema to validate employee input
type CreateEmployeeInput struct {
	Name string `json:"name" binding:"required"`
	Role string `json:"role" binding:"required"`
}
