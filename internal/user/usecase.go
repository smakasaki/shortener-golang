package user

import (
	"context"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//

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
	existingUser, _ := u.userRepository.GetUserByEmail(ctx, user.Email)
	if existingUser != nil {
		return ErrUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return u.userRepository.CreateUser(ctx, user)
}

func (u *userUseCase) GetUserByID(ctx context.Context, id uuid.UUID) (*User, error) {
	return u.userRepository.GetUserByID(ctx, id)
}
