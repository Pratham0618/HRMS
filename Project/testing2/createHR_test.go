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

func TestCreateHR(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("successful creation", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO employee").
			WithArgs(sqlmock.AnyArg(), "John HR", "john.hr@example.com", 1234567890, "123 HR St", "1990-01-01", 1, 2).
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectExec("INSERT INTO hr").
			WithArgs(sqlmock.AnyArg(), 1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		newHR := Employee.HR{
			Employee: Employee.Employee{
				Name:      "John HR",
				Email:     "john.hr@example.com",
				Phone:     1234567890,
				Address:   "123 HR St",
				DOB:       "1990-01-01",
				DeptID:    func(i int) *int { return &i }(1),
				ManagerID: func(i int) *int { return &i }(2),
			},
			HR_ID: 1,
		}

		jsonValue, _ := json.Marshal(newHR)
		req, _ := http.NewRequest(http.MethodPost, "/hrs", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/hrs", repo.CreateHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Contains(t, response, "id")
	})

	t.Run("invalid input", func(t *testing.T) {
		newHR := Employee.HR{
			Employee: Employee.Employee{
				Name: "John HR",
				// Missing required fields
			},
		}

		jsonValue, _ := json.Marshal(newHR)
		req, _ := http.NewRequest(http.MethodPost, "/hrs", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/hrs", repo.CreateHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		expected := `{"error":"Invalid input"}`
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("database error", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO employee").
			WithArgs(sqlmock.AnyArg(), "John HR", "john.hr@example.com", "1234567890", "123 HR St", "1990-01-01", 1, 2).
			WillReturnError(sql.ErrConnDone)

		newHR := Employee.HR{
			Employee: Employee.Employee{
				Name:      "John HR",
				Email:     "john.hr@example.com",
				Phone:     1234567890,
				Address:   "123 HR St",
				DOB:       "1990-01-01",
				DeptID:    func(i int) *int { return &i }(1),
				ManagerID: func(i int) *int { return &i }(2),
			},
			HR_ID: 1,
		}

		jsonValue, _ := json.Marshal(newHR)
		req, _ := http.NewRequest(http.MethodPost, "/hrs", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/hrs", repo.CreateHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		expected := `{"error":"driver: bad connection"}`
		assert.JSONEq(t, expected, w.Body.String())
	})
}
