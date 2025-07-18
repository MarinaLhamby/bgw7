package internal

// InvoiceAttributes is the struct that represents the attributes of an invoice.
type InvoiceAttributes struct {
	// Datetime is the datetime of the invoice.
	Datetime string `json:"datetime"`
	// Total is the total of the invoice.
	Total float64 `json:"total"`
	// CustomerId is the customer id of the invoice.
	CustomerId int `json:"customer_id"`
}

// Invoice is the struct that represents an invoice.
type Invoice struct {
	// Id is the id of the invoice.
	Id int `json:"id"`
	// InvoiceAttributes is the attributes of the invoice.
	InvoiceAttributes
}

type TotalInvoiceCustomerCond struct {
	Total             float64 `json:"total"`
	CustomerCondition int     `json:"condition"`
}
