package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"app/utils"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSaveWithIdProduct(t *testing.T) {
	db := sql.OpenDB(utils.Init())

	defer db.Close()
	productRepository := repository.NewProductsMySQL(db)
	t.Run("Save product with id succesfully", func(t *testing.T) {
		product := internal.Product{
			Id: 6,
			ProductAttributes: internal.ProductAttributes{
				Description: "Descrição",
				Price:       10,
			},
		}

		err := productRepository.SaveWithId(product)
		require.Nil(t, err)
	})
}

func TestGetTopFiveBestSellersProduct(t *testing.T) {
	db := sql.OpenDB(utils.Init())

	defer db.Close()
	productRepository := repository.NewProductsMySQL(db)
	t.Run("Get best sellers succesfully", func(t *testing.T) {
		expectedResponse := []internal.BestSellers{
			{
				Description: "Product 8",
				Total:       8,
			},
			{
				Description: "Product 7",
				Total:       7,
			},
			{
				Description: "Product 6",
				Total:       6,
			},
			{
				Description: "Product 5",
				Total:       5,
			},
			{
				Description: "Product 4",
				Total:       4,
			},
		}

		response, err := productRepository.GetTopFiveBestSellers()
		require.Nil(t, err)
		require.Equal(t, expectedResponse, response)
	})
}
