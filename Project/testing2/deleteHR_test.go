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

func TestDeleteHR(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("successful deletion", func(t *testing.T) {
		mock.ExpectExec("DELETE e, h FROM employee e JOIN hr h ON e.EMPLOYEE_ID = h.EMPLOYEE_ID WHERE h.HR_ID = ?").
			WithArgs(101).
			WillReturnResult(sqlmock.NewResult(0, 1))

		req, _ := http.NewRequest(http.MethodDelete, "/hrs/101", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.DELETE("/hrs/:hrId", repo.DeleteHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		expected := `{"message":"HR deleted successfully"}`
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("hr not found", func(t *testing.T) {
		mock.ExpectExec("DELETE e, h FROM employee e JOIN hr h ON e.EMPLOYEE_ID = h.EMPLOYEE_ID WHERE h.HR_ID = ?").
			WithArgs(999).
			WillReturnResult(sqlmock.NewResult(0, 0))

		req, _ := http.NewRequest(http.MethodDelete, "/hrs/999", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.DELETE("/hrs/:hrId", repo.DeleteHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		expected := `{"error":"HR not found"}`
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("invalid hr id", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/hrs/invalid", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.DELETE("/hrs/:hrId", repo.DeleteHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		expected := `{"error":"Invalid HR ID"}`
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("database error", func(t *testing.T) {
		mock.ExpectExec("DELETE e, h FROM employee e JOIN hr h ON e.EMPLOYEE_ID = h.EMPLOYEE_ID WHERE h.HR_ID = ?").
			WithArgs(101).
			WillReturnError(sql.ErrConnDone)

		req, _ := http.NewRequest(http.MethodDelete, "/hrs/101", nil)
		w := httptest.NewRecorder()

		router := gin.Default()
		router.DELETE("/hrs/:hrId", repo.DeleteHR)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		expected := `{"error":"driver: bad connection"}`
		assert.JSONEq(t, expected, w.Body.String())
	})
}
