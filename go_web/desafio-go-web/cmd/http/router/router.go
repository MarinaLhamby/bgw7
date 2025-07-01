package router

import (
	"net/http"
	"time"

	_ "github.com/MarinaLhamby/bgw7/desafio_de_fechamento/docs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type router struct {
}

func (router *router) MapRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.StripSlashes,
		middleware.Timeout(5*time.Second),
		middleware.Heartbeat("/ping"),
	)

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Route("/tickets", func(rp chi.Router) {
		rp.Mount("/", buildTickesRoutes())
	})

	return r
}

func NewRouter() *router {
	return &router{}
}
