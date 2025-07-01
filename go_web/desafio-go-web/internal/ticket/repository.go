package ticket

import (
	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/internal/domain"
)

type ticketRepository struct {
	db []domain.Ticket
}

func NewTicketRepository(tickets []domain.Ticket) *ticketRepository {
	return &ticketRepository{
		db: tickets,
	}
}

// GetTotalticketsByDestination calculates the total number of tickets for a given destination.
func (r *ticketRepository) GetTotalTicketsByDestination(destination string) int {
	count := 0
	for _, ticket := range r.db {
		if ticket.Destination == destination {
			count++
		}
	}
	return count
}

// GetCountByPeriods calculates the total number os tickets in the informed period that can be <início da manhã, manhã, tarde, noite>.
func (r *ticketRepository) GetCountByPeriod(period domain.TicketPeriod) int {
	startingRange, endRange := domain.TicketPeriodToTimeRange[period][0], domain.TicketPeriodToTimeRange[period][1]

	count := 0
	for _, ticket := range r.db {
		if ticket.DepartureTime.Hour() >= startingRange && ticket.DepartureTime.Hour() <= endRange {
			count++
		}
	}
	return count
}

// AverageDestination calculates the average number of tickets for a given destination.
func (r *ticketRepository) AverageDestination(destination string) float64 {
	total := len(r.db)
	if total == 0 {
		return 0
	}
	countDestination := r.GetTotalTicketsByDestination(destination)

	average := (float64(countDestination) / float64(total)) * 100
	return average
}
