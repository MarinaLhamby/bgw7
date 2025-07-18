package service

import (
	"app/internal"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const (
	customersFile = "customers.json"
	invoicesFile  = "invoices.json"
	productsFile  = "products.json"
	salesFile     = "sales.json"
)

type dataLoadDefault struct {
	customerRepository internal.RepositoryCustomer
	invoiceRepository  internal.RepositoryInvoice
	productRepository  internal.RepositoryProduct
	saleRepository     internal.RepositorySale
}

func NewDataLoadDefault(cr internal.RepositoryCustomer,
	ir internal.RepositoryInvoice,
	pr internal.RepositoryProduct,
	sr internal.RepositorySale) *dataLoadDefault {
	return &dataLoadDefault{
		customerRepository: cr,
		invoiceRepository:  ir,
		productRepository:  pr,
		saleRepository:     sr,
	}
}

func (d *dataLoadDefault) Load(path string) error {
	customers, err := loadJSON[internal.Customer](path, customersFile)
	if err != nil {
		return err
	}

	for _, c := range customers {
		err := d.customerRepository.SaveWithId(c)
		if err != nil {
			continue
		}
	}

	err = d.loadProducts(path)
	if err != nil {
		return err
	}

	invoices, err := d.loadInvoices(path)
	if err != nil {
		return err
	}

	err = d.loadSales(path)
	if err != nil {
		return err
	}

	err = d.updateInvoicesAfterLoadingSales(invoices)
	if err != nil {
		return err
	}
	return nil
}

func (d *dataLoadDefault) loadProducts(path string) error {
	products, err := loadJSON[internal.Product](path, productsFile)
	if err != nil {
		return err
	}

	for _, p := range products {
		err := d.productRepository.SaveWithId(p)
		if err != nil {
			continue
		}
	}
	return nil
}

func (d *dataLoadDefault) loadSales(path string) error {
	sales, err := loadJSON[internal.Sale](path, salesFile)
	if err != nil {
		return err
	}

	for _, s := range sales {
		err := d.saleRepository.SaveWithId(s)
		if err != nil {
			continue
		}
	}

	return nil
}

func (d *dataLoadDefault) loadInvoices(path string) ([]internal.Invoice, error) {
	invoices, err := loadJSON[internal.Invoice](path, invoicesFile)
	if err != nil {
		return []internal.Invoice{}, err
	}

	for _, i := range invoices {
		err = d.invoiceRepository.SaveWithId(i)
		if err != nil {
			continue
		}
	}

	return invoices, nil
}

func (d *dataLoadDefault) updateInvoicesAfterLoadingSales(invoices []internal.Invoice) error {
	for _, i := range invoices {
		total, err := d.saleRepository.GetTotalByInvoiceId(i.Id)
		if err != nil {
			continue
		}

		err = d.invoiceRepository.UpdateTotalById(i.Id, total)
		if err != nil {
			continue
		}
	}

	return nil
}

func loadJSON[T any](path, filename string) ([]T, error) {
	var result []T

	jsonFile, err := os.Open(fmt.Sprintf("%s/%s", path, filename))
	if err != nil {
		return result, fmt.Errorf("[%w] error opening %s file", err, filename)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return result, fmt.Errorf("[%w] error reading %s file", err, filename)
	}

	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return result, fmt.Errorf("[%w] error unmarshaling %s file", err, filename)
	}

	return result, err
}
