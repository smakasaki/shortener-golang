package url

import (
	"time"

	"github.com/google/uuid"
)

type URL struct {
	ID          int        `json:"id"`
	UserID      *uuid.UUID `json:"userID,omitempty"`
	OriginalURL string     `json:"originalURL"`
	ShortCode   string     `json:"shortCode"`
	ClickCount  int        `json:"clickCount"`
	CreatedAt   time.Time  `json:"createdAt"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty"`
}

type URLStats struct {
	ClickCount  int `json:"clickCount"`
	TotalClicks int `json:"totalClicks"`
}
