package session_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/smakasaki/shortener/internal/common"
	"github.com/smakasaki/shortener/internal/session"
	"github.com/smakasaki/shortener/internal/session/mocks"
	usermocks "github.com/smakasaki/shortener/internal/user/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthMiddleware_CheckSession(t *testing.T) {
	mockSessionRepo := new(mocks.MockRepository)
	mockUserRepo := new(usermocks.MockRepository)

	sessionID := uuid.New()
	userID := uuid.New()
	validSession := &session.Session{
		ID:        sessionID,
		UserID:    userID,
		CreatedAt: time.Now(),
	}

	expiredSession := &session.Session{
		ID:        sessionID,
		UserID:    userID,
		CreatedAt: time.Now().Add(-25 * time.Hour),
	}

	testCases := []struct {
		name           string
		cookieValue    string
		mockSetup      func()
		expectedStatus int
		expectedBody   map[string]string
	}{
		{
			name:        "success",
			cookieValue: sessionID.String(),
			mockSetup: func() {
				mockSessionRepo.On("Get", mock.Anything, sessionID).Return(validSession, nil).Once()
			},
			expectedStatus: http.StatusOK,
			expectedBody:   nil,
		},
		{
			name:           "unauthorized - no cookie",
			cookieValue:    "",
			mockSetup:      func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedBody: map[string]string{
				"error": "Unauthorized",
			},
		},
		{
			name:           "unauthorized - invalid session ID",
			cookieValue:    "invalid-uuid",
			mockSetup:      func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedBody: map[string]string{
				"error": "Unauthorized",
			},
		},
		{
			name:        "unauthorized - session not found",
			cookieValue: sessionID.String(),
			mockSetup: func() {
				mockSessionRepo.On("Get", mock.Anything, sessionID).Return(nil, errors.New("session not found")).Once()
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody: map[string]string{
				"error": "Unauthorized",
			},
		},
		{
			name:        "unauthorized - session expired",
			cookieValue: sessionID.String(),
			mockSetup: func() {
				mockSessionRepo.On("Get", mock.Anything, sessionID).Return(expiredSession, nil).Once()
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody: map[string]string{
				"error": session.ErrSessionExpired.Error(),
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()

			if tt.cookieValue != "" {
				cookie := &http.Cookie{
					Name:  common.SessionCookieName,
					Value: tt.cookieValue,
				}
				req.AddCookie(cookie)
			}

			c := e.NewContext(req, rec)
			c.SetRequest(req)

			middleware := session.NewAuthMiddleware(mockSessionRepo, mockUserRepo)
			handler := middleware.CheckSession(func(c echo.Context) error {
				return c.JSON(http.StatusOK, map[string]string{"message": "success"})
			})

			tt.mockSetup()

			err := handler(c)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			if tt.expectedBody != nil {
				var responseBody map[string]string
				err := json.Unmarshal(rec.Body.Bytes(), &responseBody)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, responseBody)
			}

			mockSessionRepo.AssertExpectations(t)
		})
	}
}
