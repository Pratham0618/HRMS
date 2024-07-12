package Testify

import (
	"Project/Employee"
	"Project/main/funcs"
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateEmployee(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// Mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	// Initialize repository with mock DB
	repo := funcs.NewRepo(db)

	t.Run("successful update", func(t *testing.T) {
		// Mocking expected SQL behavior
		mock.ExpectExec("UPDATE EMPLOYEE").WithArgs("John Doe", "john.doe@example.com", int64(1234567890), "456 New St", "1990-01-01", sql.NullInt64{}, sql.NullInt64{}, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// Prepare updated employee data
		updatedEmployee := Employee.Employee{
			EmpID:     1,
			Name:      "John Doe",
			Email:     "john.doe@example.com",
			Phone:     int64(1234567890),
			Address:   "456 New St",
			DOB:       "1990-01-01",
			DeptID:    nil, // Pointer to nil for null value
			ManagerID: nil, // Pointer to nil for null value
		}

		// Mock HTTP request
		jsonValue, _ := json.Marshal(updatedEmployee)
		req, _ := http.NewRequest(http.MethodPut, "/employees/1", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Create Gin router and define route handler
		router := gin.Default()
		router.PUT("/employees/:id", repo.UpdateEmployee)
		router.ServeHTTP(w, req)

		// Verify HTTP status code
		assert.Equal(t, http.StatusOK, w.Code)

		// Verify response body
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse JSON response: %v", err)
		}

		// Adjust expected values to match the type and format of actual values
		expected := map[string]interface{}{
			"emp_id":     float64(updatedEmployee.EmpID),
			"name":       updatedEmployee.Name,
			"email":      updatedEmployee.Email,
			"phone":      float64(updatedEmployee.Phone), // Convert int64 to float64 for JSON comparison
			"address":    updatedEmployee.Address,
			"dob":        updatedEmployee.DOB,
			"dept_id":    nil, // Check for nil value
			"manager_id": nil, // Check for nil value
		}

		assert.Equal(t, expected, response)
	})

	t.Run("invalid employee ID", func(t *testing.T) {
		// Mocking the database (not used in this case)
		db, _, _ := sqlmock.New()

		// Create repository using NewRepo function from funcs package
		repo := funcs.NewRepo(db)

		// Creating a test router and request
		router := gin.Default()
		router.PUT("/employees/:id", repo.UpdateEmployee)

		// Prepare JSON payload
		updateData := map[string]interface{}{
			"Name":    "John Doe",
			"Email":   "john.doe@example.com",
			"Phone":   int64(1234567890),
			"Address": "456 New St",
		}
		jsonValue, _ := json.Marshal(updateData)
		req, _ := http.NewRequest(http.MethodPut, "/employees/invalid_id", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		// Serve the request
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Verify response status code
		assert.Equal(t, http.StatusBadRequest, w.Code)

		// Verify response body
		expected := `{"error":"Invalid employee ID"}`
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("database error", func(t *testing.T) {
		// Mocking expected SQL behavior with an error
		mock.ExpectExec("UPDATE EMPLOYEE").
			WithArgs("John Doe", "john.doe@example.com", int64(1234567890), "456 New St", "1990-01-01", nil, nil, 1).
			WillReturnError(errors.New("database error"))

		// Prepare updated employee data
		updatedEmployee := Employee.Employee{
			EmpID:     1,
			Name:      "John Doe",
			Email:     "john.doe@example.com",
			Phone:     int64(1234567890),
			Address:   "456 New St",
			DOB:       "1990-01-01",
			DeptID:    nil, // Pointer to nil for null value
			ManagerID: nil, // Pointer to nil for null value
		}

		// Mock HTTP request
		jsonValue, _ := json.Marshal(updatedEmployee)
		req, _ := http.NewRequest(http.MethodPut, "/employees/1", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Create Gin router and define route handler
		router := gin.Default()
		router.PUT("/employees/:id", repo.UpdateEmployee)
		router.ServeHTTP(w, req)

		// Verify HTTP status code for internal server error
		assert.Equal(t, http.StatusInternalServerError, w.Code)

		// Verify error message in response body
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse JSON response: %v", err)
		}

		expectedError := "Failed to update employee: database error"
		if errorMessage, ok := response["error"].(string); !ok || errorMessage != expectedError {
			t.Errorf("Expected error message '%s', got '%s'", expectedError, errorMessage)
		}
	})

	// Add more test cases for other scenarios if needed

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}

}
