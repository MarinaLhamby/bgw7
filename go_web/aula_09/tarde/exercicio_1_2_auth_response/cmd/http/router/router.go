package router

import (
	"net/http"

	"github.com/bgw7/exercicio_1_2_auth_response/pkg/middleware/auth"
	"github.com/bgw7/exercicio_1_2_auth_response/pkg/middleware/request"
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
		r.Use(request.LogRequest)
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
