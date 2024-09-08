package user_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/smakasaki/shortener/internal/user"
	"github.com/smakasaki/shortener/internal/user/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestUserHandler_Create tests the Create method in userHandler
func TestUserHandler_Create(t *testing.T) {
	// Mock the UseCase interface
	mockUseCase := new(mocks.MockUseCase)

	e := echo.New()

	tests := []struct {
		name           string
		inputBody      string
		mockBehavior   func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:      "Success",
			inputBody: `{"email": "test@example.com", "password": "password123"}`,
			mockBehavior: func() {
				mockUseCase.On("CreateUser", mock.Anything, mock.Anything).Return(nil).Once()
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   "User created",
		},
		{
			name:      "Validation Error - Missing Fields",
			inputBody: `{"email": "", "password": ""}`,
			mockBehavior: func() {
				// No need to call CreateUser, because validation will fail
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"errors":["Email is required","Email is not a valid email","Password is required","Password must be between 6 and 50 characters"]}`,
		},
		{
			name:      "UseCase Error",
			inputBody: `{"email": "test@example.com", "password": "password123"}`,
			mockBehavior: func() {
				mockUseCase.On("CreateUser", mock.Anything, mock.Anything).Return(assert.AnError).Once()
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"Could not create user"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a request with the test input
			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(tt.inputBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			// Create the echo context
			c := e.NewContext(req, rec)

			// Set up the handler
			h := user.NewUserHandler(mockUseCase)

			// Execute the mock behavior for each test
			tt.mockBehavior()

			// Call the Create handler
			err := h.Create(c)
			if err != nil {
				t.Errorf("handler error: %v", err)
			}

			// Assert the status code
			assert.Equal(t, tt.expectedStatus, rec.Code)

			// Remove any trailing newlines from the response body before asserting
			actualBody := strings.TrimSpace(rec.Body.String())

			// Assert the response body
			assert.Equal(t, tt.expectedBody, actualBody)

			// Assert that all expectations were met
			mockUseCase.AssertExpectations(t)
		})
	}
}
