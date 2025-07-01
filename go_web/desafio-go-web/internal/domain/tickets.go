package domain

import (
	"strconv"
	"time"

	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/pkg/apperrors"
)

// Ticket represents the ticket structure that is read from the CSV file.
type Ticket struct {
	ID            int64
	Name          string
	Email         string
	Destination   string
	DepartureTime time.Time
	Price         float64
}

// StringsToTicket converts a slice of strings in the format <id,name,email,destination,departure_time,price> to a Ticket struct.
func StringsToTicket(s []string) (Ticket, *apperrors.HandledError) {
	id, err := strconv.ParseInt(s[0], 10, 64)
	if err != nil {
		return Ticket{}, &apperrors.ErrParsingTicket
	}

	departureTime, err := time.Parse("15:04", s[4])
	if err != nil {
		return Ticket{}, &apperrors.ErrParsingTicket
	}

	price, err := strconv.ParseFloat(s[5], 64)
	if err != nil {
		return Ticket{}, &apperrors.ErrParsingTicket
	}
	return Ticket{
		ID:            id,
		Name:          s[1],
		Email:         s[2],
		Destination:   s[3],
		DepartureTime: departureTime,
		Price:         price,
	}, nil
}

type TicketPeriod int

const (
	// EarlyMorning represents the period (0 → 6).
	EarlyMorning = iota
	// Morning represents the period (7 → 12).
	Morning
	// Afternoon represents the period (13 → 19).
	Afternoon
	// Evening represents the period (20 → 23).
	Evening
)

var ticketPeriodMap = map[string]TicketPeriod{
	"início da manhã": EarlyMorning,
	"manhã":           Morning,
	"tarde":           Afternoon,
	"noite":           Evening,
}

var TicketPeriodToTimeRange = map[TicketPeriod][2]int{
	EarlyMorning: {0, 6},
	Morning:      {7, 12},
	Afternoon:    {13, 19},
	Evening:      {20, 23},
}

func StringToTicketPeriod(s string) (TicketPeriod, error) {
	period, exists := ticketPeriodMap[s]
	if !exists {
		return -1, &apperrors.ErrInvalidPeriod
	}
	return period, nil
}

type TicketRepository interface {
	GetTotalTicketsByDestination(destination string) int
	GetCountByPeriod(period TicketPeriod) int
	AverageDestination(destination string) float64
}

type TicketService interface {
	GetTotalTicketsByDestination(destination string) int
	GetCountByPeriod(period TicketPeriod) int
	AverageDestination(destination string) float64
}

type TotalTicketsResponse struct {
	Total int `json:"totalTickets"`
}

type AverageTicketsResponse struct {
	Average float64 `json:"average"`
}
