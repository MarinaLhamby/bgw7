package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/bgw7/exercicio_2_falha/internal/domain"
	"github.com/bgw7/exercicio_2_falha/internal/product"
	"github.com/bgw7/exercicio_2_falha/internal/storage"
	"github.com/bgw7/exercicio_2_falha/pkg/middleware/auth"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestGetAllProductsHandler(t *testing.T) {
	t.Run("success getting all products", func(t *testing.T) {
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

		handler.GetAllProductsHandler().ServeHTTP(res, req)

		require.Equal(t, http.StatusOK, res.Code)
		require.JSONEq(t, expectedResponse, res.Body.String())
	})
}

func TestGetProductByIdHandler(t *testing.T) {
	t.Run("success getting a single product by id", func(t *testing.T) {
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

		handler.GetProductByIdHandler().ServeHTTP(res, req)

		require.Equal(t, http.StatusOK, res.Code)
		require.JSONEq(t, expectedResponse, res.Body.String())
	})

	t.Run("bad request by passing invalid id", func(t *testing.T) {
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

		expectedResponse := "error parsing url param\n"
		req := httptest.NewRequest("GET", "/products", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "AX@@@@@")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()

		storage := storage.NewStorageMock()
		repo := product.NewProductJsonRepository(storage, products)
		handler := NewProductHandler(repo)

		handler.GetProductByIdHandler().ServeHTTP(res, req)

		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, expectedResponse, res.Body.String())
	})

	t.Run("not found", func(t *testing.T) {
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

		expectedResponse := "resource product of id 105 not found\n"
		req := httptest.NewRequest("GET", "/products", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "105")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()

		storage := storage.NewStorageMock()
		repo := product.NewProductJsonRepository(storage, products)
		handler := NewProductHandler(repo)

		handler.GetProductByIdHandler().ServeHTTP(res, req)

		require.Equal(t, http.StatusNotFound, res.Code)
		require.Equal(t, expectedResponse, res.Body.String())
	})
}

func TestPostProductHandler(t *testing.T) {
	t.Run("success adding a new product", func(t *testing.T) {
		productStr, _ := json.Marshal(domain.Product{
			Name:        "Dell Monitor",
			Quantity:    30,
			CodeValue:   "Dell-2025-B",
			IsPublished: true,
			Expiration:  "12/10/2026",
			Price:       1000.00,
		})
		expectedResponse := "{\"id\":1}"

		req := httptest.NewRequest("POST", "/products", strings.NewReader(string(productStr)))
		res := httptest.NewRecorder()

		storage := storage.NewStorageMock()
		repo := product.NewProductJsonRepository(storage, []domain.Product{})
		handler := NewProductHandler(repo)

		handler.PostProductHandler().ServeHTTP(res, req)

		require.Equal(t, http.StatusCreated, res.Code)
		require.JSONEq(t, expectedResponse, res.Body.String())
	})

	t.Run("bad request by passing invalid json", func(t *testing.T) {
		productStr, _ := json.Marshal(domain.Product{})

		expectedResponse := "name should be informed\n"

		req := httptest.NewRequest("POST", "/products", strings.NewReader(string(productStr)))
		res := httptest.NewRecorder()

		storage := storage.NewStorageMock()
		repo := product.NewProductJsonRepository(storage, []domain.Product{})
		handler := NewProductHandler(repo)

		handler.PostProductHandler().ServeHTTP(res, req)

		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, expectedResponse, res.Body.String())
	})

	t.Run("unauthorized", func(t *testing.T) {
		productStr, _ := json.Marshal(domain.Product{})
		os.Setenv("TOKEN", "TOKEN_TEST")
		expectedResponse := "invalid token\n"

		req := httptest.NewRequest("POST", "/products", strings.NewReader(string(productStr)))
		res := httptest.NewRecorder()

		storage := storage.NewStorageMock()
		repo := product.NewProductJsonRepository(storage, []domain.Product{})
		handler := NewProductHandler(repo)

		middleware := auth.Authenticate(handler.PostProductHandler())
		middleware.ServeHTTP(res, req)

		require.Equal(t, http.StatusUnauthorized, res.Code)
		require.Equal(t, expectedResponse, res.Body.String())
	})
}

