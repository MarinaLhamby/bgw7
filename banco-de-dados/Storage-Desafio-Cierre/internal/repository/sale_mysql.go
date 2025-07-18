package repository

import (
	"database/sql"

	"app/internal"
)

// NewSalesMySQL creates new mysql repository for sale entity.
func NewSalesMySQL(db *sql.DB) *SalesMySQL {
	return &SalesMySQL{db}
}

// SalesMySQL is the MySQL repository implementation for sale entity.
type SalesMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all sales from the database.
func (r *SalesMySQL) FindAll() (s []internal.Sale, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `quantity`, `product_id`, `invoice_id` FROM sales")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var sa internal.Sale
		// scan the row into the sale
		err := rows.Scan(&sa.Id, &sa.Quantity, &sa.ProductId, &sa.InvoiceId)
		if err != nil {
			return nil, err
		}
		// append the sale to the slice
		s = append(s, sa)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the sale into the database.
func (r *SalesMySQL) Save(s *internal.Sale) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO sales (`quantity`, `product_id`, `invoice_id`) VALUES (?, ?, ?)",
		(*s).Quantity, (*s).ProductId, (*s).InvoiceId,
	)
	if err != nil {
		return err
	}

	// get the last inserted id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set the id
	(*s).Id = int(id)

	return
}

// SaveWithId saves a sale into the database with a specific ID.
func (r *SalesMySQL) SaveWithId(s internal.Sale) (err error) {
	_, err = r.db.Exec(
		"INSERT INTO sales (`id`, `quantity`, `product_id`, `invoice_id`) VALUES (?, ?, ?, ?) ON DUPLICATE KEY UPDATE `id`=VALUES(`id`), `quantity`=VALUES(`quantity`), `product_id`=VALUES(`product_id`), `invoice_id`=VALUES(`invoice_id`)",
		s.Id, s.Quantity, s.ProductId, s.InvoiceId,
	)
	return
}

// GetTotalByInvoiceId gets the total bought by invoice id.
func (r *SalesMySQL) GetTotalByInvoiceId(invoiceId int) (float64, error) {
	var total float64
	if err := r.db.QueryRow(`SELECT SUM(p.price*s.quantity)
							FROM sales s
							INNER JOIN products p ON s.product_id = p.id
							WHERE s.invoice_id = ?`,
		invoiceId).Scan(&total); err != nil {
		return total, err
	}

	return total, nil
}
