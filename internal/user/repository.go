package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/smakasaki/typing-trainer/domain"
)

type Repository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error) // For authentication
}
