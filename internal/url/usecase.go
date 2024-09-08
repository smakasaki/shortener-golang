package url

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Repository interface {
	GetByShortCode(ctx context.Context, shortCode string) (*URL, error)
	Create(ctx context.Context, url *URL) (*URL, error)
	GetAllByUser(ctx context.Context, userID *uuid.UUID, limit, offset int) ([]*URL, error)
	Delete(ctx context.Context, shortCode string, userID *uuid.UUID) error
	GetStats(ctx context.Context, shortCode string, userID *uuid.UUID) (*URLStats, error)
	IncrementClick(ctx context.Context, urlID int) error
	CreateClick(ctx context.Context, urlID int, ipAddress, userAgent, referer string) error
}

type urlUseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *urlUseCase {
	return &urlUseCase{
		repo: repo,
	}
}

func (uc *urlUseCase) GetByShortCode(ctx context.Context, shortCode string) (*URL, error) {
	return uc.repo.GetByShortCode(ctx, shortCode)
}

func (uc *urlUseCase) Create(ctx context.Context, userID *uuid.UUID, originalURL string) (*URL, error) {
	url := &URL{
		UserID:      userID,
		OriginalURL: originalURL,
		ShortCode:   "computeMe",
		CreatedAt:   time.Now(),
	}

	newUrl, err := uc.repo.Create(ctx, url)
	if err != nil {
		return nil, err
	}
	return newUrl, nil
}

func (uc *urlUseCase) GetAll(ctx context.Context, userID *uuid.UUID, limit, offset int) ([]*URL, error) {
	return uc.repo.GetAllByUser(ctx, userID, limit, offset)
}

func (uc *urlUseCase) Delete(ctx context.Context, shortCode string, userID *uuid.UUID) error {
	return uc.repo.Delete(ctx, shortCode, userID)
}

func (uc *urlUseCase) GetStats(ctx context.Context, shortCode string, userID *uuid.UUID) (*URLStats, error) {
	return uc.repo.GetStats(ctx, shortCode, userID)
}

func (uc *urlUseCase) IncrementClickCount(ctx context.Context, urlID int, ipAddress, userAgent, referer string) error {
	err := uc.repo.CreateClick(ctx, urlID, ipAddress, userAgent, referer)
	if err != nil {
		return err
	}
	return uc.repo.IncrementClick(ctx, urlID)
}
