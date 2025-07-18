package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"app/utils"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSaveWithIdCustomer(t *testing.T) {
	db := sql.OpenDB(utils.Init())

	defer db.Close()
	customerRepository := repository.NewCustomersMySQL(db)
	t.Run("Save customer with id succesfully", func(t *testing.T) {
		customer := internal.Customer{
			Id: 6,
			CustomerAttributes: internal.CustomerAttributes{
				FirstName: "Nome",
				LastName:  "Sobrenome",
				Condition: 1,
			},
		}

		err := customerRepository.SaveWithId(customer)
		require.Nil(t, err)
	})

}

func TestGetTopBuyersCustomer(t *testing.T) {
	db := sql.OpenDB(utils.Init())
	customerRepository := repository.NewCustomersMySQL(db)

	defer db.Close()

	t.Run("Save customer with id succesfully", func(t *testing.T) {
		expectedBuyers :=
			[]internal.TopBuyer{
				{
					FirstName: "Jane",
					LastName:  "Smith",
					Amount:    800,
				},
				{
					FirstName: "John",
					LastName:  "Smith",
					Amount:    700,
				},
				{
					FirstName: "Jane",
					LastName:  "Doe",
					Amount:    600,
				},
				{
					FirstName: "John",
					LastName:  "Doe",
					Amount:    500,
				},
				{
					FirstName: "Jane",
					LastName:  "Smith",
					Amount:    400,
				},
			}
		buyers, err := customerRepository.GetTopBuyers()

		require.Nil(t, err)
		require.Equal(t, expectedBuyers, buyers)
	})

}
