package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"app/utils"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSaveWithIdInvoice(t *testing.T) {
	db := sql.OpenDB(utils.Init())

	defer db.Close()
	invoiceRepository := repository.NewInvoicesMySQL(db)
	t.Run("Save invoice with id succesfully", func(t *testing.T) {
		invoice := internal.Invoice{
			Id: 6,
			InvoiceAttributes: internal.InvoiceAttributes{
				Datetime:   "2022-08-08",
				Total:      450,
				CustomerId: 1,
			},
		}

		err := invoiceRepository.SaveWithId(invoice)
		require.Nil(t, err)
	})
}

func TestUpdateTotalByIdInvoice(t *testing.T) {
	db := sql.OpenDB(utils.Init())

	defer db.Close()
	invoiceRepository := repository.NewInvoicesMySQL(db)
	t.Run("Update total invoice succesfully", func(t *testing.T) {
		total := 1.0
		id := 1

		err := invoiceRepository.UpdateTotalById(id, total)
		require.Nil(t, err)
	})
}

func TestGetTotalAndCustomerConditionInvoice(t *testing.T) {
	db := sql.OpenDB(utils.Init())

	defer db.Close()
	invoiceRepository := repository.NewInvoicesMySQL(db)
	t.Run("Get total and customer condition invoice succesfully", func(t *testing.T) {
		expected := []internal.TotalInvoiceCustomerCond{
			{
				Total:             1000,
				CustomerCondition: 1,
			},
			{
				Total:             2600,
				CustomerCondition: 0,
			},
		}

		topInvoices, err := invoiceRepository.GetTotalAndCustomerCondition()
		require.Nil(t, err)
		require.ElementsMatch(t, expected, topInvoices)
	})
}
