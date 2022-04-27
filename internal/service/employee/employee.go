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

// func AddEmployee(db database.Database, emp *models.EmployeeDataResponse) error {
// 	var emp_id int
// 	query := `INSERT INTO Employee (name, email) VALUES ($1, $2) RETURNING emp_id`
// 	err := db.Conn.QueryRow(query, emp.Name, emp.Email).Scan(&emp_id)
// 	if err != nil {
// 		return err
// 	}
// 	emp.ID = emp_id
// 	return nil
// }

// func GetEmployeeByEmployee_Id(db database.Database, empID int) (models.EmployeeDataResponse, error) {
// 	emm := models.EmployeeDataResponse{}
// 	query := `SELECT * FROM Employee WHERE employee_id = $1;`
// 	row := db.Conn.QueryRow(query, empID)
// 	switch err := row.Scan(&emm.ID, &emm.Name, &emm.Email); err {
// 	case sql.ErrNoRows:
// 		return emm, ErrNoMatch
// 	default:
// 		return emm, err
// 	}
// }

// func DeleteEmployee(db database.Database, empID int) error {
// 	query := `DELETE FROM  WHERE employee_id = $1;`
// 	_, err := db.Conn.Exec(query, empID)
// 	switch err {
// 	case sql.ErrNoRows:
// 		return ErrNoMatch
// 	default:
// 		return err
// 	}
// }

// func UpdateItem(db database.Database, empID int, empData models.EmployeeDataResponse) (models.EmployeeDataResponse, error) {
// 	emp := models.EmployeeDataResponse{}
// 	query := `UPDATE Employee SET name=$1, email=$2 WHERE employee_id=$3 RETURNING emoloyee_id, name, email;`
// 	err := db.Conn.QueryRow(query, empData.Name, empData.Email, empID).Scan(&emp.ID, &emp.Name, &emp.Email)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return emp, ErrNoMatch
// 		}
// 		return emp, err
// 	}
// 	return emp, nil
// }
