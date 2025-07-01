package router

import (
	"net/http"

	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/cmd/http/handlers"
	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/internal/file"
	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/internal/ticket"
	"github.com/go-chi/chi/v5"
)

func buildTickesRoutes() http.Handler {
	r := chi.NewRouter()

	ticketsList, err := file.GetTicketsFromFile()
	if err != nil {
		panic(err)
	}

	ticketRepository := ticket.NewTicketRepository(ticketsList)
	ticketService := ticket.NewTicketService(ticketRepository)

	handler := handlers.NewTicketHandler(ticketService)
	r.Get("/destination/{destination}", handler.GetTicketsByDestination())
	r.Get("/destination/average/{destination}", handler.GetAverage())
	r.Get("/", handler.GetByPeriod())
	return r
}
