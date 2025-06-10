package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bgw7/exercicio_2_storage/internal/domain"
	"github.com/bgw7/exercicio_2_storage/internal/product"
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

func (h *ProductHandler) GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	products := h.productRepository.GetAllProducts()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) GetProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	idParam := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idParam, 10, 64)

	product := h.productRepository.GetProductById(id)
	if product == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) GetProductsByFilterPriceHandler(w http.ResponseWriter, r *http.Request) {
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

func (h *ProductHandler) PostProduct(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	response := domain.PostResponse{ID: id}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *ProductHandler) PutProductHandler(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedProduct)
}

func (h *ProductHandler) PatchHandler(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "error parsing url param", http.StatusBadRequest)
		return
	}

	err = h.productRepository.DeleteProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusNoContent)
}
