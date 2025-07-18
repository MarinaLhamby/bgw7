package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"app/utils"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSaveWithIdSale(t *testing.T) {
	db := sql.OpenDB(utils.Init())

	defer db.Close()
	saleRepository := repository.NewSalesMySQL(db)
	t.Run("Save sale with id succesfully", func(t *testing.T) {
		sale := internal.Sale{
			Id: 6,
			SaleAttributes: internal.SaleAttributes{
				Quantity:  1,
				ProductId: 1,
				InvoiceId: 1,
			},
		}

		err := saleRepository.SaveWithId(sale)
		require.Nil(t, err)
	})
}

func TestGetTotalByInvoiceIdSale(t *testing.T) {
	db := sql.OpenDB(utils.Init())

	defer db.Close()
	saleRepository := repository.NewSalesMySQL(db)
	t.Run("Get total by invoice id succesfully", func(t *testing.T) {
		expectedResponse := 100.0
		id := 1

		response, err := saleRepository.GetTotalByInvoiceId(id)
		require.Nil(t, err)
		require.Equal(t, expectedResponse, response)
	})
}
