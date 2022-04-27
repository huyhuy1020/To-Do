package employee

import (
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
