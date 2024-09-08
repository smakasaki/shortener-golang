package session_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/smakasaki/shortener/internal/session"
	"github.com/stretchr/testify/assert"
)

func TestCreateSession(t *testing.T) {
	sessionID := uuid.New()
	userID := uuid.New()

	testCases := []struct {
		name          string
		userID        uuid.UUID
		mockSetup     func(sqlmock.Sqlmock)
		expectedID    uuid.UUID
		expectedError error
	}{
		{
			name:   "success",
			userID: userID,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("INSERT INTO sessions").
					WithArgs(userID).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(sessionID))
			},
			expectedID:    sessionID,
			expectedError: nil,
		},
		{
			name:   "db error",
			userID: userID,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("INSERT INTO sessions").
					WithArgs(userID).
					WillReturnError(errors.New("db error"))
			},
			expectedID:    uuid.Nil,
			expectedError: errors.New("db error"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			tt.mockSetup(mock)

			repo := session.NewRepository(db)
			id, err := repo.Create(context.Background(), tt.userID)

			assert.Equal(t, tt.expectedID, id)
			assert.Equal(t, tt.expectedError, err)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetSession(t *testing.T) {
	sessionID := uuid.New()
	userID := uuid.New()
	now := time.Now()

	testCases := []struct {
		name          string
		sessionID     uuid.UUID
		mockSetup     func(sqlmock.Sqlmock)
		expected      *session.Session
		expectedError error
	}{
		{
			name:      "success",
			sessionID: sessionID,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT .* FROM sessions WHERE id = \\$1").
					WithArgs(sessionID).
					WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "created_at"}).
						AddRow(sessionID, userID, now))
			},
			expected: &session.Session{
				ID:        sessionID,
				UserID:    userID,
				CreatedAt: now,
			},
			expectedError: nil,
		},
		{
			name:      "session not found",
			sessionID: sessionID,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT .* FROM sessions WHERE id = \\$1").
					WithArgs(sessionID).
					WillReturnError(sql.ErrNoRows)
			},
			expected:      nil,
			expectedError: session.ErrSessionNotFound,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			tt.mockSetup(mock)

			repo := session.NewRepository(db)
			sess, err := repo.Get(context.Background(), tt.sessionID)

			assert.Equal(t, tt.expected, sess)
			assert.Equal(t, tt.expectedError, err)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestDeleteSession(t *testing.T) {
	sessionID := uuid.New()

	testCases := []struct {
		name          string
		sessionID     uuid.UUID
		mockSetup     func(sqlmock.Sqlmock)
		expectedError error
	}{
		{
			name:      "success",
			sessionID: sessionID,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("DELETE FROM sessions WHERE id = \\$1").
					WithArgs(sessionID).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name:      "session not found",
			sessionID: sessionID,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("DELETE FROM sessions WHERE id = \\$1").
					WithArgs(sessionID).
					WillReturnError(sql.ErrNoRows)
			},
			expectedError: session.ErrSessionNotFound,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			tt.mockSetup(mock)

			repo := session.NewRepository(db)
			err = repo.Delete(context.Background(), tt.sessionID)

			assert.Equal(t, tt.expectedError, err)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
