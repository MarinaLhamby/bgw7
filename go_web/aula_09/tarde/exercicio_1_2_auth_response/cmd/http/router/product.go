package router

import (
	"net/http"

	"github.com/bgw7/exercicio_1_2_auth_response/cmd/http/handler"
	"github.com/bgw7/exercicio_1_2_auth_response/internal/product"
	"github.com/bgw7/exercicio_1_2_auth_response/internal/storage"
	"github.com/go-chi/chi/v5"
)

func buildProductsRoutes() http.Handler {
	r := chi.NewRouter()

	str := storage.NewStorageJson()
	products, err := str.Get()
	if err != nil {
		panic(err)
	}
	productRepository := product.NewProductJsonRepository(str, products)
	handler := handler.NewProductHandler(productRepository)

	r.Get("/", handler.GetAllProductsHandler())
	r.Get(`/{id:[0-9]+}`, handler.GetProductByIdHandler())
	r.Get("/search", handler.GetProductsByFilterPriceHandler())
	r.Post("/", handler.PostProductHandler())
	r.Put(`/{id:[0-9]+}`, handler.PutProductHandler())
	r.Patch(`/{id:[0-9]+}`, handler.PatchHandler())
	r.Delete(`/{id:[0-9]+}`, handler.DeleteHandler())
	return r
}
