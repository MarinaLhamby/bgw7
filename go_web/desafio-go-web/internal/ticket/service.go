package ticket

import "github.com/MarinaLhamby/bgw7/desafio_de_fechamento/internal/domain"

type productService struct {
	ticketRepository domain.TicketRepository
}

func NewTicketService(repository domain.TicketRepository) *productService {
	return &productService{
		ticketRepository: repository,
	}
}

func (s *productService) GetTotalTicketsByDestination(destination string) int {
	return s.ticketRepository.GetTotalTicketsByDestination(destination)
}

func (s *productService) GetCountByPeriod(period domain.TicketPeriod) int {
	return s.ticketRepository.GetCountByPeriod(period)
}

func (s *productService) AverageDestination(destination string) float64 {
	return s.ticketRepository.AverageDestination(destination)
}
