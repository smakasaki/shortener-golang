package url

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/smakasaki/shortener/internal/common"
)

func RegisterEndpoints(e *echo.Echo, urlUseCase UseCase, authMiddleware common.AuthMiddleware) {
	h := NewHandler(urlUseCase)
	e.GET("/s/:shortCode", h.Redirect)
	e.GET("/urls", h.GetAll, authMiddleware.CheckSession)
	e.POST("/urls", h.Create, authMiddleware.OptionalSession)
	e.DELETE("/urls/:shortCode", h.Delete, authMiddleware.CheckSession)
	e.GET("/urls/:shortCode/stats", h.GetStats, authMiddleware.CheckSession)
}

type UseCase interface {
	GetByShortCode(ctx context.Context, shortCode string) (*URL, error)
	Create(ctx context.Context, userID *uuid.UUID, originalURL string) (*URL, error)
	GetAll(ctx context.Context, userID *uuid.UUID, limit, offset int) ([]*URL, error)
	Delete(ctx context.Context, shortCode string, userID *uuid.UUID) error
	GetStats(ctx context.Context, shortCode string, userID *uuid.UUID) (*URLStats, error)
	IncrementClickCount(ctx context.Context, urlID int, ipAddress, userAgent, referer string) error
}

type urlHandler struct {
	urlUseCase UseCase
}

func NewHandler(urlUseCase UseCase) *urlHandler {
	return &urlHandler{
		urlUseCase: urlUseCase,
	}
}

func (h *urlHandler) Redirect(c echo.Context) error {
	shortCode := c.Param("shortCode")
	url, err := h.urlUseCase.GetByShortCode(c.Request().Context(), shortCode)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "URL not found"})
	}

	ipAddress := c.RealIP()
	userAgent := c.Request().UserAgent()
	referer := c.Request().Referer()

	h.urlUseCase.IncrementClickCount(c.Request().Context(), url.ID, ipAddress, userAgent, referer)
	return c.Redirect(http.StatusFound, url.OriginalURL)
}

func (h *urlHandler) GetAll(c echo.Context) error {
	userID := c.Get(common.SessionEchoStorageKey).(*common.Session).UserID
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid limit"})
	}
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid offset"})
	}

	urls, err := h.urlUseCase.GetAll(c.Request().Context(), &userID, limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch URLs"})
	}
	return c.JSON(http.StatusOK, urls)
}

func (h *urlHandler) Create(c echo.Context) error {
	var input struct {
		OriginalURL string `json:"originalUrl"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	parsedURL, err := url.ParseRequestURI(input.OriginalURL)
	if err != nil || !parsedURL.IsAbs() {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid URL"})
	}

	var userID *uuid.UUID
	sessionValue := c.Get(common.SessionEchoStorageKey)
	if sessionValue != nil {
		userID = &sessionValue.(*common.Session).UserID
	} else {
		userID = nil
	}

	shortURL, err := h.urlUseCase.Create(c.Request().Context(), userID, input.OriginalURL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, shortURL)
}

func (h *urlHandler) Delete(c echo.Context) error {
	shortCode := c.Param("shortCode")
	userID := c.Get(common.SessionEchoStorageKey).(*common.Session).UserID

	err := h.urlUseCase.Delete(c.Request().Context(), shortCode, &userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "URL not found or unauthorized"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "URL deleted"})
}

func (h *urlHandler) GetStats(c echo.Context) error {
	shortCode := c.Param("shortCode")
	userID := c.Get(common.SessionEchoStorageKey).(*common.Session).UserID

	stats, err := h.urlUseCase.GetStats(c.Request().Context(), shortCode, &userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "URL not found or unauthorized"})
	}
	return c.JSON(http.StatusOK, stats)
}
