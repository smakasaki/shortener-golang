package common

import "github.com/labstack/echo/v4"

type AuthMiddleware interface {
	CheckSession(next echo.HandlerFunc) echo.HandlerFunc
}
