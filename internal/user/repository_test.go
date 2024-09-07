package user_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/smakasaki/shortener/internal/user"
	"github.com/stretchr/testify/assert"
)

func TestRepoGetUserByID(t *testing.T) {
	// Create a context
	ctx := context.Background()

	// Table-driven tests
	tests := []struct {
		name         string
		setupMock    func(mock sqlmock.Sqlmock, userID uuid.UUID)
		inputID      uuid.UUID
		expectedErr  error
		expectedUser *user.User
	}{
		{
			name: "successfully retrieves user by ID",
			setupMock: func(mock sqlmock.Sqlmock, userID uuid.UUID) {
				query := regexp.QuoteMeta("SELECT id, email, created_at FROM users WHERE id = $1")
				rows := sqlmock.NewRows([]string{"id", "email", "created_at"}).
					AddRow(userID, "test@example.com", time.Now())
				mock.ExpectQuery(query).WithArgs(userID).WillReturnRows(rows)
			},
			inputID:     uuid.New(),
			expectedErr: nil,
			expectedUser: &user.User{
				Email: "test@example.com",
			},
		},
		{
			name: "returns error when user not found",
			setupMock: func(mock sqlmock.Sqlmock, userID uuid.UUID) {
				query := regexp.QuoteMeta("SELECT id, email, created_at FROM users WHERE id = $1")
				mock.ExpectQuery(query).WithArgs(userID).WillReturnError(sql.ErrNoRows)
			},
			inputID:      uuid.New(),
			expectedErr:  user.ErrUserNotFound,
			expectedUser: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a database mock and sqlmock
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			// Create a repository
			repo := user.NewRepository(db)

			// Set up the mock
			tt.setupMock(mock, tt.inputID)

			// Call the method
			result, err := repo.GetUserByID(ctx, tt.inputID)

			// Check the result
			assert.Equal(t, tt.expectedErr, err)
			if tt.expectedUser != nil {
				assert.NotNil(t, result)
				assert.Equal(t, tt.expectedUser.Email, result.Email)
			} else {
				assert.Nil(t, result)
			}

			// Make sure all expectations were met
			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}

func TestRepoCreateUser(t *testing.T) {
	// Create a context
	ctx := context.Background()

	// Table-driven tests
	tests := []struct {
		name        string
		inputUser   *user.User
		setupMock   func(mock sqlmock.Sqlmock, u *user.User)
		expectedErr error
	}{
		{
			name: "successfully creates user",
			inputUser: &user.User{
				Email:    "test@example.com",
				Password: "hashedpassword",
			},
			setupMock: func(mock sqlmock.Sqlmock, u *user.User) {
				query := regexp.QuoteMeta(`
					INSERT INTO users (email, password, created_at, updated_at)
					VALUES ($1, $2, $3, $4)
				`)
				mock.ExpectExec(query).
					WithArgs(u.Email, u.Password, sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedErr: nil,
		},
		{
			name: "fails when user already exists",
			inputUser: &user.User{
				Email:    "test@example.com",
				Password: "hashedpassword",
			},
			setupMock: func(mock sqlmock.Sqlmock, u *user.User) {
				query := regexp.QuoteMeta(`
					INSERT INTO users (email, password, created_at, updated_at)
					VALUES ($1, $2, $3, $4)
				`)
				mock.ExpectExec(query).
					WithArgs(u.Email, u.Password, sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(user.ErrUserAlreadyExists)
			},
			expectedErr: user.ErrUserAlreadyExists,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a database mock and sqlmock
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			// Create a repository
			repo := user.NewRepository(db)

			// Set up the mock
			tt.setupMock(mock, tt.inputUser)

			// Call the method
			err = repo.CreateUser(ctx, tt.inputUser)

			// Check the result
			assert.Equal(t, tt.expectedErr, err)

			// Make sure all expectations were met
			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}

func TestRepoGetUserByEmail(t *testing.T) {
	// Create a context
	ctx := context.Background()

	// Table-driven tests
	tests := []struct {
		name         string
		inputEmail   string
		setupMock    func(mock sqlmock.Sqlmock, email string)
		expectedErr  error
		expectedUser *user.User
	}{
		{
			name:       "successfully retrieves user by email",
			inputEmail: "test@example.com",
			setupMock: func(mock sqlmock.Sqlmock, email string) {
				query := regexp.QuoteMeta("SELECT id, email, password, created_at, updated_at FROM users WHERE email = $1")
				userID := uuid.New()
				now := time.Now()
				rows := sqlmock.NewRows([]string{"id", "email", "password", "created_at", "updated_at"}).
					AddRow(userID, email, "hashedpassword", now, now)
				mock.ExpectQuery(query).WithArgs(email).WillReturnRows(rows)
			},
			expectedErr: nil,
			expectedUser: &user.User{
				Email: "test@example.com",
			},
		},
		{
			name:       "returns error when user not found",
			inputEmail: "unknown@example.com",
			setupMock: func(mock sqlmock.Sqlmock, email string) {
				query := regexp.QuoteMeta("SELECT id, email, password, created_at, updated_at FROM users WHERE email = $1")
				mock.ExpectQuery(query).WithArgs(email).WillReturnError(sql.ErrNoRows)
			},
			expectedErr:  user.ErrUserNotFound,
			expectedUser: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a database mock and sqlmock
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			// Create a repository
			repo := user.NewRepository(db)

			// Set up the mock
			tt.setupMock(mock, tt.inputEmail)

			// Call the method
			result, err := repo.GetUserByEmail(ctx, tt.inputEmail)

			// Check the result
			assert.Equal(t, tt.expectedErr, err)
			if tt.expectedUser != nil {
				assert.NotNil(t, result)
				assert.Equal(t, tt.expectedUser.Email, result.Email)
			} else {
				assert.Nil(t, result)
			}

			// Make sure all expectations were met
			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}
