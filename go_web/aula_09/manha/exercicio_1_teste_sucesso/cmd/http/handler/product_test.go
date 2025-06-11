package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bgw7/exercicio_1_teste_sucesso/internal/domain"
	"github.com/bgw7/exercicio_1_teste_sucesso/internal/product"
	"github.com/bgw7/exercicio_1_teste_sucesso/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestGetAllProductsHandler(t *testing.T) {
	products := []domain.Product{
		{
			ID:          101,
			Name:        "Wireless Mouse",
			Quantity:    50,
			CodeValue:   "WM-2023-A",
			IsPublished: true,
			Expiration:  "10/10/2025",
			Price:       25.99,
		},
		{
			ID:          102,
			Name:        "Mechanical Keyboard",
			Quantity:    30,
			CodeValue:   "MK-2023-B",
			IsPublished: false,
			Expiration:  "12/10/2020",
			Price:       75.00,
		},
	}
	expectedResponse := "[{\"id\":101,\"name\":\"Wireless Mouse\",\"quantity\":50,\"code_value\":\"WM-2023-A\",\"is_published\":true,\"expiration\":\"10/10/2025\",\"price\":25.99},{\"id\":102,\"name\":\"Mechanical Keyboard\",\"quantity\":30,\"code_value\":\"MK-2023-B\",\"is_published\":false,\"expiration\":\"12/10/2020\",\"price\":75}]\n"
	req := httptest.NewRequest("GET", "/products", nil)
	res := httptest.NewRecorder()

	storage := storage.NewStorageMock()
	repo := product.NewProductJsonRepository(storage, products)
	handler := NewProductHandler(repo)

	handler.GetAllProductsHandler(res, req)

	require.Equal(t, http.StatusOK, res.Code)
	require.JSONEq(t, expectedResponse, res.Body.String())
}

func TestGetProductByIdHandler(t *testing.T) {
	products := []domain.Product{
		{
			ID:          101,
			Name:        "Wireless Mouse",
			Quantity:    50,
			CodeValue:   "WM-2023-A",
			IsPublished: true,
			Expiration:  "10/10/2025",
			Price:       25.99,
		},
		{
			ID:          102,
			Name:        "Mechanical Keyboard",
			Quantity:    30,
			CodeValue:   "MK-2023-B",
			IsPublished: false,
			Expiration:  "12/10/2020",
			Price:       75.00,
		},
	}

	expectedResponse := "{\"id\":101,\"name\":\"Wireless Mouse\",\"quantity\":50,\"code_value\":\"WM-2023-A\",\"is_published\":true,\"expiration\":\"10/10/2025\",\"price\":25.99}"
	req := httptest.NewRequest("GET", "/products/101", nil)
	chiCtx := chi.NewRouteContext()
	chiCtx.URLParams.Add("id", "101")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
	res := httptest.NewRecorder()

	storage := storage.NewStorageMock()
	repo := product.NewProductJsonRepository(storage, products)
	handler := NewProductHandler(repo)

	handler.GetProductByIdHandler(res, req)

	require.Equal(t, http.StatusOK, res.Code)
	require.JSONEq(t, expectedResponse, res.Body.String())
}

func TestPostProductHandler(t *testing.T) {
	products := []domain.Product{
		{
			ID:          101,
			Name:        "Wireless Mouse",
			Quantity:    50,
			CodeValue:   "WM-2023-A",
			IsPublished: true,
			Expiration:  "10/10/2025",
			Price:       25.99,
		},
		{
			ID:          102,
			Name:        "Mechanical Keyboard",
			Quantity:    30,
			CodeValue:   "MK-2023-B",
			IsPublished: false,
			Expiration:  "12/10/2020",
			Price:       75.00,
		},
	}

	productStr, _ := json.Marshal(domain.Product{
		Name:        "Dell Monitor",
		Quantity:    30,
		CodeValue:   "Dell-2025-B",
		IsPublished: true,
		Expiration:  "12/10/2026",
		Price:       1000.00,
	})
	expectedResponse := "{\"id\":103}"

	req := httptest.NewRequest("POST", "/products", strings.NewReader(string(productStr)))
	res := httptest.NewRecorder()

	storage := storage.NewStorageMock()
	repo := product.NewProductJsonRepository(storage, products)
	handler := NewProductHandler(repo)

	handler.PostProductHandler(res, req)

	require.Equal(t, http.StatusCreated, res.Code)
	require.JSONEq(t, expectedResponse, res.Body.String())
}

func TestPutProductHandler(t *testing.T) {
	products := []domain.Product{
		{
			ID:          101,
			Name:        "Wireless Mouse",
			Quantity:    50,
			CodeValue:   "WM-2023-A",
			IsPublished: true,
			Expiration:  "10/10/2025",
			Price:       25.99,
		},
		{
			ID:          102,
			Name:        "Mechanical Keyboard",
			Quantity:    30,
			CodeValue:   "MK-2023-B",
			IsPublished: false,
			Expiration:  "12/10/2020",
			Price:       75.00,
		},
	}

	req := httptest.NewRequest("DELETE", "/products", nil)
	ctxChi := chi.NewRouteContext()
	ctxChi.URLParams.Add("id", "101")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctxChi))
	res := httptest.NewRecorder()

	storage := storage.NewStorageMock()
	repo := product.NewProductJsonRepository(storage, products)
	handler := NewProductHandler(repo)

	handler.DeleteHandler(res, req)

	require.Equal(t, http.StatusNoContent, res.Code)
	require.Equal(t, "", res.Body.String())
}
