package session

import (
	"context"

	"github.com/google/uuid"
	"github.com/smakasaki/shortener/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	Create(ctx context.Context, userId uuid.UUID) (uuid.UUID, error)
	Get(ctx context.Context, sessionId uuid.UUID) (*Session, error)
	Delete(ctx context.Context, sessionId uuid.UUID) error
}

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*user.User, error)
}

type sessionUseCase struct {
	repo     Repository
	userRepo UserRepository
}

func NewUseCase(repo Repository, userRepo UserRepository) *sessionUseCase {
	return &sessionUseCase{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (uc *sessionUseCase) Create(ctx context.Context, email string, password string) (uuid.UUID, error) {
	user, err := uc.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return uuid.Nil, ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return uuid.Nil, ErrInvalidCredentials
	}

	return uc.repo.Create(ctx, user.ID)
}

func (uc *sessionUseCase) Delete(ctx context.Context, sessionID uuid.UUID, userID uuid.UUID) error {
	session, err := uc.repo.Get(ctx, sessionID)
	if err != nil {
		return ErrSessionNotFound
	}

	if session.UserID != userID {
		return ErrSessionNotFound
	}

	return uc.repo.Delete(ctx, sessionID)
}
