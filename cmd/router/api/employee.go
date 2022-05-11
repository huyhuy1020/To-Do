package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"todo/internal/models"
	response "todo/internal/responses"
	employeeService "todo/internal/service/employee"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var (
	EmployeeID = "EmpmID"
	itemIDKey  = "itemID"
)

func Employee(router chi.Router) {
	router.Get("/", getAllEmployees)
	router.Post("/", createEmployee)
	router.Route("/{id}", func(router chi.Router) {
		router.Get("/detail", getEmployee)
		router.Put("/update", updateEmployee)
		router.Delete("/delete", deleteEmployee)
	})
}

func getAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := employeeService.GetAllEmployee(dbInstance)
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	result := translateToDistrictResponses(employees)
	response.JSON(w, http.StatusOK, result)
}

func translateToDistrictResponses(employees []employeeService.EmployeeResponses) EmployeeResponses {
	employeeResponses := EmployeeResponses{}

	for _, employee := range employees {
		employeeResponses.Employees = append(employeeResponses.Employees, EmployeeDataResponse{
			ID:    employee.ID,
			Name:  employee.Name,
			Email: employee.Email,
		})
	}

	return employeeResponses
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var employee employeeService.EmployeeResquests
	err = json.Unmarshal(reqBytes, &employee)
}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func Addition(employees []employeeService.EmployeeResponses) EmployeeResponses {
	newEmployee := EmployeeResponses{}
	for _, employee := range employees {
		newEmployee.Employees = append(newEmployee.Employees, EmployeeDataResponse{
			ID:    employee.ID,
			Name:  employee.Name,
			Email: employee.Email,
		})
	}
	return newEmployee
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	// handle logic here
	empid := chi.URLParam(r, "id")
	if empid == "" {
		render.Render(w, r, ServerErrorRenderer(errors.New("id is empty")))
		return
	}

	parseId, err := strconv.Atoi(empid)
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(errors.New("id parse failed")))
		return
	}

	resp, err := employeeService.GetEmployeeByID(dbInstance, parseId)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	response.JSON(w, http.StatusOK, resp)
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	// handle logic here
	empID := r.Context().Value(EmployeeID).(int)
	empData := models.Employee{}
	if err := render.Bind(r, &empData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	item, err := employeeService.updateEmp(empID)
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	// handle logic here
	empmID := r.Context().Value(EmployeeID).(int)
	_, err := employeeService.DeleteEmployees(dbInstance, empmID)
	if err != nil {
		if err == employeeService.ErrNoMatch(empmID) {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
	}
	return
}
