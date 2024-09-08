package session

import (
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/smakasaki/shortener/internal/common"
	"github.com/smakasaki/shortener/internal/user"
)

type authMiddleware struct {
	sessionRepo Repository
	userRepo    user.Repository
}

func NewAuthMiddleware(sessionRepo Repository, userRepo user.Repository) *authMiddleware {
	return &authMiddleware{
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
	}
}

func (m *authMiddleware) CheckSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(common.SessionCookieName)
		if err != nil {
			return c.JSON(401, map[string]string{"error": "Unauthorized"})
		}

		sessionID, err := uuid.Parse(cookie.Value)
		if err != nil {
			return c.JSON(401, map[string]string{"error": "Unauthorized"})
		}
		session, err := m.sessionRepo.Get(c.Request().Context(), sessionID)
		if err != nil {
			return c.JSON(401, map[string]string{"error": "Unauthorized"})
		}

		if time.Since(session.CreatedAt) > 24*time.Hour {
			return c.JSON(401, map[string]string{"error": ErrSessionExpired.Error()})
		}

		commonSession := &common.Session{
			ID:        session.ID,
			UserID:    session.UserID,
			CreatedAt: session.CreatedAt,
		}

		c.Set(common.SessionEchoStorageKey, commonSession)
		return next(c)
	}
}

func (m *authMiddleware) OptionalSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(common.SessionCookieName)
		if err != nil {
			return next(c)
		}

		sessionID, err := uuid.Parse(cookie.Value)
		if err != nil {
			return next(c)
		}
		session, err := m.sessionRepo.Get(c.Request().Context(), sessionID)
		if err != nil {
			return next(c)
		}

		if time.Since(session.CreatedAt) > 24*time.Hour {
			return next(c)
		}

		commonSession := &common.Session{
			ID:        session.ID,
			UserID:    session.UserID,
			CreatedAt: session.CreatedAt,
		}

		c.Set(common.SessionEchoStorageKey, commonSession)
		return next(c)
	}
}
