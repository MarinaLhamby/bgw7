package main

import (
	"fmt"

	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/internal/file"
	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/internal/tickets"
)

func main() {
	ticketsList, err := file.GetTicketsFromFile()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	total := tickets.GetTotalTicketsByDestination("Brazil", ticketsList)
	average := tickets.AverageDestination("Brazil", ticketsList)
	morningTickets, err := tickets.GetCountByPeriod("início da manhã", ticketsList)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Total tickets for Brazil:", total)
	fmt.Printf("Average of tickets for Brazil: %.2f%%\n", average)
	fmt.Println("Morning tickets:", morningTickets)
}
