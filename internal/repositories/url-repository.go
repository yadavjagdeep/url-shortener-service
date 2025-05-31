package repositories

import (
	"database/sql"
	"fmt"

	"github.com/jagdeep/url-shortener-service/internal/models"
	"github.com/jagdeep/url-shortener-service/internal/utils"
)

type URLRepository struct {
	db *sql.DB
}

func NewURLRepository() *URLRepository {
	return &URLRepository{db: GetMySqlClient()}
}

func (r *URLRepository) Save(data *models.URL) (*models.URL, error) {
	result, err := r.db.Exec("INSERT INTO urls (long_url, short_url) VALUES (?, ?)", data.LongURL, data.ShortURL)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	data.ID = int(id)
	return data, nil
}

func (r *URLRepository) Get(shortURL string) (*models.URL, error) {
	row := r.db.QueryRow("SELECT id, long_url, short_url, created_at FROM urls WHERE short_url = ?", shortURL)

	var url models.URL

	err := row.Scan(
		&url.ID,
		&url.LongURL,
		&url.ShortURL,
		&url.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NotFoundError(fmt.Sprintf("long url not found with short url %s not found", shortURL))
		}

		return nil, fmt.Errorf("error in fetching long url: %v", err)
	}

	return &url, nil
}
