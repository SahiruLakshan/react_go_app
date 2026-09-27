package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"employee-api/models"
)

type EmployeeHandler struct {
	DB *sql.DB
}

func NewEmployeeHandler(db *sql.DB) *EmployeeHandler {
	return &EmployeeHandler{
		DB: db,
	}
}

func (h *EmployeeHandler) GetEmployees(
	w http.ResponseWriter,
	r *http.Request,
) {

	rows, err := h.DB.QueryContext(
		context.Background(),
		`
		SELECT
			id,
			name,
			email,
			department,
			salary
		FROM employees
		ORDER BY id
		`,
	)

	if err != nil {
		http.Error(
			w,
			"Failed to get employees",
			http.StatusInternalServerError,
		)
		return
	}

	defer rows.Close()

	employees := []models.Employee{}

	for rows.Next() {

		var employee models.Employee

		err := rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Email,
			&employee.Department,
			&employee.Salary,
		)

		if err != nil {
			http.Error(
				w,
				"Failed to read employee",
				http.StatusInternalServerError,
			)
			return
		}

		employees = append(employees, employee)
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(employees)
}

func (h *EmployeeHandler) CreateEmployee(
	w http.ResponseWriter,
	r *http.Request,
) {
	var employee models.Employee

	err := json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		http.Error(
			w,
			"Invalid request",
			http.StatusBadRequest,
		)
		return
	}

	var id int64

	_, err = h.DB.ExecContext(
		context.Background(),
		`
		INSERT INTO employees
		(
			name,
			email,
			department,
			salary
		)
		VALUES
		(
			:1,
			:2,
			:3,
			:4
		)
		RETURNING id INTO :5
		`,
		employee.Name,
		employee.Email,
		employee.Department,
		employee.Salary,
		sql.Out{Dest: &id},
	)

	if err != nil {
		http.Error(
			w,
			"Failed to create employee",
			http.StatusInternalServerError,
		)
		return
	}

	employee.ID = id

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(employee)
}

func (h *EmployeeHandler) DeleteEmployee(
	w http.ResponseWriter,
	r *http.Request,
) {

	idText := strings.TrimPrefix(
		r.URL.Path,
		"/api/employees/",
	)

	id, err := strconv.ParseInt(idText, 10, 64)

	if err != nil {
		http.Error(
			w,
			"Invalid employee ID",
			http.StatusBadRequest,
		)
		return
	}

	result, err := h.DB.ExecContext(
		context.Background(),
		`
		DELETE FROM employees
		WHERE id = :1
		`,
		id,
	)

	if err != nil {
		http.Error(
			w,
			"Failed to delete employee",
			http.StatusInternalServerError,
		)
		return
	}

	count, _ := result.RowsAffected()

	if count == 0 {
		http.Error(
			w,
			"Employee not found",
			http.StatusNotFound,
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}