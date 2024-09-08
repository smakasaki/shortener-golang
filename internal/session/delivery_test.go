package session_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/smakasaki/shortener/internal/common"
	"github.com/smakasaki/shortener/internal/session"
	"github.com/smakasaki/shortener/internal/session/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSessionHandler_Create(t *testing.T) {
	mockUseCase := new(mocks.MockUseCase)

	// Тестовые данные
	email := "test@example.com"
	password := "password123"
	sessionID := uuid.New()

	testCases := []struct {
		name           string
		input          map[string]string
		mockSetup      func()
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "success",
			input: map[string]string{
				"email":    email,
				"password": password,
			},
			mockSetup: func() {
				mockUseCase.On("Create", mock.Anything, email, password).Return(sessionID, nil).Once()
			},
			expectedStatus: http.StatusCreated,
			expectedBody: map[string]interface{}{
				"message": "Session created",
			},
		},
		{
			name: "invalid input",
			input: map[string]string{
				"email":    "invalid-email",
				"password": "123",
			},
			mockSetup:      func() {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"errors": []interface{}{
					"Email is not a valid email",
					"Password must be between 6 and 50 characters",
				},
			},
		},
		{
			name: "invalid credentials",
			input: map[string]string{
				"email":    email,
				"password": "wrongpassword",
			},
			mockSetup: func() {
				mockUseCase.On("Create", mock.Anything, email, "wrongpassword").Return(uuid.Nil, session.ErrInvalidCredentials).Once()
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "invalid credentials",
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			h := session.NewHandler(mockUseCase)

			body, _ := json.Marshal(tt.input)
			req := httptest.NewRequest(http.MethodPost, "/sessions", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			tt.mockSetup()

			err := h.Create(c)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			var responseBody map[string]interface{}
			json.Unmarshal(rec.Body.Bytes(), &responseBody)
			assert.Equal(t, tt.expectedBody, responseBody)

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestSessionHandler_Delete(t *testing.T) {
	mockUseCase := new(mocks.MockUseCase)

	sessionID := uuid.New()
	userID := uuid.New()

	testCases := []struct {
		name           string
		mockSetup      func()
		session        *common.Session
		expectedStatus int
		expectedBody   map[string]string
	}{
		{
			name: "success",
			mockSetup: func() {
				mockUseCase.On("Delete", mock.Anything, sessionID, userID).Return(nil).Once()
			},
			session:        &common.Session{ID: sessionID, UserID: userID},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]string{
				"message": "Session deleted",
			},
		},
		{
			name: "session not found",
			mockSetup: func() {
				mockUseCase.On("Delete", mock.Anything, sessionID, userID).Return(session.ErrSessionNotFound).Once()
			},
			session:        &common.Session{ID: sessionID, UserID: userID},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]string{
				"error": "session not found",
			},
		},
		{
			name: "delete error",
			mockSetup: func() {
				mockUseCase.On("Delete", mock.Anything, sessionID, userID).Return(errors.New("delete error")).Once()
			},
			session:        &common.Session{ID: sessionID, UserID: userID},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]string{
				"error": "delete error",
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			h := session.NewHandler(mockUseCase)

			req := httptest.NewRequest(http.MethodDelete, "/sessions", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.Set(common.SessionEchoStorageKey, tt.session)

			tt.mockSetup()

			err := h.Delete(c)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			var responseBody map[string]string
			json.Unmarshal(rec.Body.Bytes(), &responseBody)
			assert.Equal(t, tt.expectedBody, responseBody)

			mockUseCase.AssertExpectations(t)
		})
	}
}
