package models

import "time"

type URL struct {
	ID        int       `json:"id" db:"id"`
	LongURL   string    `json:"long_url" db:"long_url"`
	ShortURL  string    `json:"short_url" db:"short_url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
