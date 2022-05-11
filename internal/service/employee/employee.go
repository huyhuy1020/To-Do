package employee

import (
	"todo/database"
	"todo/internal/models"
	"todo/internal/repository/employee"
)

func GetAllEmployee(db database.Database) ([]EmployeeResponses, error) {
	employees, err := employee.RetrieveEmployee(db)
	resp := translateToEmployeeResp(employees)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func translateToEmployeeResp(employees []models.Employee) []EmployeeResponses {
	employeeResponses := []EmployeeResponses{}
	for _, employee := range employees {
		employeeResponse := EmployeeResponses{
			ID:    employee.ID,
			Name:  employee.Name,
			Email: employee.Email,
		}
		employeeResponses = append(employeeResponses, employeeResponse)
	}

	return employeeResponses
}

func AddEmployee(db database.Database) error {
	err := AddEmployee(db)
	if err != nil {
		return nil
	}
	return err
}

func GetEmployeeByID(db database.Database, id int) (EmployeeResponses, error) {
	employee, err := employee.RetrieveEmployeeById(db, id)
	resp := translateToGetID(employee)
	if err != nil {
		return EmployeeResponses{}, err
	}

	return resp, nil
}

func translateToGetID(employee models.Employee) EmployeeResponses {
	return EmployeeResponses{
		ID:    employee.ID,
		Email: employee.Email,
		Name:  employee.Name,
	}
}

func DeleteEmployees(db database.Database, id int) error {
	emp := employee.DeleteEmployee(db, id)
	err := emp
	if err != nil {
		return nil
	}
	return err
}
