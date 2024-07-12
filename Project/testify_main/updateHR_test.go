package Testify

import (
	"Project/Employee"
	"Project/main/funcs"
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

	// Initialize mock database and repository
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := funcs.NewRepo(db)

	t.Run("successful update", func(t *testing.T) {
		deptID := 1
		managerID := 2
		updatedHR := Employee.HR{
			Employee: Employee.Employee{
				Name:      "John Doe Updated",
				Email:     "john.updated@example.com",
				Phone:     9876543210,
				Address:   "456 Updated St",
				DOB:       "1990-02-02",
				DeptID:    &deptID,
				ManagerID: &managerID,
			},
		}

		mock.ExpectExec("UPDATE employee e JOIN hr h ON e.EMPLOYEE_ID = h.EMPLOYEE_ID SET").
			WithArgs(updatedHR.Name, updatedHR.Email, updatedHR.Phone, updatedHR.Address, updatedHR.DOB, updatedHR.DeptID, updatedHR.ManagerID, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		jsonValue, _ := json.Marshal(updatedHR)
		req, _ := http.NewRequest(http.MethodPut, "/hr/1", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.PUT("/hr/:hrId", repo.UpdateHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response Employee.HR
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, 1, response.HR_ID)
		assert.Equal(t, "John Doe Updated", response.Name)
		assert.Equal(t, "john.updated@example.com", response.Email)
		assert.Equal(t, int64(9876543210), response.Phone)
		assert.Equal(t, "456 Updated St", response.Address)
		assert.Equal(t, "1990-02-02", response.DOB)
		assert.Equal(t, 1, *response.DeptID)
		assert.Equal(t, 2, *response.ManagerID)
	})

	t.Run("invalid HR ID", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPut, "/hr/invalid", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.PUT("/hr/:hrId", repo.UpdateHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid HR ID")
	})

	t.Run("invalid input", func(t *testing.T) {
		invalidHR := struct {
			Name string `json:"name"`
		}{
			Name: "Invalid",
		}

		jsonValue, _ := json.Marshal(invalidHR)
		req, _ := http.NewRequest(http.MethodPut, "/hr/1", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.PUT("/hr/:hrId", repo.UpdateHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("database error", func(t *testing.T) {
		deptID := 1
		managerID := 2
		updatedHR := Employee.HR{
			Employee: Employee.Employee{
				Name:      "John Doe Updated",
				Email:     "john.updated@example.com",
				Phone:     9876543210,
				Address:   "456 Updated St",
				DOB:       "1990-02-02",
				DeptID:    &deptID,
				ManagerID: &managerID,
			},
		}

		mock.ExpectExec("UPDATE employee e JOIN hr h ON e.EMPLOYEE_ID = h.EMPLOYEE_ID SET").
			WithArgs(updatedHR.Name, updatedHR.Email, updatedHR.Phone, updatedHR.Address, updatedHR.DOB, updatedHR.DeptID, updatedHR.ManagerID, 1).
			WillReturnError(sql.ErrConnDone)

		jsonValue, _ := json.Marshal(updatedHR)
		req, _ := http.NewRequest(http.MethodPut, "/hr/1", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.PUT("/hr/:hrId", repo.UpdateHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "sql: connection is already closed")
	})

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
