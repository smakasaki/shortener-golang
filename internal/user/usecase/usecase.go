package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/smakasaki/typing-trainer/domain"
	"github.com/smakasaki/typing-trainer/internal/user"
)

type UserUseCase struct {
	userRepository user.Repository
}

func NewUserUseCase(userRepository user.Repository) user.UseCase {
	return &UserUseCase{
		userRepository: userRepository,
	}
}

func (u *UserUseCase) CreateUser(ctx context.Context, user *domain.User) error {
	return nil
}

func (u *UserUseCase) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return nil, nil
}
