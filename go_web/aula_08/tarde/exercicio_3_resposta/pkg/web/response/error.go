package response

import (
	"fmt"
	"net/http"
	"strconv"
)

type ApiError struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
	Status     string `json:"status"`
}

var (
	ErrEntityNotFound           = New("resource %s of id %s not found", http.StatusNotFound)
	ErrorDecoding               = New("error decoding %s", http.StatusBadRequest)
	ErrorEncoding               = New("error encoding %s", http.StatusBadRequest)
	ErrFile                     = New("error while manipulating file", http.StatusBadRequest)
	ErrProductCodeAlreadyExists = New("error code already exists", http.StatusConflict)
	ErrValidation               = New("%s", http.StatusBadRequest)
)

func New(message string, code int) ApiError {
	return ApiError{
		StatusCode: code,
		Message:    message,
		Status:     strconv.Itoa(code),
	}
}

func (e ApiError) Format(args ...interface{}) ApiError {
	return ApiError{
		StatusCode: e.StatusCode,
		Message:    fmt.Sprintf(e.Message, args...),
		Status:     e.Status,
	}
}

func (e ApiError) Error() string {
	return e.Message
}
