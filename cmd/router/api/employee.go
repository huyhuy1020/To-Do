package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

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
	router.Route("/{itemId}", func(router chi.Router) {
		router.Use(ItemContext)
		router.Get("/", getEmployee)
		router.Put("/", updateEmployee)
		router.Delete("/", deleteEmployee)
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
	// handle logic here
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	// handle logic here
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	// handle logic here
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	// handle logic here
}

func ItemContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemId := chi.URLParam(r, "itemId")
		if itemId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("item ID is required")))
			return
		}
		id, err := strconv.Atoi(itemId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid item ID")))
		}
		ctx := context.WithValue(r.Context(), itemIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
