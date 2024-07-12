package Testify

import (
	"Project/main/funcs"
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
	// Set up the mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Create a new repo instance with the mock db
	repo := funcs.NewRepo(db)

	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	t.Run("successful creation", func(t *testing.T) {
		department := map[string]string{
			"name": "New Department",
		}
		jsonValue, _ := json.Marshal(department)

		mock.ExpectExec("INSERT INTO department").
			WithArgs(department["name"]).
			WillReturnResult(sqlmock.NewResult(1, 1))

		req, _ := http.NewRequest(http.MethodPost, "/departments", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/departments", repo.CreateDepartment)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		expected := `{"message":"Department created successfully","id":1}`
		assert.JSONEq(t, expected, w.Body.String())

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}
	})

	t.Run("duplicate department name", func(t *testing.T) {
		department := map[string]string{
			"name": "Existing Department",
		}
		jsonValue, _ := json.Marshal(department)

		mock.ExpectExec("INSERT INTO department").
			WithArgs(department["name"]).
			WillReturnError(&mysql.MySQLError{Number: 1062, Message: "Duplicate entry 'Existing Department' for key 'name'"})

		req, _ := http.NewRequest(http.MethodPost, "/departments", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/departments", repo.CreateDepartment)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusConflict, w.Code)
		expected := `{"error":"Department name already exists"}`
		assert.JSONEq(t, expected, w.Body.String())

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}
	})

	t.Run("database error", func(t *testing.T) {
		department := map[string]string{
			"name": "Error Department",
		}
		jsonValue, _ := json.Marshal(department)

		mock.ExpectExec("INSERT INTO department").
			WithArgs(department["name"]).
			WillReturnError(sql.ErrConnDone)

		req, _ := http.NewRequest(http.MethodPost, "/departments", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/departments", repo.CreateDepartment)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		expected := `{"error":"sql: connection is already closed"}`
		assert.JSONEq(t, expected, w.Body.String())

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}
	})
}
