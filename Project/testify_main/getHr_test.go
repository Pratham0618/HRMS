package Testify

import (
	"Project/main/funcs"
	"database/sql"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var db *sql.DB
var mock sqlmock.Sqlmock
var repo *funcs.Repo

func init() {
	var err error
	db, mock, err = sqlmock.New()
	if err != nil {
		panic("failed to open a stub database connection")
	}
	repo = funcs.NewRepo(db)
}

func TestGetHR(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("valid HR records", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"EMPLOYEE_ID", "EMPLOYEE_NAME", "EMPLOYEE_EMAIL", "EMPLOYEE_PHONE", "EMPLOYEE_ADDRESS", "EMPLOYEE_DOB", "DEPT_ID", "MANAGER_ID", "HR_ID"}).
			AddRow(1, "John HR", "john.hr@example.com", 1234567890, "123 HR St", "1990-01-01", 1, 2, 101).
			AddRow(2, "Jane HR", "jane.hr@example.com", 9087654321, "456 HR St", "1992-02-02", 2, 3, 102)

		mock.ExpectQuery("SELECT e.EMPLOYEE_ID, e.EMPLOYEE_NAME, e.EMPLOYEE_EMAIL, e.EMPLOYEE_PHONE, e.EMPLOYEE_ADDRESS, e.EMPLOYEE_DOB, e.DEPT_ID, e.MANAGER_ID, h.HR_ID FROM employee e JOIN hr h ON e.EMPLOYEE_ID = h.EMPLOYEE_ID").
			WillReturnRows(rows)

		req, _ := http.NewRequest(http.MethodGet, "/hrs", nil)
		w := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/hrs", repo.GetHR)
		router.ServeHTTP(w, req)

		expected := `[{"emp_id":1,"name":"John HR","email":"john.hr@example.com","phone":1234567890,"address":"123 HR St","dob":"1990-01-01","dept_id":1,"manager_id":2,"hr_id":101},{"emp_id":2,"name":"Jane HR","email":"jane.hr@example.com","phone":9087654321,"address":"456 HR St","dob":"1992-02-02","dept_id":2,"manager_id":3,"hr_id":102}]`
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("no HR records", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"EMPLOYEE_ID", "EMPLOYEE_NAME", "EMPLOYEE_EMAIL", "EMPLOYEE_PHONE", "EMPLOYEE_ADDRESS", "EMPLOYEE_DOB", "DEPT_ID", "MANAGER_ID", "HR_ID"})

		mock.ExpectQuery("SELECT e.EMPLOYEE_ID, e.EMPLOYEE_NAME, e.EMPLOYEE_EMAIL, e.EMPLOYEE_PHONE, e.EMPLOYEE_ADDRESS, e.EMPLOYEE_DOB, e.DEPT_ID, e.MANAGER_ID, h.HR_ID FROM employee e JOIN hr h ON e.EMPLOYEE_ID = h.EMPLOYEE_ID").
			WillReturnRows(rows)

		req, _ := http.NewRequest(http.MethodGet, "/hrs", nil)
		w := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/hrs", repo.GetHR)
		router.ServeHTTP(w, req)

		expected := `[]`
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("database error", func(t *testing.T) {
		mock.ExpectQuery("SELECT e.EMPLOYEE_ID, e.EMPLOYEE_NAME, e.EMPLOYEE_EMAIL, e.EMPLOYEE_PHONE, e.EMPLOYEE_ADDRESS, e.EMPLOYEE_DOB, e.DEPT_ID, e.MANAGER_ID, h.HR_ID FROM employee e JOIN hr h ON e.EMPLOYEE_ID = h.EMPLOYEE_ID").
			WillReturnError(sql.ErrConnDone)

		req, _ := http.NewRequest(http.MethodGet, "/hrs", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/hrs", repo.GetHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		// Verify the error message
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("failed to unmarshal JSON response: %v", err)
		}

		expectedError := "sql: connection is already closed" // Update this if your error message is different
		actualError, ok := response["error"].(string)
		if !ok || actualError != expectedError {
			t.Errorf("expected error message %q, got %q", expectedError, actualError)
		}
	})
}
