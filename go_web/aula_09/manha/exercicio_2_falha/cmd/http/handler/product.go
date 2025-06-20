package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/bgw7/exercicio_2_falha/internal/domain"
	"github.com/bgw7/exercicio_2_falha/internal/product"
	"github.com/bgw7/exercicio_2_falha/pkg/web/response"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	productRepository product.ProductRepository
}

func NewProductHandler(repo product.ProductRepository) *ProductHandler {
	return &ProductHandler{
		productRepository: repo,
	}
}

func (h *ProductHandler) GetAllProductsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		products := h.productRepository.GetAllProducts()

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
	}
}

func (h *ProductHandler) GetProductByIdHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		idParam := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			http.Error(w, "error parsing url param", http.StatusBadRequest)
			return
		}

		product := h.productRepository.GetProductById(id)
		if product == nil {
			err := response.ErrEntityNotFound.Format("product", idParam)
			http.Error(w, err.Message, err.StatusCode)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
	}

}

func (h *ProductHandler) GetProductsByFilterPriceHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		minPriceQuery := r.URL.Query().Get("priceGt")

		if minPriceQuery == "" {
			http.Error(w, "priceGt query parameter is required", http.StatusBadRequest)
			return
		}

		minPrice, err := strconv.ParseFloat(minPriceQuery, 64)
		if err != nil {
			http.Error(w, "Invalid priceGt value", http.StatusBadRequest)
			return
		}

		products := h.productRepository.FilterByPrice(minPrice)
		if len(products) == 0 {
			http.Error(w, "No products found with the specified minimum price", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
	}
}

func (h *ProductHandler) PostProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		var productReq domain.PostOrPutRequest
		err := json.NewDecoder(r.Body).Decode(&productReq)
		if err != nil {
			http.Error(w, "error processing the body", http.StatusUnprocessableEntity)
			return
		}

		err = productReq.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := h.productRepository.AddProduct(productReq.ToProduct())
		if err != nil {
			var apiError response.ApiError
			if errors.As(err, &apiError) {
				http.Error(w, err.Error(), apiError.StatusCode)
				return
			}
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		response := domain.PostResponse{ID: id}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (h *ProductHandler) PutProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "error parsing url param", http.StatusBadRequest)
			return
		}

		var requestBody domain.PostOrPutRequest
		err = json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, "error parsing body", http.StatusUnprocessableEntity)
			return
		}

		product := requestBody.ToProduct()
		product.ID = id

		updatedProduct, err := h.productRepository.UpdateProduct(product)
		if err != nil {
			var apiError response.ApiError
			if errors.As(err, &apiError) {
				http.Error(w, err.Error(), apiError.StatusCode)
				return
			}
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedProduct)
	}
}

func (h *ProductHandler) PatchHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			http.Error(w, "error parsing url param", http.StatusBadRequest)
			return
		}

		var partialProduct domain.PartialProduct
		err = json.NewDecoder(r.Body).Decode(&partialProduct)
		if err != nil {
			http.Error(w, "error parsing body", http.StatusUnprocessableEntity)
			return
		}

		product, err := h.productRepository.PartialUpdateProduct(id, partialProduct)
		if err != nil {
			var apiError response.ApiError
			if errors.As(err, &apiError) {
				http.Error(w, err.Error(), apiError.StatusCode)
				return
			}
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(product)
	}
}

func (h *ProductHandler) DeleteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			http.Error(w, "error parsing url param", http.StatusBadRequest)
			return
		}

		err = h.productRepository.DeleteProduct(id)
		if err != nil {
			var apiError response.ApiError
			if errors.As(err, &apiError) {
				http.Error(w, err.Error(), apiError.StatusCode)
				return
			}
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
