package main

import (
	"fmt"
	"net/http"

	"github.com/bgw7/exercicio_1_dominios/internal/handler"
	"github.com/bgw7/exercicio_1_dominios/internal/product"
	"github.com/go-chi/chi/v5"
)

func main() {
	product.LoadProducts()
	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})
	r.Route("/products", func(r chi.Router) {
		r.Get("/", handler.GetAllProductsHandler)
		r.Get(`/{id:[0-9]+}`, handler.GetProductByIdHandler)
		r.Get("/search", handler.GetProductsByFilterPriceHandler)
		r.Post("/", handler.PostProduct)
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
