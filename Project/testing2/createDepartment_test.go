package Testify

import (
	"Project/Employee"
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateDepartment(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("successful creation", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO department").
			WithArgs("Engineering").
			WillReturnResult(sqlmock.NewResult(1, 1))

		newDept := Employee.Department{
			Dept_ID:   1, // Ensure Dept_ID is initialized
			Dept_Name: "Engineering",
		}

		jsonValue, _ := json.Marshal(newDept)
		req, _ := http.NewRequest(http.MethodPost, "/departments", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/departments", repo.CreateDepartment)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		expected := `{"dept_id":1,"Dept_Name":"Engineering"}`
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("invalid input", func(t *testing.T) {
		newDept := Employee.Department{
			// Missing Dept_Name
		}

		jsonValue, _ := json.Marshal(newDept)
		req, _ := http.NewRequest(http.MethodPost, "/departments", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/departments", repo.CreateDepartment)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		expected := `{"error":"Invalid input: Department name is required"}`
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("database error", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO department").
			WithArgs("Finance").
			WillReturnError(sql.ErrConnDone)

		newDept := Employee.Department{
			Dept_ID:   1, // Ensure Dept_ID is initialized
			Dept_Name: "Finance",
		}

		jsonValue, _ := json.Marshal(newDept)
		req, _ := http.NewRequest(http.MethodPost, "/departments", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/departments", repo.CreateDepartment)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		expected := `{"error":"Failed to create department: driver: bad connection"}`
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("duplicate department name", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO department").
			WithArgs("Human Resources").
			WillReturnError(&mysql.MySQLError{Number: 1062, Message: "Duplicate entry 'Human Resources' for key 'dept_name'"})

		newDept := Employee.Department{
			Dept_ID:   1, // Ensure Dept_ID is initialized
			Dept_Name: "Human Resources",
		}

		jsonValue, _ := json.Marshal(newDept)
		req, _ := http.NewRequest(http.MethodPost, "/departments", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/departments", repo.CreateDepartment)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusConflict, w.Code)
		expected := `{"error":"Department name already exists"}`
		assert.JSONEq(t, expected, w.Body.String())
	})
}
