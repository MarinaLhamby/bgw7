package ticket

import (
	"testing"
	"time"

	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/internal/domain"
	"github.com/stretchr/testify/require"
)

func TestGetTotalTicketsByDestination(t *testing.T) {
	dataTest := []struct {
		name        string
		destination string
		tickets     []domain.Ticket
		expected    int
	}{
		{
			name:        "Succesfully get total tickets by destination",
			destination: "Brazil",
			tickets: []domain.Ticket{
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
			tickets: []domain.Ticket{
				{Destination: "USA"},
				{Destination: "Canada"},
			},
			expected: 0,
		},
		{
			name:        "Empty tickets slice",
			destination: "Brazil",
			tickets:     []domain.Ticket{},
			expected:    0,
		},
	}
	for _, tt := range dataTest {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewTicketRepository(tt.tickets)

			result := repo.GetTotalTicketsByDestination(tt.destination)

			require.Equal(t, tt.expected, result)
		})
	}
}

func TestGetCountByPeriod(t *testing.T) {
	tickets := []domain.Ticket{
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

	repo := NewTicketRepository(tickets)

	dataTest := []struct {
		name        string
		period      domain.TicketPeriod
		tickets     []domain.Ticket
		expected    int
		expectedErr error
	}{
		{
			name:     "Successfully get count by period - início da manhã",
			period:   domain.EarlyMorning,
			tickets:  tickets,
			expected: 2,
		},
		{
			name:     "Successfully get count by period - manhã",
			period:   domain.Morning,
			tickets:  tickets,
			expected: 5,
		},
		{
			name:     "Successfully get count by period - tarde",
			period:   domain.Afternoon,
			tickets:  tickets,
			expected: 7,
		},
		{
			name:     "Successfully get count by period - noite",
			period:   domain.Evening,
			tickets:  tickets,
			expected: 4,
		},
	}

	for _, tt := range dataTest {
		t.Run(tt.name, func(t *testing.T) {
			result := repo.GetCountByPeriod(tt.period)

			require.Equal(t, tt.expected, result)
		})
	}
}

func TestAverageDestination(t *testing.T) {
	tickets := []domain.Ticket{
		{Destination: "Brazil"},
		{Destination: "Brazil"},
		{Destination: "USA"},
		{Destination: "Brazil"},
	}

	repo := NewTicketRepository(tickets)

	dataTest := []struct {
		name        string
		destination string
		tickets     []domain.Ticket
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
			result := repo.AverageDestination(tt.destination)
			if result != 0 {
				require.InEpsilon(t, tt.expected, result, 0.01)
			}
		})
	}
}
