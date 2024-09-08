package url_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/smakasaki/shortener/internal/url"
	"github.com/stretchr/testify/assert"
)

func TestURLRepository_GetByShortCode(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := url.NewRepository(db)
	userID := uuid.New()

	shortCode := "abc123"
	expectedURL := &url.URL{
		ID:          1,
		UserID:      &userID,
		OriginalURL: "http://example.com",
		ShortCode:   shortCode,
		ClickCount:  5,
		CreatedAt:   time.Now(),
	}

	testCases := []struct {
		name          string
		shortCode     string
		mockSetup     func()
		expectedURL   *url.URL
		expectedError error
	}{
		{
			name:      "success",
			shortCode: shortCode,
			mockSetup: func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "original_url", "short_code", "click_count", "created_at"}).
					AddRow(expectedURL.ID, expectedURL.UserID, expectedURL.OriginalURL, expectedURL.ShortCode, expectedURL.ClickCount, expectedURL.CreatedAt)
				mock.ExpectQuery("SELECT id, user_id, original_url, short_code, click_count, created_at FROM urls").
					WithArgs(shortCode).
					WillReturnRows(rows)
			},
			expectedURL:   expectedURL,
			expectedError: nil,
		},
		{
			name:      "not found",
			shortCode: shortCode,
			mockSetup: func() {
				mock.ExpectQuery("SELECT id, user_id, original_url, short_code, click_count, created_at FROM urls").
					WithArgs(shortCode).
					WillReturnError(sql.ErrNoRows)
			},
			expectedURL:   nil,
			expectedError: sql.ErrNoRows,
		},
		{
			name:      "db error",
			shortCode: shortCode,
			mockSetup: func() {
				mock.ExpectQuery("SELECT id, user_id, original_url, short_code, click_count, created_at FROM urls").
					WithArgs(shortCode).
					WillReturnError(errors.New("db error"))
			},
			expectedURL:   nil,
			expectedError: errors.New("db error"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			url, err := repo.GetByShortCode(context.Background(), tt.shortCode)

			assert.Equal(t, tt.expectedURL, url)
			assert.Equal(t, tt.expectedError, err)

			mock.ExpectationsWereMet()
		})
	}
}

func TestURLRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := url.NewRepository(db)

	userID := uuid.New()
	newURL := &url.URL{
		UserID:      &userID,
		OriginalURL: "http://example.com",
		CreatedAt:   time.Now(),
	}

	testCases := []struct {
		name          string
		mockSetup     func()
		expectedError error
	}{
		{
			name: "success",
			mockSetup: func() {
				mock.ExpectBegin()
				mock.ExpectQuery("INSERT INTO urls").WithArgs(userID, newURL.OriginalURL, newURL.CreatedAt, sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id", "created_at"}).AddRow(1, newURL.CreatedAt))
				mock.ExpectExec("UPDATE urls").WithArgs(sqlmock.AnyArg(), 1).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectedError: nil,
		},
		{
			name: "db error",
			mockSetup: func() {
				mock.ExpectBegin()
				mock.ExpectQuery("INSERT INTO urls").WillReturnError(errors.New("db error"))
				mock.ExpectRollback()
			},
			expectedError: errors.New("db error"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			_, err := repo.Create(context.Background(), newURL)

			assert.Equal(t, tt.expectedError, err)
			mock.ExpectationsWereMet()
		})
	}
}

func TestURLRepository_GetAllByUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := url.NewRepository(db)

	userID := uuid.New()
	fixedTime := time.Now()

	testCases := []struct {
		name          string
		limit, offset int
		mockSetup     func()
		expectedURLs  []*url.URL
		expectedError error
	}{
		{
			name:   "success",
			limit:  10,
			offset: 0,
			mockSetup: func() {
				rows := sqlmock.NewRows([]string{"id", "original_url", "short_code", "click_count", "created_at"}).
					AddRow(1, "http://example1.com", "abc123", 5, fixedTime).
					AddRow(2, "http://example2.com", "def456", 3, fixedTime)
				mock.ExpectQuery("SELECT id, original_url, short_code, click_count, created_at FROM urls").
					WithArgs(userID, 10, 0).
					WillReturnRows(rows)
			},
			expectedURLs: []*url.URL{
				{
					ID:          1,
					OriginalURL: "http://example1.com",
					ShortCode:   "abc123",
					ClickCount:  5,
					CreatedAt:   fixedTime,
				},
				{
					ID:          2,
					OriginalURL: "http://example2.com",
					ShortCode:   "def456",
					ClickCount:  3,
					CreatedAt:   fixedTime,
				},
			},
			expectedError: nil,
		},
		{
			name:   "db error",
			limit:  10,
			offset: 0,
			mockSetup: func() {
				mock.ExpectQuery("SELECT id, original_url, short_code, click_count, created_at FROM urls").
					WithArgs(userID, 10, 0).
					WillReturnError(errors.New("db error"))
			},
			expectedURLs:  nil,
			expectedError: errors.New("db error"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			urls, err := repo.GetAllByUser(context.Background(), &userID, tt.limit, tt.offset)

			assert.Equal(t, tt.expectedURLs, urls)
			assert.Equal(t, tt.expectedError, err)

			mock.ExpectationsWereMet()
		})
	}
}
