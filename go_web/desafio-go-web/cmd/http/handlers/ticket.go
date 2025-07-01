package handlers

import (
	"net/http"

	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/internal/domain"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

type TicketHandler struct {
	service domain.TicketService
}

func NewTicketHandler(service domain.TicketService) *TicketHandler {
	return &TicketHandler{
		service: service,
	}
}

// @Summary Fetch number of tickets by destination
// @Description Fetch number of tickets by destination
// @Tags tickets
// @Accept  json
// @Produce  json
// @Param destination path string true "Destination"
// @Success 200 {object} map[string]domain.TotalTicketsResponse
// @Router /tickets/destination/{destination} [get]
func (h *TicketHandler) GetTicketsByDestination() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		destination := chi.URLParam(r, "destination")

		total := h.service.GetTotalTicketsByDestination(destination)

		response.JSON(w, http.StatusOK, map[string]any{
			"data": domain.TotalTicketsResponse{
				Total: total,
			},
		})
	}
}

// @Summary Fetch average number of tickets by destination
// @Description Fetch average number of tickets by destination
// @Tags tickets
// @Accept  json
// @Produce  json
// @Param destination path string true "Destination"
// @Success 200 {object} map[string]domain.AverageTicketsResponse
// @Router /tickets/destination/average/{destination} [get]
func (h *TicketHandler) GetAverage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		destination := chi.URLParam(r, "destination")

		average := h.service.AverageDestination(destination)

		response.JSON(w, http.StatusOK, map[string]any{
			"data": domain.AverageTicketsResponse{
				Average: average,
			},
		})
	}
}

// @Summary Fetch total number of tickets by period
// @Description Fetch total number of tickets by period
// @Tags tickets
// @Accept  json
// @Produce  json
// @Param period query string true "início da manhã, manhã, tarde ou noite"
// @Success 200 {object} map[string]domain.TotalTicketsResponse
// @failure      400              {string}  string    "error"
// @Router /tickets [get]
func (h *TicketHandler) GetByPeriod() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		timeStr := r.URL.Query().Get("period")
		time, err := domain.StringToTicketPeriod(timeStr)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
		}

		total := h.service.GetCountByPeriod(time)

		response.JSON(w, http.StatusOK, map[string]any{
			"data": domain.TotalTicketsResponse{
				Total: total,
			},
		})
	}
}
