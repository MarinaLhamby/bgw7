package file

import (
	"encoding/csv"
	"os"

	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/internal/internal_error"
	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/internal/tickets"
)

const (
	filename = "tickets.csv"
)

// GetTicketsFromFile reads tickets from a CSV file and returns a slice of Ticket structs.
func GetTicketsFromFile(strings ...string) ([]tickets.Ticket, error) {
	var path string
	if len(strings) > 0 {
		path = strings[0]
	} else {
		path = filename
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, &internal_error.ErrReadingFile
	}
	defer file.Close()
	csv := csv.NewReader(file)

	records, err := csv.ReadAll()
	if err != nil {
		return nil, &internal_error.ErrReadingFile
	}

	var ticketsList []tickets.Ticket
	for _, ticketRecord := range records {
		ticket, err := tickets.StringsToTicket(ticketRecord)
		if err != nil {
			return nil, err
		}
		ticketsList = append(ticketsList, ticket)
	}

	return ticketsList, nil
}
