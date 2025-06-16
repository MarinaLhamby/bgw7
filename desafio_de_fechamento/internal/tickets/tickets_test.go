package tickets

import (
	"testing"
	"time"

	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/internal/internal_error"
	"github.com/stretchr/testify/require"
)

func TestGetTotalTicketsByDestination(t *testing.T) {
	dataTest := []struct {
		name        string
		destination string
		tickets     []Ticket
		expected    int
	}{
		{
			name:        "Succesfully get total tickets by destination",
			destination: "Brazil",
			tickets: []Ticket{
				{Destination: "Brazil"},
				{Destination: "Brazil"},
				{Destination: "USA"},
				{Destination: "Brazil"},
				{Destination: "Canada"},
			},
			expected: 3,
		},
		{
			name:        "No tickets for destination",
			destination: "Brazil",
			tickets: []Ticket{
				{Destination: "USA"},
				{Destination: "Canada"},
			},
			expected: 0,
		},
		{
			name:        "Empty tickets slice",
			destination: "Brazil",
			tickets:     []Ticket{},
			expected:    0,
		},
	}
	for _, tt := range dataTest {
		t.Run(tt.name, func(t *testing.T) {
			result := GetTotalTicketsByDestination(tt.destination, tt.tickets)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestGetCountByPeriod(t *testing.T) {
	tickets := []Ticket{
		{DepartureTime: time.Date(2023, 1, 1, 4, 0, 0, 0, time.UTC)},
		{DepartureTime: time.Date(2023, 1, 1, 5, 0, 0, 0, time.UTC)},
		{DepartureTime: time.Date(2023, 1, 1, 8, 0, 0, 0, time.UTC)},
		{DepartureTime: time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC)},
		{DepartureTime: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC)},
		{DepartureTime: time.Date(2023, 1, 1, 11, 0, 0, 0, time.UTC)},
		{DepartureTime: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)},
		{DepartureTime: time.Date(2023, 1, 1, 13, 0, 0, 0, time.UTC)},
		{DepartureTime: time.Date(2023, 1, 1, 14, 0, 0, 0, time.UTC)},
		{DepartureTime: time.Date(2023, 1, 1, 15, 0, 0, 0, time.UTC)},
		{DepartureTime: time.Date(2023, 1, 1, 16, 0, 0, 0, time.UTC)},
		{DepartureTime: time.Date(2023, 1, 1, 17, 0, 0, 0, time.UTC)},
		{DepartureTime: time.Date(2023, 1, 1, 18, 0, 0, 0, time.UTC)},
		{DepartureTime: time.Date(2023, 1, 1, 19, 0, 0, 0, time.UTC)},
		{DepartureTime: time.Date(2023, 1, 1, 20, 0, 0, 0, time.UTC)},
		{DepartureTime: time.Date(2023, 1, 1, 21, 0, 0, 0, time.UTC)},
		{DepartureTime: time.Date(2023, 1, 1, 22, 0, 0, 0, time.UTC)},
		{DepartureTime: time.Date(2023, 1, 1, 23, 0, 0, 0, time.UTC)},
	}

	dataTest := []struct {
		name        string
		period      string
		tickets     []Ticket
		expected    int
		expectedErr error
	}{
		{
			name:     "Successfully get count by period - início da manhã",
			period:   "início da manhã",
			tickets:  tickets,
			expected: 2,
		},
		{
			name:     "Successfully get count by period - manhã",
			period:   "manhã",
			tickets:  tickets,
			expected: 5,
		},
		{
			name:     "Successfully get count by period - tarde",
			period:   "tarde",
			tickets:  tickets,
			expected: 7,
		},
		{
			name:     "Successfully get count by period - noite",
			period:   "noite",
			tickets:  tickets,
			expected: 4,
		},
		{
			name:        "Invalid period",
			period:      "invalid",
			tickets:     []Ticket{},
			expectedErr: &internal_error.ErrInvalidPeriod,
		},
	}

	for _, tt := range dataTest {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetCountByPeriod(tt.period, tt.tickets)

			require.Equal(t, tt.expected, result)
			if err != nil {
				require.Equal(t, tt.expectedErr.Error(), err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAverageDestination(t *testing.T) {
	tickets := []Ticket{
		{Destination: "Brazil"},
		{Destination: "Brazil"},
		{Destination: "USA"},
		{Destination: "Brazil"},
	}

	dataTest := []struct {
		name        string
		destination string
		tickets     []Ticket
		expected    float64
	}{
		{
			name:        "Successfully calculate average for Brazil",
			destination: "Brazil",
			tickets:     tickets,
			expected:    75,
		},
		{
			name:        "No tickets for destination",
			destination: "Canada",
			tickets:     tickets,
			expected:    0.0,
		},
	}

	for _, tt := range dataTest {
		t.Run(tt.name, func(t *testing.T) {
			result := AverageDestination(tt.destination, tt.tickets)
			if result != 0 {
				require.InEpsilon(t, tt.expected, result, 0.01)
			}
		})
	}
}
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
			period, err := stringToTicketPeriod(tt.input)
			if tt.expectError {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, period)
			}
		})
	}
}
