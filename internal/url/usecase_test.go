package url_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/smakasaki/shortener/internal/url"
	"github.com/smakasaki/shortener/internal/url/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestURLUseCase_GetByShortCode(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	useCase := url.NewUseCase(mockRepo)

	shortCode := "abc123"
	userId := uuid.New()
	expectedURL := &url.URL{
		ID:          1,
		UserID:      &userId,
		OriginalURL: "http://example.com",
		ShortCode:   shortCode,
		ClickCount:  5,
		CreatedAt:   time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetByShortCode", mock.Anything, shortCode).Return(expectedURL, nil).Once()

		url, err := useCase.GetByShortCode(context.Background(), shortCode)

		assert.NoError(t, err)
		assert.Equal(t, expectedURL, url)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("GetByShortCode", mock.Anything, shortCode).Return(nil, errors.New("db error")).Once()

		url, err := useCase.GetByShortCode(context.Background(), shortCode)

		assert.Nil(t, url)
		assert.EqualError(t, err, "db error")
		mockRepo.AssertExpectations(t)
	})
}

func TestURLUseCase_Create(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	useCase := url.NewUseCase(mockRepo)

	userID := uuid.New()
	originalURL := "http://example.com"
	expectedURL := &url.URL{
		ID:          1,
		UserID:      &userID,
		OriginalURL: originalURL,
		ShortCode:   "abc123",
		CreatedAt:   time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*url.URL")).Return(expectedURL, nil).Once()

		url, err := useCase.Create(context.Background(), &userID, originalURL)

		assert.NoError(t, err)
		assert.Equal(t, expectedURL, url)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*url.URL")).Return(nil, errors.New("db error")).Once()

		url, err := useCase.Create(context.Background(), &userID, originalURL)

		assert.Nil(t, url)
		assert.EqualError(t, err, "db error")
		mockRepo.AssertExpectations(t)
	})
}

func TestURLUseCase_GetAll(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	useCase := url.NewUseCase(mockRepo)

	userID := uuid.New()
	limit, offset := 10, 0
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
		mockRepo.On("GetAllByUser", mock.Anything, &userID, limit, offset).Return(expectedURLs, nil).Once()

		urls, err := useCase.GetAll(context.Background(), &userID, limit, offset)

		assert.NoError(t, err)
		assert.Equal(t, expectedURLs, urls)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("GetAllByUser", mock.Anything, &userID, limit, offset).Return(nil, errors.New("db error")).Once()

		urls, err := useCase.GetAll(context.Background(), &userID, limit, offset)

		assert.Nil(t, urls)
		assert.EqualError(t, err, "db error")
		mockRepo.AssertExpectations(t)
	})
}

func TestURLUseCase_Delete(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	useCase := url.NewUseCase(mockRepo)

	shortCode := "abc123"
	userID := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockRepo.On("Delete", mock.Anything, shortCode, &userID).Return(nil).Once()

		err := useCase.Delete(context.Background(), shortCode, &userID)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("Delete", mock.Anything, shortCode, &userID).Return(errors.New("db error")).Once()

		err := useCase.Delete(context.Background(), shortCode, &userID)

		assert.EqualError(t, err, "db error")
		mockRepo.AssertExpectations(t)
	})
}

func TestURLUseCase_GetStats(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	useCase := url.NewUseCase(mockRepo)

	shortCode := "abc123"
	userID := uuid.New()
	expectedStats := &url.URLStats{
		ClickCount:  10,
		TotalClicks: 20,
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetStats", mock.Anything, shortCode, &userID).Return(expectedStats, nil).Once()

		stats, err := useCase.GetStats(context.Background(), shortCode, &userID)

		assert.NoError(t, err)
		assert.Equal(t, expectedStats, stats)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("GetStats", mock.Anything, shortCode, &userID).Return(nil, errors.New("db error")).Once()

		stats, err := useCase.GetStats(context.Background(), shortCode, &userID)

		assert.Nil(t, stats)
		assert.EqualError(t, err, "db error")
		mockRepo.AssertExpectations(t)
	})
}

func TestURLUseCase_IncrementClickCount(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	useCase := url.NewUseCase(mockRepo)

	urlID := 1
	ipAddress := "127.0.0.1"
	userAgent := "Mozilla"
	referer := "http://example.com"

	t.Run("success", func(t *testing.T) {
		mockRepo.On("CreateClick", mock.Anything, urlID, ipAddress, userAgent, referer).Return(nil).Once()
		mockRepo.On("IncrementClick", mock.Anything, urlID).Return(nil).Once()

		err := useCase.IncrementClickCount(context.Background(), urlID, ipAddress, userAgent, referer)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("create click error", func(t *testing.T) {
		mockRepo.On("CreateClick", mock.Anything, urlID, ipAddress, userAgent, referer).Return(errors.New("db error")).Once()

		err := useCase.IncrementClickCount(context.Background(), urlID, ipAddress, userAgent, referer)

		assert.EqualError(t, err, "db error")
		mockRepo.AssertExpectations(t)
	})

	t.Run("increment click error", func(t *testing.T) {
		mockRepo.On("CreateClick", mock.Anything, urlID, ipAddress, userAgent, referer).Return(nil).Once()
		mockRepo.On("IncrementClick", mock.Anything, urlID).Return(errors.New("db error")).Once()

		err := useCase.IncrementClickCount(context.Background(), urlID, ipAddress, userAgent, referer)

		assert.EqualError(t, err, "db error")
		mockRepo.AssertExpectations(t)
	})
}
