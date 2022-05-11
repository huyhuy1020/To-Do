package employee

import (
	"database/sql"
	"todo/database"
	"todo/internal/models"
)

func RetrieveEmployee(db database.Database) ([]models.Employee, error) {
	employees := []models.Employee{}
	rows, err := db.Conn.Query("SELECT * FROM Employee ORDER BY employee_id DESC")
	if err != nil {
		return employees, err
	}

	for rows.Next() {
		var emp models.Employee
		err := rows.Scan(&emp.ID, &emp.Name, &emp.Email)
		if err != nil {
			return employees, err
		}
		employees = append(employees, emp)
	}

	return employees, nil
}

func AddEmploy(db database.Database, emp *models.Employee) error {
	var InsertID int
	err := db.Conn.QueryRow("INSERT INTO Employee(name, email VALUES($1, $2) returning employee_id", emp.Name, emp.Email).Scan(&InsertID)
	if err != nil {
		return err
	}
	emp.ID = InsertID
	return nil
}

func RetrieveEmployeeById(db database.Database, empId int) (models.Employee, error) {
	emp := models.Employee{}
	query := `SELECT * FROM Employee WHERE emoloyee_id = $1;`
	row := db.Conn.QueryRow(query, empId)
	switch err := row.Scan(&emp.ID, &emp.Name, &emp.Email); err {
	case sql.ErrNoRows:
		return emp, database.ErrNoMatch
	default:
		return emp, err
	}
}

func DeleteEmployee(db database.Database, empId int) error {
	query := `DELETE FROM Employee WHERE employee_id = $1`
	_, err := db.Conn.Exec(query, empId)
	switch err {
	case sql.ErrNoRows:
		return database.ErrNoMatch
	default:
		return err
	}

}
func updateEmployees(db database.Database, empId int, empData models.Employee) (models.Employee, error) {
	employee := models.Employee{}
	query := `UPDATE Employee SET name=$1, email=$2 WHERE employee_id=$3 RETURNING employee_id, name, description;`
	err := db.Conn.QueryRow(query, empData.Name, empData.Email, empId).Scan(&employee.ID, &employee.Name, &employee.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return employee, database.ErrNoMatch
		}
		return employee, err
	}
	return employee, nil
}
