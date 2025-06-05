package internal_error

import (
	"fmt"
	"net/http"
)

// HandledError is a custom error for all error in the application domain.
type HandledError struct {
	Message      string
	Code         int
	InternalCode string
}

func New(message string, code int, internalCode string) HandledError {
	return HandledError{
		Message:      message,
		Code:         code,
		InternalCode: internalCode,
	}
}

func (e HandledError) Error() string {
	return fmt.Sprintf("Error: %s\nInternal Code: %s\nStatus Code: %d", e.Message, e.InternalCode, e.Code)
}

var (
	// ErrReadingFile is returned when there is an error while reading the ticket file.
	ErrReadingFile = New("error reading file", http.StatusInternalServerError, "FILE_ERROR_1")
	// ErrParsingTicket is returned when there is an error while parsing a ticket.
	ErrParsingTicket = New("error parsing ticket", http.StatusUnprocessableEntity, "TICKET_ERROR_1")
	// ErrInvalidPeriod is returned when an invalid period is provided.
	ErrInvalidPeriod = New("invalid period", http.StatusBadRequest, "TICKET_ERROR_2")
)
