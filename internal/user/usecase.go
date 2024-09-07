package user

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error) // For authentication
}

type userUseCase struct {
	userRepository Repository
}

func NewUseCase(userRepository Repository) *userUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (u *userUseCase) CreateUser(ctx context.Context, user *User) error {
	return nil
}

func (u *userUseCase) GetUserByID(ctx context.Context, id uuid.UUID) (*User, error) {
	return nil, nil
}
