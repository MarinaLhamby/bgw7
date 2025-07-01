package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestStringsToTicket(t *testing.T) {
	tests := []struct {
		name        string
		input       []string
		expected    Ticket
		expectError bool
	}{
		{
			name: "Valid input",
			input: []string{
				"123", "John Doe", "john@example.com", "Brazil", "14:30", "199.99",
			},
			expected: Ticket{
				ID:            123,
				Name:          "John Doe",
				Email:         "john@example.com",
				Destination:   "Brazil",
				DepartureTime: time.Date(0, 1, 1, 14, 30, 0, 0, time.UTC),
				Price:         199.99,
			},
			expectError: false,
		},
		{
			name:        "Invalid ID",
			input:       []string{"abc", "John Doe", "john@example.com", "Brazil", "14:30", "199.99"},
			expectError: true,
		},
		{
			name:        "Invalid DepartureTime",
			input:       []string{"123", "John Doe", "john@example.com", "Brazil", "invalid", "199.99"},
			expectError: true,
		},
		{
			name:        "Invalid Price",
			input:       []string{"123", "John Doe", "john@example.com", "Brazil", "14:30", "invalid"},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ticket, err := StringsToTicket(tt.input)
			if tt.expectError {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected.ID, ticket.ID)
				require.Equal(t, tt.expected.Name, ticket.Name)
				require.Equal(t, tt.expected.Email, ticket.Email)
				require.Equal(t, tt.expected.Destination, ticket.Destination)
				require.Equal(t, tt.expected.DepartureTime.Hour(), ticket.DepartureTime.Hour())
				require.Equal(t, tt.expected.DepartureTime.Minute(), ticket.DepartureTime.Minute())
				require.Equal(t, tt.expected.Price, ticket.Price)
			}
		})
	}
}

func TestStringToTicketPeriod(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    TicketPeriod
		expectError bool
	}{
		{
			name:     "Valid - início da manhã",
			input:    "início da manhã",
			expected: EarlyMorning,
		},
		{
			name:     "Valid - manhã",
			input:    "manhã",
			expected: Morning,
		},
		{
			name:     "Valid - tarde",
			input:    "tarde",
			expected: Afternoon,
		},
		{
			name:     "Valid - noite",
			input:    "noite",
			expected: Evening,
		},
		{
			name:        "Invalid period",
			input:       "madrugada",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			period, err := StringToTicketPeriod(tt.input)
			if tt.expectError {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, period)
			}
		})
	}
}
