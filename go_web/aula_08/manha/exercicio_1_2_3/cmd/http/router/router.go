package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type router struct {
}

func (router *router) MapRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})

	r.Route("/products", func(rp chi.Router) {
		rp.Mount("/", buildProductsRoutes())
	})
	return r
}

func NewRouter() *router {
	return &router{}
}
