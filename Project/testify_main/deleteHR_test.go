package Testify

import (
	"Project/main/funcs"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteHR(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Initialize mock database and repository
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := funcs.NewRepo(db)

	t.Run("successful deletion", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM hr WHERE HR_ID = ?").
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectExec("DELETE FROM employee WHERE EMPLOYEE_ID = \\(SELECT EMPLOYEE_ID FROM hr WHERE HR_ID = \\?\\)").
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		req, _ := http.NewRequest(http.MethodDelete, "/hr/1", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.DELETE("/hr/:hrId", repo.DeleteHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "HR deleted successfully")
	})

	t.Run("invalid HR ID", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/hr/invalid", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.DELETE("/hr/:hrId", repo.DeleteHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid HR ID")
	})

	t.Run("HR not found", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM hr WHERE HR_ID = ?").
			WithArgs(999).
			WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectRollback()

		req, _ := http.NewRequest(http.MethodDelete, "/hr/999", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.DELETE("/hr/:hrId", repo.DeleteHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), "HR not found")
	})

	t.Run("transaction begin error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(sql.ErrConnDone)

		req, _ := http.NewRequest(http.MethodDelete, "/hr/1", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.DELETE("/hr/:hrId", repo.DeleteHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "Failed to start transaction")
	})

	t.Run("hr deletion error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM hr WHERE HR_ID = ?").
			WithArgs(1).
			WillReturnError(sql.ErrConnDone)
		mock.ExpectRollback()

		req, _ := http.NewRequest(http.MethodDelete, "/hr/1", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.DELETE("/hr/:hrId", repo.DeleteHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "sql: connection is already closed")
	})

	t.Run("employee deletion error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM hr WHERE HR_ID = ?").
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectExec("DELETE FROM employee WHERE EMPLOYEE_ID = \\(SELECT EMPLOYEE_ID FROM hr WHERE HR_ID = \\?\\)").
			WithArgs(1).
			WillReturnError(sql.ErrConnDone)
		mock.ExpectRollback()

		req, _ := http.NewRequest(http.MethodDelete, "/hr/1", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.DELETE("/hr/:hrId", repo.DeleteHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "sql: connection is already closed")
	})

	t.Run("commit error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM hr WHERE HR_ID = ?").
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectExec("DELETE FROM employee WHERE EMPLOYEE_ID = \\(SELECT EMPLOYEE_ID FROM hr WHERE HR_ID = \\?\\)").
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit().WillReturnError(sql.ErrConnDone)

		req, _ := http.NewRequest(http.MethodDelete, "/hr/1", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.DELETE("/hr/:hrId", repo.DeleteHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "Failed to commit transaction")
	})

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
