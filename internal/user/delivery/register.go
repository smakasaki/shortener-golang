package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/smakasaki/typing-trainer/internal/user"
)

func RegisterEndpoints(e *echo.Echo, uc user.UseCase) {
	h := NewUserHandler(uc)
	e.POST("/users", h.Create)
	e.GET("/users/profile", h.GetProfile)

}
