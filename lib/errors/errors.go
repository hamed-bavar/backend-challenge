package errors

import "net/http"

type AppError struct {
	Code    int
	Message string `json:"message"`
}

func NotFoundError(message string) *AppError {
	return &AppError{http.StatusNotFound, message}
}
func InternalServerError(message string) *AppError {
	return &AppError{http.StatusInternalServerError, message}
}
func ValidationError(message string) *AppError {
	return &AppError{http.StatusBadRequest, message}
}
