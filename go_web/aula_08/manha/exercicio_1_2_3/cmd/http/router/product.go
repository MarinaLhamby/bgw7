package router

import (
	"net/http"

	"github.com/bgw7/exercicio_1_2_3/cmd/http/handler"
	"github.com/bgw7/exercicio_1_2_3/internal/product"
	"github.com/go-chi/chi/v5"
)

func buildProductsRoutes() http.Handler {
	r := chi.NewRouter()

	productRepository := product.NewProductJsonRepository()
	handler := handler.NewProductHandler(productRepository)

	r.Get("/", handler.GetAllProductsHandler)
	r.Get(`/{id:[0-9]+}`, handler.GetProductByIdHandler)
	r.Get("/search", handler.GetProductsByFilterPriceHandler)
	r.Post("/", handler.PostProduct)
	r.Put(`/{id:[0-9]+}`, handler.PutProductHandler)
	r.Patch(`/{id:[0-9]+}`, handler.PatchHandler)
	r.Delete(`/{id:[0-9]+}`, handler.DeleteHandler)
	return r
}
