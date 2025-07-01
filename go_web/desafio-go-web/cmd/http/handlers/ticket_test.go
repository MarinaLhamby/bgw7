package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/cmd/http/handlers"
	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/internal/file"
	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/internal/ticket"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

const (
	pathToTestFile = "../../../internal/file/data/tickets_test.csv"
)

func TestTicketsByDestination(t *testing.T) {
	t.Run("get successfully", func(t *testing.T) {
		ticketsList, _ := file.GetTicketsFromFile(pathToTestFile)

		repository := ticket.NewTicketRepository(ticketsList)
		service := ticket.NewTicketService(repository)
		handler := handlers.NewTicketHandler(service)
		method := handler.GetTicketsByDestination()
		req := httptest.NewRequest(http.MethodGet, "/tickets/destination", nil)
		res := httptest.NewRecorder()
		expectedBody := `{"data":{"totalTickets":45}}`
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("destination", "Brazil")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		method.ServeHTTP(res, req)

		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, expectedBody, res.Body.String())
	})
}
func TestGetAverage(t *testing.T) {
	t.Run("get average successfully", func(t *testing.T) {
		ticketsList, _ := file.GetTicketsFromFile(pathToTestFile)

		repository := ticket.NewTicketRepository(ticketsList)
		service := ticket.NewTicketService(repository)
		handler := handlers.NewTicketHandler(service)
		method := handler.GetAverage()
		req := httptest.NewRequest(http.MethodGet, "/tickets/average", nil)
		res := httptest.NewRecorder()
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("destination", "Brazil")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		method.ServeHTTP(res, req)

		require.Equal(t, http.StatusOK, res.Code)
		expectedBody := `{"data":{"average":4.5}}`
		require.JSONEq(t, expectedBody, res.Body.String())
	})
}

func TestGetByPeriod(t *testing.T) {
	t.Run("get by period successfully", func(t *testing.T) {
		ticketsList, _ := file.GetTicketsFromFile(pathToTestFile)

		repository := ticket.NewTicketRepository(ticketsList)
		service := ticket.NewTicketService(repository)
		handler := handlers.NewTicketHandler(service)
		method := handler.GetByPeriod()
		req := httptest.NewRequest(http.MethodGet, "/tickets?period=tarde", nil)
		res := httptest.NewRecorder()

		method.ServeHTTP(res, req)

		require.Equal(t, http.StatusOK, res.Code)
		expectedBody := `{"data":{"totalTickets":289}}`
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("invalid period returns error", func(t *testing.T) {
		ticketsList, _ := file.GetTicketsFromFile(pathToTestFile)

		repository := ticket.NewTicketRepository(ticketsList)
		service := ticket.NewTicketService(repository)
		handler := handlers.NewTicketHandler(service)
		method := handler.GetByPeriod()
		req := httptest.NewRequest(http.MethodGet, "/tickets/period?period=invalid", nil)
		res := httptest.NewRecorder()

		method.ServeHTTP(res, req)

		require.Equal(t, http.StatusBadRequest, res.Code)
	})
}
