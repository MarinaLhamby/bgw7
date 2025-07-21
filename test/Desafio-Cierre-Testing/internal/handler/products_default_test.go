package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/mocks"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	url := "/product"
	t.Run("invalid id", func(t *testing.T) {
		repositoryMock := mocks.NewMockRepositoryProducts(ctrl)
		productHandler := handler.NewProductsDefault(repositoryMock)

		req := httptest.NewRequest(http.MethodGet, url+"?id=teste", nil)
		res := httptest.NewRecorder()
		expectedStatusCode := http.StatusBadRequest
		expectedResponse := `{"status":"Bad Request","message":"invalid id"}`
		repositoryMock.EXPECT().SearchProducts(gomock.Any()).Times(0)

		productHandler.Get().ServeHTTP(res, req)

		require.Equal(t, expectedResponse, res.Body.String())
		require.Equal(t, expectedStatusCode, res.Code)
	})

	t.Run("no id informed succesfully calls the search method", func(t *testing.T) {
		repositoryMock := mocks.NewMockRepositoryProducts(ctrl)
		productHandler := handler.NewProductsDefault(repositoryMock)

		req := httptest.NewRequest(http.MethodGet, url, nil)
		res := httptest.NewRecorder()
		expectedStatusCode := http.StatusOK
		methodReturn := map[int]internal.Product{
			1: {
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "description",
					Price:       10.0,
					SellerId:    1,
				},
			},
		}
		expectedResponse := `{"data":{"1":{"id":1,"description":"description","price":10,"seller_id":1}},"message":"success"}`

		repositoryMock.EXPECT().SearchProducts(internal.ProductQuery{
			Id: 0,
		}).Return(methodReturn, nil).Times(1)

		productHandler.Get().ServeHTTP(res, req)

		require.Equal(t, expectedResponse, res.Body.String())
		require.Equal(t, expectedStatusCode, res.Code)

	})

	t.Run("error ocurred while searching products", func(t *testing.T) {
		repositoryMock := mocks.NewMockRepositoryProducts(ctrl)
		productHandler := handler.NewProductsDefault(repositoryMock)

		req := httptest.NewRequest(http.MethodGet, url+"?id=1", nil)
		res := httptest.NewRecorder()
		expectedStatusCode := http.StatusOK
		methodReturn := map[int]internal.Product{
			1: {
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "description",
					Price:       10.0,
					SellerId:    1,
				},
			},
		}
		expectedResponse := `{"data":{"1":{"id":1,"description":"description","price":10,"seller_id":1}},"message":"success"}`

		repositoryMock.EXPECT().SearchProducts(internal.ProductQuery{
			Id: 1,
		}).Return(methodReturn, nil).Times(1)

		productHandler.Get().ServeHTTP(res, req)

		require.Equal(t, expectedResponse, res.Body.String())
		require.Equal(t, expectedStatusCode, res.Code)

	})

	t.Run("id informed succesfully calls the search method", func(t *testing.T) {
		repositoryMock := mocks.NewMockRepositoryProducts(ctrl)
		productHandler := handler.NewProductsDefault(repositoryMock)

		req := httptest.NewRequest(http.MethodGet, url+"?id=1", nil)
		res := httptest.NewRecorder()
		expectedStatusCode := http.StatusInternalServerError
		expectedResponse := `{"status":"Internal Server Error","message":"internal error"}`

		repositoryMock.EXPECT().SearchProducts(internal.ProductQuery{
			Id: 1,
		}).Return(make(map[int]internal.Product), sql.ErrNoRows).Times(1)

		productHandler.Get().ServeHTTP(res, req)

		require.Equal(t, expectedResponse, res.Body.String())
		require.Equal(t, expectedStatusCode, res.Code)

	})
}
