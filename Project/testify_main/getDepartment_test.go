package Testify

import (
	"Project/Employee"
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

func TestGetDepartments(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Initialize mock database and repository
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := funcs.NewRepo(db)

	t.Run("successful retrieval", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"DEPT_ID", "DEPT_NAME"}).
			AddRow(1, "HR").
			AddRow(2, "Engineering").
			AddRow(3, "Marketing")

		mock.ExpectQuery("SELECT DEPT_ID, DEPT_NAME FROM department").
			WillReturnRows(rows)

		req, _ := http.NewRequest(http.MethodGet, "/departments", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/departments", repo.GetDepartments)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var departments []Employee.Department
		err := json.Unmarshal(w.Body.Bytes(), &departments)
		assert.NoError(t, err)

		assert.Len(t, departments, 3)
		assert.Equal(t, 1, departments[0].Dept_ID)
		assert.Equal(t, "HR", departments[0].Dept_Name)
		assert.Equal(t, 2, departments[1].Dept_ID)
		assert.Equal(t, "Engineering", departments[1].Dept_Name)
		assert.Equal(t, 3, departments[2].Dept_ID)
		assert.Equal(t, "Marketing", departments[2].Dept_Name)
	})

	t.Run("empty result", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"DEPT_ID", "DEPT_NAME"})

		mock.ExpectQuery("SELECT DEPT_ID, DEPT_NAME FROM department").
			WillReturnRows(rows)

		req, _ := http.NewRequest(http.MethodGet, "/departments", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/departments", repo.GetDepartments)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var departments []Employee.Department
		err := json.Unmarshal(w.Body.Bytes(), &departments)
		assert.NoError(t, err)

		assert.Len(t, departments, 0)
	})

	t.Run("database error", func(t *testing.T) {
		mock.ExpectQuery("SELECT DEPT_ID, DEPT_NAME FROM department").
			WillReturnError(sql.ErrConnDone)

		req, _ := http.NewRequest(http.MethodGet, "/departments", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/departments", repo.GetDepartments)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "sql: connection is already closed")
	})

	t.Run("scan error", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"DEPT_ID", "DEPT_NAME"}).
			AddRow("invalid", "HR") // This will cause a scan error

		mock.ExpectQuery("SELECT DEPT_ID, DEPT_NAME FROM department").
			WillReturnRows(rows)

		req, _ := http.NewRequest(http.MethodGet, "/departments", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/departments", repo.GetDepartments)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "sql: Scan error")
	})

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
