package ticket_test

import (
	"testing"

	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/internal/domain"
	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/internal/file"
	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/internal/ticket"
	"github.com/stretchr/testify/require"
)

const (
	pathToTestFile = "../file/data/tickets_test.csv"
)

func TestGetTotalTicketsByDestination(t *testing.T) {
	t.Run("get by destination successfully", func(t *testing.T) {
		ticketsList, _ := file.GetTicketsFromFile(pathToTestFile)

		repository := ticket.NewTicketRepository(ticketsList)
		service := ticket.NewTicketService(repository)
		expectedResponse := 45

		response := service.GetTotalTicketsByDestination("Brazil")

		require.Equal(t, expectedResponse, response)
	})
}
func TestGetCountByPeriod(t *testing.T) {
	t.Run("get count by period successfully", func(t *testing.T) {
		ticketsList, _ := file.GetTicketsFromFile(pathToTestFile)

		repository := ticket.NewTicketRepository(ticketsList)
		service := ticket.NewTicketService(repository)
		period, _ := domain.StringToTicketPeriod("início da manhã")
		expectedCount := 304

		response := service.GetCountByPeriod(period)

		require.Equal(t, expectedCount, response)
	})
}

func TestAverageDestination(t *testing.T) {
	t.Run("get average by destination successfully", func(t *testing.T) {
		ticketsList, _ := file.GetTicketsFromFile(pathToTestFile)

		repository := ticket.NewTicketRepository(ticketsList)
		service := ticket.NewTicketService(repository)
		destination := "Brazil"
		expectedAverage := 4.5

		response := service.AverageDestination(destination)

		require.InDelta(t, expectedAverage, response, 0.0001)
	})
}
