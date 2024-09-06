package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/smakasaki/typing-trainer/domain"
)

type UseCase interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
}
