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

func TestGetDepartmentByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("successful retrieval", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"DEPT_ID", "DEPT_NAME"}).
			AddRow(1, "Human Resources")

		mock.ExpectQuery("SELECT DEPT_ID, DEPT_NAME FROM department WHERE DEPT_ID = ?").
			WithArgs(1).
			WillReturnRows(rows)

		req, _ := http.NewRequest(http.MethodGet, "/departments/1", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/departments/:id", repo.GetDepartmentByID)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		expected := `{"dept_id":1,"Dept_Name":"Human Resources"}`
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("department not found", func(t *testing.T) {
		mock.ExpectQuery("SELECT DEPT_ID, DEPT_NAME FROM department WHERE DEPT_ID = ?").
			WithArgs(999).
			WillReturnError(sql.ErrNoRows)

		req, _ := http.NewRequest(http.MethodGet, "/departments/999", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/departments/:id", repo.GetDepartmentByID)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		expected := `{"error":"Department not found"}`
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("invalid department id", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/departments/invalid", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/departments/:id", repo.GetDepartmentByID)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		expected := `{"error":"Invalid department ID"}`
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("database error", func(t *testing.T) {
		mock.ExpectQuery("SELECT DEPT_ID, DEPT_NAME FROM department WHERE DEPT_ID = ?").
			WithArgs(1).
			WillReturnError(sql.ErrConnDone)

		req, _ := http.NewRequest(http.MethodGet, "/departments/1", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/departments/:id", repo.GetDepartmentByID)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		expected := `{"error":"driver: bad connection"}`
		assert.JSONEq(t, expected, w.Body.String())
	})
}
