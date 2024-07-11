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

func TestGetHRByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("valid HR record", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"EMPLOYEE_ID", "EMPLOYEE_NAME", "EMPLOYEE_EMAIL", "EMPLOYEE_PHONE", "EMPLOYEE_ADDRESS", "EMPLOYEE_DOB", "DEPT_ID", "MANAGER_ID", "HR_ID"}).
			AddRow(1, "John HR", "john.hr@example.com", 1234567890, "123 HR St", "1990-01-01", 1, 2, 101)

		mock.ExpectQuery("SELECT e.EMPLOYEE_ID, e.EMPLOYEE_NAME, e.EMPLOYEE_EMAIL, e.EMPLOYEE_PHONE, e.EMPLOYEE_ADDRESS, e.EMPLOYEE_DOB, e.DEPT_ID, e.MANAGER_ID, h.HR_ID FROM employee e JOIN hr h ON e.EMPLOYEE_ID = h.EMPLOYEE_ID WHERE h.HR_ID = ?").
			WithArgs(101).
			WillReturnRows(rows)

		req, _ := http.NewRequest(http.MethodGet, "/hrs/101", nil)
		w := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/hrs/:hrId", repo.GetHRByID)
		router.ServeHTTP(w, req)

		expected := `{"EmpID":1,"Name":"John HR","Email":"john.hr@example.com","Phone":1234567890,"Address":"123 HR St","DOB":"1990-01-01","DeptID":1,"ManagerID":2,"HR_ID":101}`
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("HR not found", func(t *testing.T) {
		mock.ExpectQuery("SELECT e.EMPLOYEE_ID, e.EMPLOYEE_NAME, e.EMPLOYEE_EMAIL, e.EMPLOYEE_PHONE, e.EMPLOYEE_ADDRESS, e.EMPLOYEE_DOB, e.DEPT_ID, e.MANAGER_ID, h.HR_ID FROM employee e JOIN hr h ON e.EMPLOYEE_ID = h.EMPLOYEE_ID WHERE h.HR_ID = ?").
			WithArgs(999).
			WillReturnError(sql.ErrNoRows)

		req, _ := http.NewRequest(http.MethodGet, "/hrs/999", nil)
		w := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/hrs/:hrId", repo.GetHRByID)
		router.ServeHTTP(w, req)

		expected := `{"error":"HR not found"}`
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("invalid HR ID", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/hrs/invalid", nil)
		w := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/hrs/:hrId", repo.GetHRByID)
		router.ServeHTTP(w, req)

		expected := `{"error":"Invalid HR ID"}`
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("database error", func(t *testing.T) {
		mock.ExpectQuery("SELECT e.EMPLOYEE_ID, e.EMPLOYEE_NAME, e.EMPLOYEE_EMAIL, e.EMPLOYEE_PHONE, e.EMPLOYEE_ADDRESS, e.EMPLOYEE_DOB, e.DEPT_ID, e.MANAGER_ID, h.HR_ID FROM employee e JOIN hr h ON e.EMPLOYEE_ID = h.EMPLOYEE_ID WHERE h.HR_ID = ?").
			WithArgs(101).
			WillReturnError(sql.ErrConnDone)

		req, _ := http.NewRequest(http.MethodGet, "/hrs/101", nil)
		w := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/hrs/:hrId", repo.GetHRByID)
		router.ServeHTTP(w, req)

		expected := `{"error":"driver: bad connection"}`
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, expected, w.Body.String())
	})
}
