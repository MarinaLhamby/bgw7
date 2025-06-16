package internal_error

import (
	"fmt"
)

// HandledError is a custom error for all error in the application domain.
type HandledError struct {
	Message      string
	InternalCode string
}

// New creates a HandledError
func New(message string, internalCode string) HandledError {
	return HandledError{
		Message:      message,
		InternalCode: internalCode,
	}
}

func (e HandledError) Error() string {
	return fmt.Sprintf("Error: %s\nInternal Code: %s", e.Message, e.InternalCode)
}

var (
	// ErrReadingFile is returned when there is an error while reading the ticket file.
	ErrReadingFile = New("error reading file", "FILE_ERROR_1")
	// ErrParsingTicket is returned when there is an error while parsing a ticket.
	ErrParsingTicket = New("error parsing ticket", "TICKET_ERROR_1")
	// ErrInvalidPeriod is returned when an invalid period is provided.
	ErrInvalidPeriod = New("invalid period", "TICKET_ERROR_2")
)
