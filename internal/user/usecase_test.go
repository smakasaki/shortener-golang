package user_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/smakasaki/shortener/internal/user"
	"github.com/smakasaki/shortener/internal/user/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateUser(t *testing.T) {
	// Create context
	ctx := context.Background()

	// Table-driven tests
	tests := []struct {
		name        string
		input       *user.User
		mockSetup   func(repo *mocks.MockRepository)
		expectedErr error
	}{
		{
			name: "successfully creates user",
			input: &user.User{
				Email:    "test@example.com",
				Password: "password123",
			},
			mockSetup: func(repo *mocks.MockRepository) {
				// Return nil, indicating that the user is not found (unique email)
				repo.EXPECT().GetUserByEmail(ctx, "test@example.com").Return(nil, nil)
				// Expect successful user creation
				repo.EXPECT().CreateUser(ctx, mock.AnythingOfType("*user.User")).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name: "fails when user already exists",
			input: &user.User{
				Email:    "test@example.com",
				Password: "password123",
			},
			mockSetup: func(repo *mocks.MockRepository) {
				// Return an existing user to simulate duplicate email
				repo.EXPECT().GetUserByEmail(ctx, "test@example.com").Return(&user.User{}, nil)
			},
			expectedErr: user.ErrUserAlreadyExists,
		},
		{
			name: "fails when hashing password",
			input: &user.User{
				Email:    "test@example.com",
				Password: string(make([]byte, 74)),
			},
			mockSetup: func(repo *mocks.MockRepository) {
				// Do not expect CreateUser to be called as the error should occur during password hashing
				repo.EXPECT().GetUserByEmail(ctx, "test@example.com").Return(nil, nil)
			},
			expectedErr: bcrypt.ErrPasswordTooLong,
		},
	}

	// Run each test
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock repository
			repo := new(mocks.MockRepository)
			tt.mockSetup(repo)

			// Create use case with mock repository
			uc := user.NewUseCase(repo)

			// Call the method and check for errors
			err := uc.CreateUser(ctx, tt.input)
			assert.Equal(t, tt.expectedErr, err)

			if tt.name == "fails when hashing password" {
				repo.AssertNotCalled(t, "CreateUser")
			} else {
				repo.AssertExpectations(t)
			}
		})
	}
}

func TestGetUserByID(t *testing.T) {
	// Create context
	ctx := context.Background()

	// Random UUID for testing
	testID := uuid.New()

	// Table-driven tests
	tests := []struct {
		name        string
		inputID     uuid.UUID
		mockSetup   func(repo *mocks.MockRepository)
		expectedRes *user.User
		expectedErr error
	}{
		{
			name:    "successfully retrieves user by ID",
			inputID: testID,
			mockSetup: func(repo *mocks.MockRepository) {
				// Return the user
				repo.EXPECT().GetUserByID(ctx, testID).Return(&user.User{
					ID:    testID,
					Email: "test@example.com",
				}, nil)
			},
			expectedRes: &user.User{
				ID:    testID,
				Email: "test@example.com",
			},
			expectedErr: nil,
		},
		{
			name:    "fails when user not found",
			inputID: testID,
			mockSetup: func(repo *mocks.MockRepository) {
				// Return nil to simulate user not found
				repo.EXPECT().GetUserByID(ctx, testID).Return(nil, nil)
			},
			expectedRes: nil,
			expectedErr: nil, // Expect nil as it is not an error if the user is not found
		},
		{
			name:    "fails with database error",
			inputID: testID,
			mockSetup: func(repo *mocks.MockRepository) {
				// Return a database error
				repo.EXPECT().GetUserByID(ctx, testID).Return(nil, errors.New("database error"))
			},
			expectedRes: nil,
			expectedErr: errors.New("database error"),
		},
	}

	// Run each test
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock repository
			repo := new(mocks.MockRepository)
			tt.mockSetup(repo)

			// Create use case with mock repository
			uc := user.NewUseCase(repo)

			// Call the method and check the result
			result, err := uc.GetUserByID(ctx, tt.inputID)
			assert.Equal(t, tt.expectedRes, result)
			assert.Equal(t, tt.expectedErr, err)

			// Ensure all mock expectations are met
			repo.AssertExpectations(t)
		})
	}
}
