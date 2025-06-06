package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bgw7/exercicio_2_criar_servidor/json_data"
	"github.com/go-chi/chi/v5"
)

func GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	products := json_data.GetProducts()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func GetProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	idParam := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idParam, 10, 64)

	product := json_data.GetProductById(id)
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

	products := json_data.FilterByPrice(minPrice)
	if len(products) == 0 {
		http.Error(w, "No products found with the specified minimum price", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
