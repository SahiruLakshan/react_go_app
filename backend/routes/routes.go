package routes

import (
	"net/http"

	"employee-api/handlers"
)

func RegisterRoutes(
	mux *http.ServeMux,
	employeeHandler *handlers.EmployeeHandler,
) {

	mux.HandleFunc(
		"/api/employees",
		func(w http.ResponseWriter, r *http.Request) {

			switch r.Method {

			case http.MethodGet:
				employeeHandler.GetEmployees(w, r)

			case http.MethodPost:
				employeeHandler.CreateEmployee(w, r)

			default:
				http.Error(
					w,
					"Method not allowed",
					http.StatusMethodNotAllowed,
				)
			}
		},
	)

	mux.HandleFunc(
		"/api/employees/",
		func(w http.ResponseWriter, r *http.Request) {

			if r.Method == http.MethodDelete {
				employeeHandler.DeleteEmployee(w, r)
				return
			}

			http.Error(
				w,
				"Method not allowed",
				http.StatusMethodNotAllowed,
			)
		},
	)
}