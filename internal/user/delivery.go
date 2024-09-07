package user

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func RegisterEndpoints(e *echo.Echo, uc UseCase) {
	h := NewUserHandler(uc)
	e.POST("/users", h.Create)
	e.GET("/users/profile", h.GetProfile)

}

type UseCase interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByID(ctx context.Context, id uuid.UUID) (*User, error)
}

type userHandler struct {
	userUseCase UseCase
}

func NewUserHandler(userUseCase UseCase) *userHandler {
	return &userHandler{
		userUseCase: userUseCase,
	}
}

func (h *userHandler) Create(c echo.Context) error {
	return nil
}

func (h *userHandler) GetProfile(c echo.Context) error {
	return c.String(http.StatusOK, "User profile")
}
