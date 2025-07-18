package internal

// RepositorySale is the interface that wraps the basic Sale methods.
type RepositorySale interface {
	// FindAll returns all sales.
	FindAll() (s []Sale, err error)
	// Save saves a sale.
	Save(s *Sale) (err error)
	//SaveWithId saves a sale into the database with a specific ID.
	SaveWithId(s Sale) (err error)
	//GetTotalByInvoiceId gets the total bought by invoice id.
	GetTotalByInvoiceId(invoiceId int) (float64, error)
}
