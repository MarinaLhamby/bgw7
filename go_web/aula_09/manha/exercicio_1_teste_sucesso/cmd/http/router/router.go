package router

import (
	"net/http"

	"github.com/bgw7/exercicio_1_teste_sucesso/pkg/middleware/auth"
	"github.com/go-chi/chi/v5"
)

type router struct {
}

func (router *router) MapRoutes() http.Handler {
	r := chi.NewRouter()
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})

	r.Group(func(r chi.Router) {
		r.Use(auth.Authenticate)
		r.Route("/products", func(rp chi.Router) {
			rp.Mount("/", buildProductsRoutes())
		})
		return
	})

	return r
}

func NewRouter() *router {
	return &router{}
}
