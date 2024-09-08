package session

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type sessionRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *sessionRepository {
	return &sessionRepository{
		db: db,
	}
}

func (repo *sessionRepository) Create(ctx context.Context, userId uuid.UUID) (uuid.UUID, error) {
	query := `
		INSERT INTO sessions (user_id)
		VALUES ($1)
		RETURNING id`
	var id uuid.UUID
	err := repo.db.
		QueryRowContext(ctx, query, userId).
		Scan(&id)
	return id, err
}

func (repo *sessionRepository) Get(ctx context.Context, sessionId uuid.UUID) (*Session, error) {
	query := `
	SELECT * FROM sessions WHERE id = $1`
	session := &Session{}
	err := repo.db.
		QueryRowContext(ctx, query, sessionId).
		Scan(&session.ID, &session.UserID, &session.CreatedAt)
	if err != nil {
		return nil, ErrSessionNotFound
	}
	return session, nil
}

func (repo *sessionRepository) Delete(ctx context.Context, sessionId uuid.UUID) error {
	query := `
	DELETE FROM sessions WHERE id = $1`
	_, err := repo.db.ExecContext(ctx, query, sessionId)
	if err != nil {
		return ErrSessionNotFound
	}
	return nil
}
