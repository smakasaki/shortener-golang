package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/smakasaki/typing-trainer/internal/user"
)

type userHandler struct {
	userUseCase user.UseCase
}

func NewUserHandler(userUseCase user.UseCase) *userHandler {
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
