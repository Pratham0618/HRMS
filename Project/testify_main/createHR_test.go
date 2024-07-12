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

func TestCreateHR(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Initialize mock database and repository
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := funcs.NewRepo(db)

	t.Run("successful creation", func(t *testing.T) {
		// Mock begin transaction
		mock.ExpectBegin()

		// Mock employee insertion
		mock.ExpectExec("INSERT INTO employee").
			WithArgs("John Doe", "john@example.com", int64(1234567890), "123 Main St", "1990-01-01", 1, 2).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// Mock hr insertion
		mock.ExpectExec("INSERT INTO hr").
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// Mock commit
		mock.ExpectCommit()

		deptID := 1
		managerID := 2
		newHR := Employee.HR{
			Employee: Employee.Employee{
				Name:      "John Doe",
				Email:     "john@example.com",
				Phone:     1234567890,
				Address:   "123 Main St",
				DOB:       "1990-01-01",
				DeptID:    &deptID,
				ManagerID: &managerID,
			},
		}

		jsonValue, _ := json.Marshal(newHR)
		req, _ := http.NewRequest(http.MethodPost, "/hr", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/hr", repo.CreateHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response Employee.HR
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, 1, response.EmpID)
		assert.Equal(t, 1, response.HR_ID)
		assert.Equal(t, "John Doe", response.Name)
		assert.Equal(t, "john@example.com", response.Email)
		assert.Equal(t, int64(1234567890), response.Phone)
		assert.Equal(t, "123 Main St", response.Address)
		assert.Equal(t, "1990-01-01", response.DOB)
		assert.Equal(t, 1, *response.DeptID)
		assert.Equal(t, 2, *response.ManagerID)
	})

	t.Run("invalid input", func(t *testing.T) {
		newHR := Employee.HR{
			// Missing required fields
		}

		jsonValue, _ := json.Marshal(newHR)
		req, _ := http.NewRequest(http.MethodPost, "/hr", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/hr", repo.CreateHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid input")
	})

	t.Run("transaction begin error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(sql.ErrConnDone)

		deptID := 1
		managerID := 2
		newHR := Employee.HR{
			Employee: Employee.Employee{
				Name:      "John Doe",
				Email:     "john@example.com",
				Phone:     1234567890,
				Address:   "123 Main St",
				DOB:       "1990-01-01",
				DeptID:    &deptID,
				ManagerID: &managerID,
			},
		}

		jsonValue, _ := json.Marshal(newHR)
		req, _ := http.NewRequest(http.MethodPost, "/hr", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/hr", repo.CreateHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "Failed to start transaction")
	})

	t.Run("employee insertion error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO employee").
			WithArgs("John Doe", "john@example.com", int64(1234567890), "123 Main St", "1990-01-01", 1, 2).
			WillReturnError(sql.ErrConnDone)
		mock.ExpectRollback()

		deptID := 1
		managerID := 2
		newHR := Employee.HR{
			Employee: Employee.Employee{
				Name:      "John Doe",
				Email:     "john@example.com",
				Phone:     1234567890,
				Address:   "123 Main St",
				DOB:       "1990-01-01",
				DeptID:    &deptID,
				ManagerID: &managerID,
			},
		}

		jsonValue, _ := json.Marshal(newHR)
		req, _ := http.NewRequest(http.MethodPost, "/hr", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/hr", repo.CreateHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "sql: connection is already closed")
	})

	t.Run("hr insertion error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO employee").
			WithArgs("John Doe", "john@example.com", int64(1234567890), "123 Main St", "1990-01-01", 1, 2).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("INSERT INTO hr").
			WithArgs(1).
			WillReturnError(sql.ErrConnDone)
		mock.ExpectRollback()

		deptID := 1
		managerID := 2
		newHR := Employee.HR{
			Employee: Employee.Employee{
				Name:      "John Doe",
				Email:     "john@example.com",
				Phone:     1234567890,
				Address:   "123 Main St",
				DOB:       "1990-01-01",
				DeptID:    &deptID,
				ManagerID: &managerID,
			},
		}

		jsonValue, _ := json.Marshal(newHR)
		req, _ := http.NewRequest(http.MethodPost, "/hr", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/hr", repo.CreateHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "sql: connection is already closed")
	})

	t.Run("commit error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO employee").
			WithArgs("John Doe", "john@example.com", int64(1234567890), "123 Main St", "1990-01-01", 1, 2).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("INSERT INTO hr").
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit().WillReturnError(sql.ErrConnDone)

		deptID := 1
		managerID := 2
		newHR := Employee.HR{
			Employee: Employee.Employee{
				Name:      "John Doe",
				Email:     "john@example.com",
				Phone:     1234567890,
				Address:   "123 Main St",
				DOB:       "1990-01-01",
				DeptID:    &deptID,
				ManagerID: &managerID,
			},
		}

		jsonValue, _ := json.Marshal(newHR)
		req, _ := http.NewRequest(http.MethodPost, "/hr", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/hr", repo.CreateHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "Failed to commit transaction")
	})

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
