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

func TestDeleteDepartment(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("successful deletion", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec("DELETE FROM department WHERE dept_id = ?").
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		req, _ := http.NewRequest(http.MethodDelete, "/departments/1", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.DELETE("/departments/:id", repo.DeleteDepartment)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		expected := `{"message":"Department deleted successfully"}`
		assert.JSONEq(t, expected, w.Body.String())

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("department not found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec("DELETE FROM department WHERE dept_id = ?").
			WithArgs(999).
			WillReturnResult(sqlmock.NewResult(0, 0))

		req, _ := http.NewRequest(http.MethodDelete, "/departments/999", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.DELETE("/departments/:id", repo.DeleteDepartment)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		expected := `{"error":"Department not found"}`
		assert.JSONEq(t, expected, w.Body.String())

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("database error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec("DELETE FROM department WHERE dept_id = ?").
			WithArgs(1).
			WillReturnError(sql.ErrConnDone)

		req, _ := http.NewRequest(http.MethodDelete, "/departments/1", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.DELETE("/departments/:id", repo.DeleteDepartment)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		expected := `{"error":"Failed to delete department: driver: bad connection"}`
		assert.JSONEq(t, expected, w.Body.String())

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
