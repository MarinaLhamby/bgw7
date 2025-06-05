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
			expected:    233.33,
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
			require.InEpsilon(t, tt.expected, result, 0.01)
		})
	}
}
