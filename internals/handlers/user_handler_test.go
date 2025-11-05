package handlers

import (
	"base_crud_api/internals/models"
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TO implement datadriven

type MockService struct {
	Called bool
	Err    error
}

func (us *MockService) CreateUser(login models.LoginRequest) ([]byte, error) {
	us.Called = true
	if login.Username == "fail" {
		return nil, errors.New("Unique name issue, duplicate entry")
	}

	token := []byte(`"token":"tokensystem"`)

	if login.Username == "token" {
		return token, us.Err
	}
	return []byte(login.Username), us.Err
}

func (us *MockService) LoginUser(login models.LoginRequest) (string, error) {
	us.Called = true
	if login.Username == "fail" {
		return "", errors.New("Name not found in the database")
	}

	token := "tokensystem"

	if login.Username == "token" {
		return token, nil
	}
	return login.Username, nil
}

func TestUserHandler_LoginUser(t *testing.T) {
	tests := []struct {
		name       string
		body       []byte
		wantStatus int
		wantError  bool
	}{
		{
			name:       "Happy Path",
			body:       []byte(`{"username":"user1","password":"securepassword"}`),
			wantStatus: http.StatusOK,
			wantError:  false,
		},
		{
			name:       "No body",
			body:       nil, // simulate missing body
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
		{
			name:       "Invalid JSON",
			body:       []byte(`{invalid-json}`),
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			req := httptest.NewRequest("POST", "/login", bytes.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			uh := UserHandler{Service: &MockService{}}
			uh.LoginUser(c) // call your handler

			assert.Equal(t, tt.wantStatus, w.Code)
			if !tt.wantError {
				assert.Contains(t, w.Body.String(), "token")
			}

			if tt.wantError {
				// optionally assert error response body
				assert.Contains(t, w.Body.String(), "error")
				fmt.Println(w.Body.String())
			}
		})
	}
}

func TestUserHandler_SignUpUser(t *testing.T) {
	tests := []struct {
		name       string
		body       []byte
		wantStatus int
		wantError  bool
	}{
		{
			name:       "Happy Path",
			body:       []byte(`{"username":"user1","password":"securepassword"}`),
			wantStatus: http.StatusOK,
			wantError:  false,
		},
		{
			name:       "No body",
			body:       nil, // simulate missing body
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
		{
			name:       "Invalid JSON",
			body:       []byte(`{invalid-json}`),
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			req := httptest.NewRequest("POST", "/signup", bytes.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			uh := UserHandler{Service: &MockService{}}
			uh.SignUpUser(c) // call your handler

			assert.Equal(t, tt.wantStatus, w.Code)

			if tt.wantError {
				// optionally assert error response body
				assert.Contains(t, w.Body.String(), "error")
				fmt.Println(w.Body.String())
			}
		})
	}
}
