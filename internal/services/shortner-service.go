package services

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jagdeep/url-shortener-service/internal/models"
	"github.com/jagdeep/url-shortener-service/internal/repositories"
	"github.com/jagdeep/url-shortener-service/internal/utils"
)

type URLShortnerService struct {
	urlRepository *repositories.URLRepository
}

func NewURLShortnerService(urlRepository *repositories.URLRepository) *URLShortnerService {
	return &URLShortnerService{urlRepository: urlRepository}
}

func (s *URLShortnerService) ShortenURL(data *models.URL) (*models.URL, error) {
	const maxAttempts = 5
	var shortCode string

	for range maxAttempts {
		code, err := utils.GenerateRandomCode(8)
		if err != nil {
			return nil, err
		}

		existing, _ := s.urlRepository.Get(code)
		if existing == nil {
			shortCode = code
			break
		}
	}

	if shortCode == "" {
		return nil, errors.New("failed to generate a unique short code after several attempts")
	}

	data.ShortURL = shortCode
	data.CreatedAt = time.Now().UTC()

	result, err := s.urlRepository.Save(data)
	if err != nil {
		return nil, err
	}

	result.ShortURL = BuildFullShortURL(result.ShortURL)

	return result, nil
}

func BuildFullShortURL(code string) string {
	baseURL := os.Getenv("APP_BASE_URL")
	return fmt.Sprintf("%s/%s", baseURL, code)
}

func (s *URLShortnerService) GetOriginalURL(shortCode string) (*models.URL, error) {
	return s.urlRepository.Get(shortCode)
}
