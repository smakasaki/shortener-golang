package user

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/smakasaki/shortener/internal/common"
	"github.com/smakasaki/shortener/pkg/validation"
)

func RegisterEndpoints(e *echo.Echo, uc UseCase, authMiddleware common.AuthMiddleware) {
	h := NewUserHandler(uc)
	e.POST("/users", h.Create)
	e.GET("/users/profile", h.GetProfile, authMiddleware.CheckSession)

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
	var newUser User
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	rules := []validation.Rule{
		validation.ValidateRequired("Email", newUser.Email),
		validation.ValidateEmail("Email", newUser.Email),
		validation.ValidateRequired("Password", newUser.Password),
		validation.ValidateLength("Password", newUser.Password, 6, 50),
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

	if err := h.userUseCase.CreateUser(c.Request().Context(), &newUser); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create user"})
	}

	return c.String(http.StatusCreated, "User created")
}

func (h *userHandler) GetProfile(c echo.Context) error {
	userID := c.Get(common.SessionEchoStorageKey).(*common.Session).UserID
	user, err := h.userUseCase.GetUserByID(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not fetch user"})
	}

	return c.JSON(http.StatusOK, user)
}
