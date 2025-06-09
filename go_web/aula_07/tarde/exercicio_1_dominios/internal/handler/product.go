package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bgw7/exercicio_1_dominios/internal/domain"
	"github.com/bgw7/exercicio_1_dominios/internal/product"
	"github.com/go-chi/chi/v5"
)

func GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	products := product.GetAllProducts()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func GetProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	idParam := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idParam, 10, 64)

	product := product.GetProductById(id)
	if product == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func GetProductsByFilterPriceHandler(w http.ResponseWriter, r *http.Request) {
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

	products := product.FilterByPrice(minPrice)
	if len(products) == 0 {
		http.Error(w, "No products found with the specified minimum price", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func PostProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var productReq domain.PostRequest
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

	id, err := product.AddProduct(productReq.ToProduct())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	response := domain.PostResponse{ID: id}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
