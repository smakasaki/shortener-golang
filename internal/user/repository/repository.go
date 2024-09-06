package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/smakasaki/typing-trainer/domain"
	"github.com/smakasaki/typing-trainer/internal/user"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) user.Repository {
	return &userRepository{
		db: db,
	}
}

//TODO: Add migrations and prepare the database

func (r *userRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	// TODO: Implementation here
	return nil, nil
}

func (r *userRepository) CreateUser(ctx context.Context, u *domain.User) error {
	// TODO: Implementation here
	return nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return nil, nil
}
