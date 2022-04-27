package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Tasks struct {
	Task_id   string `json:"taskID"`
	Task_name string `json:"task_name"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "0942877351"
	dbname   = "todo"
)

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func GetTaskByEmployeeID(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	rows, err := db.Query("SELECT * FROM employee")
	if err != nil {
		log.Fatal(err)
	}

	var em []Employee

	for rows.Next() {
		var emp Employee
		rows.Scan(&emp.ID, &emp.Name, &emp.Email)
		em = append(em, emp)
	}

	peopleBytes, _ := json.MarshalIndent(em, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(peopleBytes)

	defer rows.Close()
	defer db.Close()
}

func PostByEmployeeID(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	var e Employee
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `INSERT INTO Employee (emp_id,name,email ) VALUES ($1, $2, $3)`
	_, err = db.Exec(sqlStatement, e.ID, e.Name, e.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}

func main() {
	http.HandleFunc("/", GetTaskByEmployeeID)
	http.HandleFunc("/insert", PostByEmployeeID)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
