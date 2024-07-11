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

func TestGetDepartment(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("successful retrieval", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"DEPT_ID", "DEPT_NAME"}).
			AddRow(1, "Human Resources").
			AddRow(2, "Engineering")

		mock.ExpectQuery("SELECT DEPT_ID, DEPT_NAME FROM department").
			WillReturnRows(rows)

		req, _ := http.NewRequest(http.MethodGet, "/departments", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/departments", repo.GetDepartments)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		expected := `[{"dept_id":1,"Dept_Name":"Human Resources"},{"dept_id":2,"Dept_Name":"Engineering"}]`
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("no departments found", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"DEPT_ID", "DEPT_NAME"})

		mock.ExpectQuery("SELECT DEPT_ID, DEPT_NAME FROM department").
			WillReturnRows(rows)

		req, _ := http.NewRequest(http.MethodGet, "/departments", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/departments", repo.GetDepartments)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		expected := `[]`
		assert.JSONEq(t, expected, w.Body.String())
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
		expected := `{"error":"driver: bad connection"}`
		assert.JSONEq(t, expected, w.Body.String())
	})
}
