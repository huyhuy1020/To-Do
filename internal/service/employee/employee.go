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

func GetEmployeeByEmployee_Id(db database.Database) ([]EmployeeResquests, error) {
	employees1, err := employee.GetEmployeeById(db)
	resp := translateToGetID(employees1)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func translateToGetID(employees []models.EmployeeRequest) []EmployeeResquests {
	emp := []EmployeeResquests{}
	for _, employee := range employees {
		empreq := models.EmployeeRequest{
			ID:    employee.ID,
			Name:  employee.Name,
			Email: employee.Email,
		}
		emp = append(emp, EmployeeResquests(empreq))
	}
	return emp
}

func DeleteEmployees(db database.Database) ([]EmployeeResponses, error) {
	employees, err := employee.DeleteEmployee(db)
	resp := translateToEmployeeResp(employees)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func updateEmp(db database.Database) ([]EmployeeResponses, error) {
	employees, err := employee.updateEmployee(db)
	resp := translateToEmployeeResp(employees)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
