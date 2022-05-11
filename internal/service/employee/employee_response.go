package employee

type EmployeeResponses struct { //item
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type EmployeeResquests struct { //body
	ID    int    `json:id`
	Name  string `json:"name"`
	Email string `json:"email"`
}
