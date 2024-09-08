package url_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/smakasaki/shortener/internal/common"
	"github.com/smakasaki/shortener/internal/url"
	"github.com/smakasaki/shortener/internal/url/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestURLHandler_Redirect(t *testing.T) {
	mockUseCase := new(mocks.MockUseCase)
	handler := url.NewHandler(mockUseCase)

	shortCode := "abc123"
	expectedURL := &url.URL{
		ID:          1,
		OriginalURL: "http://example.com",
	}

	t.Run("success", func(t *testing.T) {
		mockUseCase.On("GetByShortCode", mock.Anything, shortCode).Return(expectedURL, nil).Once()
		mockUseCase.On("IncrementClickCount", mock.Anything, expectedURL.ID, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/s/"+shortCode, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("shortCode")
		c.SetParamValues(shortCode)

		err := handler.Redirect(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusFound, rec.Code)
		assert.Equal(t, expectedURL.OriginalURL, rec.Header().Get("Location"))
		mockUseCase.AssertExpectations(t)
	})

	t.Run("URL not found", func(t *testing.T) {
		mockUseCase.On("GetByShortCode", mock.Anything, shortCode).Return(nil, errors.New("not found")).Once()

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/s/"+shortCode, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("shortCode")
		c.SetParamValues(shortCode)

		err := handler.Redirect(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestURLHandler_Create(t *testing.T) {
	mockUseCase := new(mocks.MockUseCase)
	handler := url.NewHandler(mockUseCase)

	input := map[string]string{
		"originalUrl": "http://example.com",
	}
	expectedURL := &url.URL{
		ID:          1,
		OriginalURL: input["originalUrl"],
		ShortCode:   "abc123",
	}

	t.Run("success", func(t *testing.T) {
		mockUseCase.On("Create", mock.Anything, mock.AnythingOfType("*uuid.UUID"), input["originalUrl"]).Return(expectedURL, nil).Once()

		e := echo.New()
		reqBody, _ := json.Marshal(input)
		req := httptest.NewRequest(http.MethodPost, "/urls", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handler.Create(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("invalid URL", func(t *testing.T) {
		input := map[string]string{
			"originalUrl": "invalid-url",
		}

		e := echo.New()
		reqBody, _ := json.Marshal(input)
		req := httptest.NewRequest(http.MethodPost, "/urls", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handler.Create(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("creation error", func(t *testing.T) {
		mockUseCase.On("Create", mock.Anything, mock.AnythingOfType("*uuid.UUID"), input["originalUrl"]).Return(nil, errors.New("db error")).Once()

		e := echo.New()
		reqBody, _ := json.Marshal(input)
		req := httptest.NewRequest(http.MethodPost, "/urls", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handler.Create(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestURLHandler_GetAll(t *testing.T) {
	mockUseCase := new(mocks.MockUseCase)
	handler := url.NewHandler(mockUseCase)

	userID := uuid.New()
	limit := 10
	offset := 0
	expectedURLs := []*url.URL{
		{
			ID:          1,
			OriginalURL: "http://example1.com",
			ShortCode:   "abc123",
			ClickCount:  5,
		},
		{
			ID:          2,
			OriginalURL: "http://example2.com",
			ShortCode:   "def456",
			ClickCount:  3,
		},
	}

	t.Run("success", func(t *testing.T) {
		mockUseCase.On("GetAll", mock.Anything, &userID, limit, offset).Return(expectedURLs, nil).Once()

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/url?limit="+strconv.Itoa(limit)+"&offset="+strconv.Itoa(offset), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.Set("session", &common.Session{UserID: userID})

		err := handler.GetAll(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUseCase.On("GetAll", mock.Anything, &userID, limit, offset).Return(nil, errors.New("db error")).Once()

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/url?limit="+strconv.Itoa(limit)+"&offset="+strconv.Itoa(offset), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.Set("session", &common.Session{UserID: userID})

		err := handler.GetAll(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestURLHandler_Delete(t *testing.T) {
	mockUseCase := new(mocks.MockUseCase)
	handler := url.NewHandler(mockUseCase)

	shortCode := "abc123"
	userID := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockUseCase.On("Delete", mock.Anything, shortCode, &userID).Return(nil).Once()

		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/urls/"+shortCode, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("session", &common.Session{UserID: userID})
		c.SetParamNames("shortCode")
		c.SetParamValues(shortCode)

		err := handler.Delete(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUseCase.On("Delete", mock.Anything, shortCode, &userID).Return(errors.New("db error")).Once()

		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/urls/"+shortCode, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("session", &common.Session{UserID: userID})
		c.SetParamNames("shortCode")
		c.SetParamValues(shortCode)

		err := handler.Delete(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestURLHandler_GetStats(t *testing.T) {
	mockUseCase := new(mocks.MockUseCase)
	handler := url.NewHandler(mockUseCase)

	shortCode := "abc123"
	userID := uuid.New()
	expectedStats := &url.URLStats{
		ClickCount:  10,
		TotalClicks: 20,
	}

	t.Run("success", func(t *testing.T) {
		mockUseCase.On("GetStats", mock.Anything, shortCode, &userID).Return(expectedStats, nil).Once()

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/urls/"+shortCode+"/stats", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("session", &common.Session{UserID: userID})
		c.SetParamNames("shortCode")
		c.SetParamValues(shortCode)

		err := handler.GetStats(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUseCase.On("GetStats", mock.Anything, shortCode, &userID).Return(nil, errors.New("db error")).Once()

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/urls/"+shortCode+"/stats", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("session", &common.Session{UserID: userID})
		c.SetParamNames("shortCode")
		c.SetParamValues(shortCode)

		err := handler.GetStats(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}
