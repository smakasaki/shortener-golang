package session

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/smakasaki/shortener/internal/common"
	"github.com/smakasaki/shortener/pkg/validation"
)

var (
	SessionCookieName     = "sessionID"
	SessionEchoStorageKey = "session"
)

func RegisterEndpoints(e *echo.Echo, sessionUseCase UseCase, authMiddleware common.AuthMiddleware) {
	h := NewHandler(sessionUseCase)
	e.POST("/sessions", h.Create)
	e.DELETE("/sessions", h.Delete, authMiddleware.CheckSession)
}

type UseCase interface {
	Create(ctx context.Context, email string, password string) (uuid.UUID, error)
	Delete(ctx context.Context, sessionID uuid.UUID, userID uuid.UUID) error
}

type sessionHandler struct {
	sessionUseCase UseCase
}

func NewHandler(sessionUseCase UseCase) *sessionHandler {
	return &sessionHandler{
		sessionUseCase: sessionUseCase,
	}
}

func (h *sessionHandler) Create(c echo.Context) error {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&credentials); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	rules := []validation.Rule{
		validation.ValidateEmail("Email", credentials.Email),
		validation.ValidateLength("Password", credentials.Password, 6, 50),
	}

	errors := validation.Execute(rules)

	if len(errors) > 0 {
		var errorMessages []string
		for _, err := range errors {
			errorMessages = append(errorMessages, err.Error())
		}
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors": errorMessages,
		})
	}

	session, err := h.sessionUseCase.Create(c.Request().Context(), credentials.Email, credentials.Password)
	if err != nil {
		return c.JSON(400, map[string]string{"error": ErrInvalidCredentials.Error()})
	}

	cookie := &http.Cookie{
		Name:     SessionCookieName,
		Value:    session.String(),
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusCreated, map[string]string{"message": "Session created"})
}

func (h *sessionHandler) Delete(c echo.Context) error {
	session := c.Get(SessionEchoStorageKey).(*Session)

	err := h.sessionUseCase.Delete(c.Request().Context(), session.ID, session.UserID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	cookie := &http.Cookie{
		Name:     SessionCookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{"message": "Session deleted"})
}
