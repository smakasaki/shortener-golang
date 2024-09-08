package common

import (
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware interface {
	CheckSession(next echo.HandlerFunc) echo.HandlerFunc
	OptionalSession(next echo.HandlerFunc) echo.HandlerFunc
}

type Session struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"userID"`
	CreatedAt time.Time `json:"createdAt"`
}

var (
	SessionCookieName     = "sessionID"
	SessionEchoStorageKey = "session"
)
