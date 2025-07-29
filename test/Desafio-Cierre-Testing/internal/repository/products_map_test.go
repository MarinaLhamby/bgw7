package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchProducts(t *testing.T) {
	db := map[int]internal.Product{
		101: {
			Id: 101,
			ProductAttributes: internal.ProductAttributes{
				Description: "Laptop Pro X",
				Price:       1200.50,
				SellerId:    1001,
			},
		},
		102: {
			Id: 102,
			ProductAttributes: internal.ProductAttributes{
				Description: "Wireless Mouse Ergo",
				Price:       25.99,
				SellerId:    1002,
			},
		},
		103: {
			Id: 103,
			ProductAttributes: internal.ProductAttributes{
				Description: "Mechanical Keyboard RGB",
				Price:       89.00,
				SellerId:    1001, // Same seller as Laptop Pro X
			},
		},
		104: {
			Id: 104,
			ProductAttributes: internal.ProductAttributes{
				Description: "4K Monitor UltraView",
				Price:       350.00,
				SellerId:    1003,
			},
		},
		105: {
			Id: 105,
			ProductAttributes: internal.ProductAttributes{
				Description: "USB-C Hub Multiport",
				Price:       45.75,
				SellerId:    1002,
			},
		},
		106: {
			Id: 106,
			ProductAttributes: internal.ProductAttributes{
				Description: "External SSD 1TB",
				Price:       99.99,
				SellerId:    1004,
			},
		},
	}

	t.Run("get all products succesfully when no query is informed", func(t *testing.T) {
		repo := repository.NewProductsMap(db)

		response, err := repo.SearchProducts(internal.ProductQuery{})

		require.NoError(t, err)
		require.Equal(t, db, response)
	})

	t.Run("get all products succesfully when query if informed", func(t *testing.T) {
		repo := repository.NewProductsMap(db)

		response, err := repo.SearchProducts(internal.ProductQuery{
			Id: 101,
		})

		require.NoError(t, err)
		require.Equal(t, map[int]internal.Product{101: db[101]}, response)
	})

	t.Run("get all when there`s no products", func(t *testing.T) {
		repo := repository.NewProductsMap(make(map[int]internal.Product))

		response, err := repo.SearchProducts(internal.ProductQuery{
			Id: 101,
		})

		require.NoError(t, err)
		require.Empty(t, response)
	})
}
