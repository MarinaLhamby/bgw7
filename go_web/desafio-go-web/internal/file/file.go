package file

import (
	"encoding/csv"
	"os"

	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/internal/domain"
	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/pkg/apperrors"
)

const (
	filename = "../../internal/file/data/tickets.csv"
)

// GetTicketsFromFile reads tickets from a CSV file and returns a slice of Ticket structs.
func GetTicketsFromFile(strings ...string) ([]domain.Ticket, error) {
	var path string
	if len(strings) > 0 {
		path = strings[0]
	} else {
		path = filename
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, &apperrors.ErrReadingFile
	}
	defer file.Close()
	csv := csv.NewReader(file)

	records, err := csv.ReadAll()
	if err != nil {
		return nil, &apperrors.ErrReadingFile
	}

	var ticketsList []domain.Ticket
	for _, ticketRecord := range records {
		ticket, err := domain.StringsToTicket(ticketRecord)
		if err != nil {
			return nil, err
		}
		ticketsList = append(ticketsList, ticket)
	}

	return ticketsList, nil
}
