package models

type Employee struct {
	ID    int    `json:"employee_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type EmployeeRequest struct {
	ID    int    `json:"employee_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
