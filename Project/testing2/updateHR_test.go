package Testify

import (
	"Project/Employee"
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateHR(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("successful update", func(t *testing.T) {
		mock.ExpectExec("UPDATE employee e JOIN hr h ON e.EMPLOYEE_ID = h.EMPLOYEE_ID").
			WithArgs("John Updated", "john.updated@example.com", int64(9876543210), "456 Updated St", "1991-02-02", sqlmock.AnyArg(), sqlmock.AnyArg(), 101).
			WillReturnResult(sqlmock.NewResult(0, 1))

		updatedHR := Employee.HR{
			Employee: Employee.Employee{
				Name:      "John Updated",
				Email:     "john.updated@example.com",
				Phone:     9876543210,
				Address:   "456 Updated St",
				DOB:       "1991-02-02",
				DeptID:    func(i int) *int { return &i }(2),
				ManagerID: func(i int) *int { return &i }(3),
			},
			HR_ID: 101,
		}

		jsonValue, _ := json.Marshal(updatedHR)
		req, _ := http.NewRequest(http.MethodPut, "/hrs/101", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.PUT("/hrs/:hrId", repo.UpdateHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		expected := `{"message":"HR updated successfully"}`
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("hr not found", func(t *testing.T) {
		mock.ExpectExec("UPDATE employee e JOIN hr h ON e.EMPLOYEE_ID = h.EMPLOYEE_ID").
			WithArgs("John Updated", "john.updated@example.com", int64(9876543210), "456 Updated St", "1991-02-02", sqlmock.AnyArg(), sqlmock.AnyArg(), 999).
			WillReturnResult(sqlmock.NewResult(0, 0))

		updatedHR := Employee.HR{
			Employee: Employee.Employee{
				Name:      "John Updated",
				Email:     "john.updated@example.com",
				Phone:     9876543210,
				Address:   "456 Updated St",
				DOB:       "1991-02-02",
				DeptID:    func(i int) *int { return &i }(2),
				ManagerID: func(i int) *int { return &i }(3),
			},
			HR_ID: 999,
		}

		jsonValue, _ := json.Marshal(updatedHR)
		req, _ := http.NewRequest(http.MethodPut, "/hrs/999", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.PUT("/hrs/:hrId", repo.UpdateHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		expected := `{"error":"HR not found"}`
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("invalid input", func(t *testing.T) {
		updatedHR := Employee.HR{
			Employee: Employee.Employee{
				Name: "John Updated",
				// Missing required fields
			},
			HR_ID: 101,
		}

		jsonValue, _ := json.Marshal(updatedHR)
		req, _ := http.NewRequest(http.MethodPut, "/hrs/101", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.PUT("/hrs/:hrId", repo.UpdateHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		expected := `{"error":"Invalid input"}`
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("database error", func(t *testing.T) {
		mock.ExpectExec("UPDATE employee e JOIN hr h ON e.EMPLOYEE_ID = h.EMPLOYEE_ID").
			WithArgs("John Updated", "john.updated@example.com", int64(9876543210), "456 Updated St", "1991-02-02", sqlmock.AnyArg(), sqlmock.AnyArg(), 101).
			WillReturnError(sql.ErrConnDone)

		updatedHR := Employee.HR{
			Employee: Employee.Employee{
				Name:      "John Updated",
				Email:     "john.updated@example.com",
				Phone:     9876543210,
				Address:   "456 Updated St",
				DOB:       "1991-02-02",
				DeptID:    func(i int) *int { return &i }(2),
				ManagerID: func(i int) *int { return &i }(3),
			},
			HR_ID: 101,
		}

		jsonValue, _ := json.Marshal(updatedHR)
		req, _ := http.NewRequest(http.MethodPut, "/hrs/101", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.PUT("/hrs/:hrId", repo.UpdateHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		expected := `{"error":"driver: bad connection"}`
		assert.JSONEq(t, expected, w.Body.String())
	})
}
