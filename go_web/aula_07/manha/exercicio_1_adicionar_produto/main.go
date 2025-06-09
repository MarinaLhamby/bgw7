package main

import (
	"fmt"
	"net/http"

	"github.com/bgw7/exercicio_1_adicionar_produto/handlers"
	"github.com/bgw7/exercicio_1_adicionar_produto/json_data"
	"github.com/go-chi/chi/v5"
)

func main() {
	json_data.LoadProducts()
	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})
	r.Route("/products", func(r chi.Router) {
		r.Get("/", handlers.GetAllProductsHandler)
		r.Get(`/{id:[0-9]+}`, handlers.GetProductByIdHandler)
		r.Get("/search", handlers.GetProductsByFilterPriceHandler)
		r.Post("/", handlers.PostProduct)
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
