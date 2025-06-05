package tickets

import (
	"strconv"
	"time"

	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/internal/internal_error"
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
func StringsToTicket(s []string) (Ticket, *internal_error.HandledError) {
	id, err := strconv.ParseInt(s[0], 10, 64)
	if err != nil {
		return Ticket{}, &internal_error.ErrParsingTicket
	}

	departureTime, err := time.Parse("15:04", s[4])
	if err != nil {
		return Ticket{}, &internal_error.ErrParsingTicket
	}

	price, err := strconv.ParseFloat(s[5], 64)
	if err != nil {
		return Ticket{}, &internal_error.ErrParsingTicket
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

var ticketPeriodToTimeRange = map[TicketPeriod][2]int{
	EarlyMorning: {0, 6},
	Morning:      {7, 12},
	Afternoon:    {13, 19},
	Evening:      {20, 23},
}

func stringToTicketPeriod(s string) (TicketPeriod, error) {
	period, exists := ticketPeriodMap[s]
	if !exists {
		return -1, &internal_error.ErrInvalidPeriod
	}
	return period, nil
}

// GetTotalticketsByDestination calculates the total number of tickets for a given destination.
func GetTotalTicketsByDestination(destination string, tickets []Ticket) int {
	count := 0
	for _, ticket := range tickets {
		if ticket.Destination == destination {
			count++
		}
	}
	return count
}

// GetCountByPeriods calculates the total number os tickets in the informed period that can be <início da manhã, manhã, tarde, noite>.
func GetCountByPeriod(time string, tickets []Ticket) (int, error) {
	parsedPeriod, err := stringToTicketPeriod(time)
	if err != nil {
		return 0, err
	}
	startingRange, endRange := ticketPeriodToTimeRange[parsedPeriod][0], ticketPeriodToTimeRange[parsedPeriod][1]

	count := 0
	for _, ticket := range tickets {
		if ticket.DepartureTime.Hour() >= startingRange && ticket.DepartureTime.Hour() <= endRange {
			count++
		}
	}
	return count, nil
}

// AverageDestination calculates the average number of tickets for a given destination.
func AverageDestination(destination string, tickets []Ticket) float64 {
	total := len(tickets)
	if total == 0 {
		return 0
	}
	countDestination := GetTotalTicketsByDestination(destination, tickets)

	average := (float64(countDestination) / float64(total)) * 100
	return average
}
