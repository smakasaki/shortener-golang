package session_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/smakasaki/shortener/internal/session"
	"github.com/smakasaki/shortener/internal/session/mocks"
	"github.com/smakasaki/shortener/internal/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestSessionUseCase_Create(t *testing.T) {
	userID := uuid.New()
	sessionID := uuid.New()
	password := "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	email := "test@example.com"

	mockRepo := new(mocks.MockRepository)
	mockUserRepo := new(mocks.MockUserRepository)

	testCases := []struct {
		name          string
		mockSetup     func()
		email         string
		password      string
		expectedID    uuid.UUID
		expectedError error
	}{
		{
			name: "success",
			mockSetup: func() {
				mockUser := &user.User{ID: userID, Email: email, Password: string(hashedPassword)}
				mockUserRepo.On("GetUserByEmail", mock.Anything, email).Return(mockUser, nil).Once()
				mockRepo.On("Create", mock.Anything, userID).Return(sessionID, nil).Once()
			},
			email:         email,
			password:      password,
			expectedID:    sessionID,
			expectedError: nil,
		},
		{
			name: "invalid password",
			mockSetup: func() {
				mockUser := &user.User{ID: userID, Email: email, Password: string(hashedPassword)}
				mockUserRepo.On("GetUserByEmail", mock.Anything, email).Return(mockUser, nil).Once()
			},
			email:         email,
			password:      "wrongpassword",
			expectedID:    uuid.Nil,
			expectedError: session.ErrInvalidCredentials,
		},
		{
			name: "user not found",
			mockSetup: func() {
				mockUserRepo.On("GetUserByEmail", mock.Anything, email).Return(nil, errors.New("user not found")).Once()
			},
			email:         email,
			password:      password,
			expectedID:    uuid.Nil,
			expectedError: session.ErrInvalidCredentials,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			uc := session.NewUseCase(mockRepo, mockUserRepo)
			id, err := uc.Create(context.Background(), tt.email, tt.password)

			assert.Equal(t, tt.expectedID, id)
			assert.Equal(t, tt.expectedError, err)

			mockRepo.AssertExpectations(t)
			mockUserRepo.AssertExpectations(t)
		})
	}
}

func TestSessionUseCase_Delete(t *testing.T) {
	sessionID := uuid.New()
	userID := uuid.New()

	mockRepo := new(mocks.MockRepository)
	mockUserRepo := new(mocks.MockUserRepository)

	testCases := []struct {
		name          string
		mockSetup     func()
		sessionID     uuid.UUID
		userID        uuid.UUID
		expectedError error
	}{
		{
			name: "success",
			mockSetup: func() {
				mockSession := &session.Session{ID: sessionID, UserID: userID}
				mockRepo.On("Get", mock.Anything, sessionID).Return(mockSession, nil).Once()
				mockRepo.On("Delete", mock.Anything, sessionID).Return(nil).Once()
			},
			sessionID:     sessionID,
			userID:        userID,
			expectedError: nil,
		},
		{
			name: "session not found",
			mockSetup: func() {
				mockRepo.On("Get", mock.Anything, sessionID).Return(nil, session.ErrSessionNotFound).Once()
			},
			sessionID:     sessionID,
			userID:        userID,
			expectedError: session.ErrSessionNotFound,
		},
		{
			name: "user ID mismatch",
			mockSetup: func() {
				mockSession := &session.Session{ID: sessionID, UserID: uuid.New()}
				mockRepo.On("Get", mock.Anything, sessionID).Return(mockSession, nil).Once()
			},
			sessionID:     sessionID,
			userID:        userID,
			expectedError: session.ErrSessionNotFound,
		},
		{
			name: "delete error",
			mockSetup: func() {
				mockSession := &session.Session{ID: sessionID, UserID: userID}
				mockRepo.On("Get", mock.Anything, sessionID).Return(mockSession, nil).Once()
				mockRepo.On("Delete", mock.Anything, sessionID).Return(errors.New("delete error")).Once()
			},
			sessionID:     sessionID,
			userID:        userID,
			expectedError: errors.New("delete error"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			uc := session.NewUseCase(mockRepo, mockUserRepo)
			err := uc.Delete(context.Background(), tt.sessionID, tt.userID)

			assert.Equal(t, tt.expectedError, err)

			mockRepo.AssertExpectations(t)
		})
	}
}
