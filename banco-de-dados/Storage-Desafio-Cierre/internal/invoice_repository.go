package internal

// RepositoryInvoice is the interface that wraps the basic methods that an invoice repository should implement.
type RepositoryInvoice interface {
	// FindAll returns all invoices
	FindAll() (i []Invoice, err error)
	// Save saves an invoice
	Save(i *Invoice) (err error)
	// SaveWithId saves an invoice with a specific ID
	SaveWithId(i Invoice) (err error)
	// UpdateTotalById updates the total value by id
	UpdateTotalById(id int, total float64) (err error)
	GetTotalAndCustomerCondition() ([]TotalInvoiceCustomerCond, error)
}
