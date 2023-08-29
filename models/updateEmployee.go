package models

// Schema to validate employee input
type UpdateEmployeeInput struct {
	Name string `json:"name"`
	Role string `json:"role"`
}