func TestDeleteHandler(t *testing.T) {
	t.Run("successfully delete product", func(t *testing.T) {
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

		handler.DeleteHandler().ServeHTTP(res, req)

		require.Equal(t, http.StatusNoContent, res.Code)
		require.Equal(t, "", res.Body.String())
	})

	t.Run("bad request by passing invalid id when delete product", func(t *testing.T) {
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
		expectedResponse := "error parsing url param\n"
		req := httptest.NewRequest("DELETE", "/products", nil)
		ctxChi := chi.NewRouteContext()
		ctxChi.URLParams.Add("id", "XXX0")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctxChi))
		res := httptest.NewRecorder()

		storage := storage.NewStorageMock()
		repo := product.NewProductJsonRepository(storage, products)
		handler := NewProductHandler(repo)

		handler.DeleteHandler().ServeHTTP(res, req)

		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, expectedResponse, res.Body.String())
	})

	t.Run("unauthorized", func(t *testing.T) {
		expectedResponse := "invalid token\n"
		req := httptest.NewRequest("DELETE", "/products", nil)

		res := httptest.NewRecorder()

		storage := storage.NewStorageMock()
		repo := product.NewProductJsonRepository(storage, []domain.Product{})
		handler := NewProductHandler(repo)
		middlaware := auth.Authenticate(handler.DeleteHandler())
		middlaware.ServeHTTP(res, req)

		require.Equal(t, http.StatusUnauthorized, res.Code)
		require.Equal(t, expectedResponse, res.Body.String())
	})

	t.Run("not found", func(t *testing.T) {
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
		expectedResponse := "resource product of id 105 not found\n"
		req := httptest.NewRequest("DELETE", "/products", nil)
		ctxChi := chi.NewRouteContext()
		ctxChi.URLParams.Add("id", "105")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctxChi))
		res := httptest.NewRecorder()

		storage := storage.NewStorageMock()
		repo := product.NewProductJsonRepository(storage, products)
		handler := NewProductHandler(repo)

		handler.DeleteHandler().ServeHTTP(res, req)

		require.Equal(t, http.StatusNotFound, res.Code)
		require.Equal(t, expectedResponse, res.Body.String())
	})
}

func TestPutProductHandler(t *testing.T) {
	t.Run("bad request when parsing id", func(t *testing.T) {
		expectedError := "error parsing url param\n"

		str := storage.NewStorageMock()
		repo := product.NewProductJsonRepository(str, []domain.Product{})
		h := NewProductHandler(repo)

		req := httptest.NewRequest("PUT", "/products", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "A00000")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()

		h.PutProductHandler().ServeHTTP(res, req)

		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, expectedError, res.Body.String())
	})

	t.Run("unauthorized", func(t *testing.T) {
		expectedResponse := "invalid token\n"

		str := storage.NewStorageMock()
		repo := product.NewProductJsonRepository(str, []domain.Product{})
		h := NewProductHandler(repo)

		req := httptest.NewRequest("PUT", "/products", nil)
		res := httptest.NewRecorder()
		middlaware := auth.Authenticate(h.PutProductHandler())

		middlaware.ServeHTTP(res, req)

		require.Equal(t, http.StatusUnauthorized, res.Code)
		require.Equal(t, expectedResponse, res.Body.String())
	})
}

func TestPatchHandler(t *testing.T) {
	t.Run("bad request when parsing id", func(t *testing.T) {
		expectedError := "error parsing url param\n"

		str := storage.NewStorageMock()
		repo := product.NewProductJsonRepository(str, []domain.Product{})
		h := NewProductHandler(repo)

		req := httptest.NewRequest("PATCH", "/products", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "A00000")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()

		h.PatchHandler().ServeHTTP(res, req)

		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, expectedError, res.Body.String())
	})

	t.Run("unauthorized", func(t *testing.T) {
		expectedResponse := "invalid token\n"

		str := storage.NewStorageMock()
		repo := product.NewProductJsonRepository(str, []domain.Product{})
		h := NewProductHandler(repo)

		req := httptest.NewRequest("PATCH", "/products", nil)
		res := httptest.NewRecorder()
		middleware := auth.Authenticate(h.PatchHandler())

		middleware.ServeHTTP(res, req)

		require.Equal(t, http.StatusUnauthorized, res.Code)
		require.Equal(t, expectedResponse, res.Body.String())
	})

	t.Run("not found", func(t *testing.T) {
		expectedResponse := "resource product of id 105 not found\n"

		str := storage.NewStorageMock()
		repo := product.NewProductJsonRepository(str, []domain.Product{})
		h := NewProductHandler(repo)
		partialProduct, _ := json.Marshal(domain.PartialProduct{
			Name: "another name",
		})

		req := httptest.NewRequest("PATCH", "/products", strings.NewReader(string(partialProduct)))
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "105")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()

		h.PatchHandler().ServeHTTP(res, req)

		require.Equal(t, http.StatusNotFound, res.Code)
		require.Equal(t, expectedResponse, res.Body.String())
	})
}
