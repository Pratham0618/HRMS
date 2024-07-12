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

func TestUpdateDepartment(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("successful update", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec("UPDATE department SET dept_name = ? WHERE dept_id = ?").
			WithArgs("Engineering Updated", 1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		updatedDept := Employee.Department{
			Dept_ID:   1,
			Dept_Name: "Engineering Updated",
		}

		jsonValue, _ := json.Marshal(updatedDept)
		req, _ := http.NewRequest(http.MethodPut, "/departments/1", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.PUT("/departments/:id", repo.UpdateDepartment)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		expected := `{"dept_id":1,"Dept_Name":"Engineering Updated"}`
		assert.JSONEq(t, expected, w.Body.String())

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("invalid input", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		updatedDept := Employee.Department{
			// Missing Dept_Name
		}

		jsonValue, _ := json.Marshal(updatedDept)
		req, _ := http.NewRequest(http.MethodPut, "/departments/1", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.PUT("/departments/:id", repo.UpdateDepartment)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		expected := `{"error":"Invalid input: Department name is required"}`
		assert.JSONEq(t, expected, w.Body.String())

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("database error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec("UPDATE department SET dept_name = ? WHERE dept_id = ?").
			WithArgs("Finance Updated", 1).
			WillReturnError(sql.ErrConnDone)

		updatedDept := Employee.Department{
			Dept_ID:   1,
			Dept_Name: "Finance Updated",
		}

		jsonValue, _ := json.Marshal(updatedDept)
		req, _ := http.NewRequest(http.MethodPut, "/departments/1", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.PUT("/departments/:id", repo.UpdateDepartment)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		expected := `{"error":"Failed to update department: driver: bad connection"}`
		assert.JSONEq(t, expected, w.Body.String())

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("department not found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec("UPDATE department SET dept_name = ? WHERE dept_id = ?").
			WithArgs("Non-existent Department", 999).
			WillReturnResult(sqlmock.NewResult(0, 0))

		updatedDept := Employee.Department{
			Dept_ID:   999,
			Dept_Name: "Non-existent Department",
		}

		jsonValue, _ := json.Marshal(updatedDept)
		req, _ := http.NewRequest(http.MethodPut, "/departments/999", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.PUT("/departments/:id", repo.UpdateDepartment)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		expected := `{"error":"Department not found"}`
		assert.JSONEq(t, expected, w.Body.String())

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("duplicate department name", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec("UPDATE department SET dept_name = ? WHERE dept_id = ?").
			WithArgs("Human Resources", 1).
			WillReturnError(&mysql.MySQLError{Number: 1062, Message: "Duplicate entry 'Human Resources' for key 'dept_name'"})

		updatedDept := Employee.Department{
			Dept_ID:   1,
			Dept_Name: "Human Resources",
		}

		jsonValue, _ := json.Marshal(updatedDept)
		req, _ := http.NewRequest(http.MethodPut, "/departments/1", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.PUT("/departments/:id", repo.UpdateDepartment)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusConflict, w.Code)
		expected := `{"error":"Department name already exists"}`
		assert.JSONEq(t, expected, w.Body.String())

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
