package Testify

import (
	"Project/main/funcs"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteEmployee(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	// Initialize repository with mock DB
	repo := funcs.NewRepo(db)

	t.Run("valid_employee_id", func(t *testing.T) {
		// Mocking expected SQL behavior for a valid employee ID
		id := 1
		mock.ExpectExec("DELETE FROM EMPLOYEE").
			WithArgs(id).
			WillReturnResult(sqlmock.NewResult(0, 1)) // Expect 1 row affected

		// Mock HTTP request
		req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/employees/%d", id), nil)
		w := httptest.NewRecorder()

		// Create Gin router and define route handler
		router := gin.Default()
		router.DELETE("/employees/:id", repo.DeleteEmployee)
		router.ServeHTTP(w, req)

		// Verify HTTP status code for success
		assert.Equal(t, http.StatusOK, w.Code)

		// Verify response body
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse JSON response: %v", err)
		}

		expectedMessage := "Employee deleted"
		if message, ok := response["message"].(string); !ok || message != expectedMessage {
			t.Errorf("Expected message '%s', got '%s'", expectedMessage, message)
		}
	})

	t.Run("invalid_employee_id", func(t *testing.T) {
		// Mocking expected SQL behavior for an invalid employee ID
		id := 999 // Assuming employee ID 999 does not exist
		mock.ExpectExec("DELETE FROM EMPLOYEE").
			WithArgs(id).
			WillReturnError(errors.New("employee not found")) // Simulate error for non-existent employee

		// Mock HTTP request
		req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/employees/%d", id), nil)
		w := httptest.NewRecorder()

		// Create Gin router and define route handler
		router := gin.Default()
		router.DELETE("/employees/:id", repo.DeleteEmployee)
		router.ServeHTTP(w, req)

		// Verify HTTP status code for internal server error
		assert.Equal(t, http.StatusInternalServerError, w.Code)

		// Verify error message in response body
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse JSON response: %v", err)
		}

		expectedError := "employee not found"
		if errorMessage, ok := response["error"].(string); !ok || errorMessage != expectedError {
			t.Errorf("Expected error message '%s', got '%s'", expectedError, errorMessage)
		}
	})

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
