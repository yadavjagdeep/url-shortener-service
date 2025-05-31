package utils

import (
	"fmt"
	"net/http"
)

type AppError struct {
	StatusCode int
	Message    string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s (status code: %d)", e.Message, e.StatusCode)
}

func NewError(message string, status_code int) *AppError {
	return &AppError{Message: message, StatusCode: status_code}
}

func NotFoundError(message string) *AppError {
	return NewError(message, http.StatusNotFound)
}

func InternalServerError(message string) *AppError {
	return NewError(message, http.StatusInternalServerError)
}

func UnothorizedError(message string) *AppError {
	return NewError(message, http.StatusUnauthorized)
}

func ForbiddenError(message string) *AppError {
	return NewError(message, http.StatusForbidden)
}
