package Testify

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEmployees(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("successful retrieval", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"EMPLOYEE_ID", "EMPLOYEE_NAME", "EMPLOYEE_EMAIL", "EMPLOYEE_PHONE", "EMPLOYEE_ADDRESS", "EMPLOYEE_DOB", "DEPT_ID", "MANAGER_ID"}).
			AddRow(1, "John Doe", "john@example.com", int64(1234567890), "123 Main St", "1990-01-01", 1, 2).
			AddRow(2, "Jane Smith", "jane@example.com", int64(9876543210), "456 Elm St", "1995-05-05", 2, 1)

		mock.ExpectQuery("SELECT EMPLOYEE_ID, EMPLOYEE_NAME, EMPLOYEE_EMAIL, EMPLOYEE_PHONE, EMPLOYEE_ADDRESS, EMPLOYEE_DOB, DEPT_ID, MANAGER_ID FROM employee").
			WillReturnRows(rows)

		req, _ := http.NewRequest(http.MethodGet, "/employees", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/employees", repo.GetEmployees)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		expected := `[
            {"emp_id":1,"name":"John Doe","email":"john@example.com","phone":1234567890,"address":"123 Main St","dob":"1990-01-01","dept_id":1,"manager_id":2},
            {"emp_id":2,"name":"Jane Smith","email":"jane@example.com","phone":9876543210,"address":"456 Elm St","dob":"1995-05-05","dept_id":2,"manager_id":1}
        ]`
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("no employees found", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"EMPLOYEE_ID", "EMPLOYEE_NAME", "EMPLOYEE_EMAIL", "EMPLOYEE_PHONE", "EMPLOYEE_ADDRESS", "EMPLOYEE_DOB", "DEPT_ID", "MANAGER_ID"})

		mock.ExpectQuery("SELECT EMPLOYEE_ID, EMPLOYEE_NAME, EMPLOYEE_EMAIL, EMPLOYEE_PHONE, EMPLOYEE_ADDRESS, EMPLOYEE_DOB, DEPT_ID, MANAGER_ID FROM employee").
			WillReturnRows(rows)

		req, _ := http.NewRequest(http.MethodGet, "/employees", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/employees", repo.GetEmployees)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		expected := `[]`
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("database error", func(t *testing.T) {
		mock.ExpectQuery("SELECT EMPLOYEE_ID, EMPLOYEE_NAME, EMPLOYEE_EMAIL, EMPLOYEE_PHONE, EMPLOYEE_ADDRESS, EMPLOYEE_DOB, DEPT_ID, MANAGER_ID FROM employee").
			WillReturnError(sql.ErrConnDone)

		req, _ := http.NewRequest(http.MethodGet, "/employees", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/employees", repo.GetEmployees)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		expected := `{"error":"sql: connection is already closed"}`
		assert.JSONEq(t, expected, w.Body.String())
	})
}
