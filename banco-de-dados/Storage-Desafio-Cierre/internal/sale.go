package internal

// SaleAttributes is the struct that represents the attributes of a sale.
type SaleAttributes struct {
	// Quantity is the quantity of the sale.
	Quantity int `json:"quantity"`
	// ProductId is the product id of the sale.
	ProductId int `json:"product_id"`
	// InvoiceId is the invoice id of the sale.
	InvoiceId int `json:"invoice_id"`
}

// Sale is the struct that represents a sale.
type Sale struct {
	// Id is the unique identifier of the sale.
	Id int `json:"id"`
	// SaleAttributes is the attributes of the sale.
	SaleAttributes
}
