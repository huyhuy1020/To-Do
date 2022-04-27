package database

import (
	"database/sql"

	"gitlab.com/idoko/bucketeer/models"
)

func (db Database) GetAllEmployee() (*models.Employee__, error) {
	list := &models.Employee__{}
	rows, err := db.Conn.Query("SELECT * FROM Employee ORDER BY ID DESC")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var emp models.Employee
		err := rows.Scan(&emp.ID, &emp.Name, &emp.Email)
		if err != nil {
			return list, err
		}
		list.Em = append(list.Em, emp)
	}
	return list, nil
}
func (db Database) AddEmployee(emp *models.Employee) error {
	var emp_id int
	query := `INSERT INTO Employee (name, email) VALUES ($1, $2) RETURNING emp_id`
	err := db.Conn.QueryRow(query, emp.Name, emp.Email).Scan(&emp_id)
	if err != nil {
		return err
	}
	emp.ID = emp_id
	return nil
}
func (db Database) GetEmployeeByEmployee_Id(empID int) (models.Employee, error) {
	emm := models.Employee{}
	query := `SELECT * FROM Employee WHERE employee_id = $1;`
	row := db.Conn.QueryRow(query, empID)
	switch err := row.Scan(&emm.ID, &emm.Name, &emm.Email); err {
	case sql.ErrNoRows:
		return emm, ErrNoMatch
	default:
		return emm, err
	}
}
func (db Database) DeleteEmployee(empID int) error {
	query := `DELETE FROM  WHERE employee_id = $1;`
	_, err := db.Conn.Exec(query, empID)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}
func (db Database) UpdateItem(empID int, empData models.Employee) (models.Employee, error) {
	emp := models.Employee{}
	query := `UPDATE Employee SET name=$1, email=$2 WHERE employee_id=$3 RETURNING emoloyee_id, name, email;`
	err := db.Conn.QueryRow(query, empData.Name, empData.Email, empID).Scan(&emp.ID, &emp.Name, &emp.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return emp, ErrNoMatch
		}
		return emp, err
	}
	return emp, nil
}
