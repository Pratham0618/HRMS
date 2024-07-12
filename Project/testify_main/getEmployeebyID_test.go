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

func TestGetEmployeeByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("successful retrieval", func(t *testing.T) {
		row := sqlmock.NewRows([]string{"EMPLOYEE_ID", "EMPLOYEE_NAME", "EMPLOYEE_EMAIL", "EMPLOYEE_PHONE", "EMPLOYEE_ADDRESS", "EMPLOYEE_DOB", "DEPT_ID", "MANAGER_ID"}).
			AddRow(1, "John Doe", "john@example.com", int64(1234567890), "123 Main St", "1990-01-01", 1, 2)

		mock.ExpectQuery("SELECT EMPLOYEE_ID, EMPLOYEE_NAME, EMPLOYEE_EMAIL, EMPLOYEE_PHONE, EMPLOYEE_ADDRESS, EMPLOYEE_DOB, DEPT_ID, MANAGER_ID FROM employee WHERE EMPLOYEE_ID = ?").
			WithArgs("1").WillReturnRows(row)

		req, _ := http.NewRequest(http.MethodGet, "/employees/1", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/employees/:id", repo.GetEmployeeByID)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		expected := `{"emp_id":1,"name":"John Doe","email":"john@example.com","phone":1234567890,"address":"123 Main St","dob":"1990-01-01","dept_id":1,"manager_id":2}`
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("employee not found", func(t *testing.T) {
		mock.ExpectQuery("SELECT EMPLOYEE_ID, EMPLOYEE_NAME, EMPLOYEE_EMAIL, EMPLOYEE_PHONE, EMPLOYEE_ADDRESS, EMPLOYEE_DOB, DEPT_ID, MANAGER_ID FROM employee WHERE EMPLOYEE_ID = ?").
			WithArgs("2").WillReturnError(sql.ErrNoRows)

		req, _ := http.NewRequest(http.MethodGet, "/employees/2", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/employees/:id", repo.GetEmployeeByID)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		expected := `{"error":"employee not found"}`
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("database error", func(t *testing.T) {
		mock.ExpectQuery("SELECT EMPLOYEE_ID, EMPLOYEE_NAME, EMPLOYEE_EMAIL, EMPLOYEE_PHONE, EMPLOYEE_ADDRESS, EMPLOYEE_DOB, DEPT_ID, MANAGER_ID FROM employee WHERE EMPLOYEE_ID = ?").
			WithArgs("1").WillReturnError(sql.ErrConnDone)

		req, _ := http.NewRequest(http.MethodGet, "/employees/1", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/employees/:id", repo.GetEmployeeByID)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		expected := `{"error":"sql: connection is already closed"}`
		assert.JSONEq(t, expected, w.Body.String())
	})
}
