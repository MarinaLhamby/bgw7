package router

import (
	"net/http"

	"github.com/bgw7/exercicio_3_resposta/cmd/http/handler"
	"github.com/bgw7/exercicio_3_resposta/internal/product"
	"github.com/bgw7/exercicio_3_resposta/internal/storage"
	"github.com/go-chi/chi/v5"
)

func buildProductsRoutes() http.Handler {
	r := chi.NewRouter()

	productRepository := product.NewProductJsonRepository(storage.NewStorageJson())
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
