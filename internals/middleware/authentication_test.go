package middleware

import (
	"base_crud_api/internals/models"
	"base_crud_api/internals/utils"
	"fmt"
	"net/http/httptest"
	"os"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Printf("No able to load .env file : %v\n", err.Error())
	}
	code := m.Run()
	os.Exit(code)

}

func TestAuthentication(t *testing.T) {
	username := "golden_user"
	token, err := utils.CreateToken(username)
	if err != nil {
		t.Fatal("failed to create token:", err)
	}
	bearerToken := fmt.Sprintf("Bearer %s", token)

	tests := []struct {
		name         string
		authHeader   string
		wantUsername string
		wantExists   bool
	}{
		{
			name:         "valid token",
			authHeader:   bearerToken,
			wantUsername: username,
			wantExists:   true,
		},
		{
			name:         "missing token",
			authHeader:   "",
			wantUsername: "",
			wantExists:   false,
		},
		{
			name:         "invalid token",
			authHeader:   "Bearer invalidtoken",
			wantUsername: "",
			wantExists:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			req := httptest.NewRequest("GET", "/protected", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}
			c.Request = req

			Authentication(c)

			val, exists := c.Get("claims")
			assert.Equal(t, tt.wantExists, exists)

			if tt.wantExists {
				claims := val.(models.Claims)
				assert.Equal(t, tt.wantUsername, claims.Username)
			}
		})
	}
}
