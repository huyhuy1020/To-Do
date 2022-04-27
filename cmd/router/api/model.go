package api

type EmployeeDataResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type EmployeeResponses struct {
	Employees []EmployeeDataResponse `json:"employees"`
}
