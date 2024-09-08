package url

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type urlRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *urlRepository {
	return &urlRepository{
		db: db,
	}
}

func (repo *urlRepository) GetByShortCode(ctx context.Context, shortCode string) (*URL, error) {
	query := `SELECT id, user_id, original_url, short_code, click_count, created_at FROM urls
	 		  WHERE short_code = $1 AND deleted_at IS NULL`
	row := repo.db.QueryRowContext(ctx, query, shortCode)

	var url URL
	err := row.Scan(&url.ID, &url.UserID, &url.OriginalURL, &url.ShortCode, &url.ClickCount, &url.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &url, nil
}

func (repo *urlRepository) Create(ctx context.Context, url *URL) (*URL, error) {
	var insertQuery string

	tx, err := repo.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// Step 1: Insert the URL without the short code and get the generated ID
	insertQuery = `INSERT INTO urls (user_id, original_url, created_at, short_code) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	err = tx.QueryRowContext(ctx, insertQuery, url.UserID, url.OriginalURL, url.CreatedAt, url.ShortCode).Scan(&url.ID, &url.CreatedAt)
	if err != nil {
		return nil, err
	}

	// Step 2: Generate the short code using the generated ID
	shortCode := generateShortCode(url.ID)

	// Step 3: Update the URL record with the generated short code
	updateQuery := `UPDATE urls SET short_code = $1 WHERE id = $2`
	_, err = tx.ExecContext(ctx, updateQuery, shortCode, url.ID)
	if err != nil {
		return nil, err
	}

	// Assign the generated short code back to the URL object
	url.ShortCode = shortCode
	return url, nil
}

func (repo *urlRepository) GetAllByUser(ctx context.Context, userID *uuid.UUID, limit, offset int) ([]*URL, error) {
	query := `SELECT id, original_url, short_code, click_count, created_at FROM urls WHERE user_id = $1 AND deleted_at IS NULL LIMIT $2 OFFSET $3`
	rows, err := repo.db.QueryContext(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []*URL
	for rows.Next() {
		var url URL
		err := rows.Scan(&url.ID, &url.OriginalURL, &url.ShortCode, &url.ClickCount, &url.CreatedAt)
		if err != nil {
			return nil, err
		}
		urls = append(urls, &url)
	}
	return urls, nil
}

func (repo *urlRepository) Delete(ctx context.Context, shortCode string, userID *uuid.UUID) error {
	query := `UPDATE urls SET deleted_at = $1 WHERE short_code = $2 AND user_id = $3`
	_, err := repo.db.ExecContext(ctx, query, time.Now(), shortCode, userID)
	return err
}

func (repo *urlRepository) GetStats(ctx context.Context, shortCode string, userID *uuid.UUID) (*URLStats, error) {
	query := `
		SELECT u.click_count, COUNT(uc.id) AS clicks
		FROM urls u
		LEFT JOIN url_clicks uc ON u.id = uc.url_id
		WHERE u.short_code = $1 AND u.user_id = $2 AND u.deleted_at IS NULL
		GROUP BY u.click_count`
	row := repo.db.QueryRowContext(ctx, query, shortCode, userID)

	var stats URLStats
	err := row.Scan(&stats.ClickCount, &stats.TotalClicks)
	if err != nil {
		return nil, err
	}
	return &stats, nil
}

func (repo *urlRepository) IncrementClick(ctx context.Context, urlID int) error {
	query := `UPDATE urls SET click_count = click_count + 1 WHERE id = $1`
	_, err := repo.db.ExecContext(ctx, query, urlID)
	return err
}

func (repo *urlRepository) CreateClick(ctx context.Context, urlID int, ipAddress, userAgent, referer string) error {
	query := `INSERT INTO url_clicks (url_id, ip_address, user_agent, referer) VALUES ($1, $2, $3, $4)`
	_, err := repo.db.ExecContext(ctx, query, urlID, ipAddress, userAgent, referer)
	return err
}

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func generateShortCode(id int) string {
	if id == 0 {
		return string(base62Chars[0])
	}

	code := ""
	for id > 0 {
		remainder := id % 62
		code = string(base62Chars[remainder]) + code
		id = id / 62
	}

	return code
}
