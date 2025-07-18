package repository

import (
	"database/sql"

	"app/internal"
)

// NewInvoicesMySQL creates new mysql repository for invoice entity.
func NewInvoicesMySQL(db *sql.DB) *InvoicesMySQL {
	return &InvoicesMySQL{db}
}

// InvoicesMySQL is the MySQL repository implementation for invoice entity.
type InvoicesMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all invoices from the database.
func (r *InvoicesMySQL) FindAll() (i []internal.Invoice, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `datetime`, `total`, `customer_id` FROM invoices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var iv internal.Invoice
		// scan the row into the invoice
		err := rows.Scan(&iv.Id, &iv.Datetime, &iv.Total, &iv.CustomerId)
		if err != nil {
			return nil, err
		}
		// append the invoice to the slice
		i = append(i, iv)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the invoice into the database.
func (r *InvoicesMySQL) Save(i *internal.Invoice) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO invoices (`datetime`, `total`, `customer_id`) VALUES (?, ?, ?)",
		(*i).Datetime, (*i).Total, (*i).CustomerId,
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
	(*i).Id = int(id)

	return
}

// SaveWithId saves the invoice into the database with a specific ID.
func (r *InvoicesMySQL) SaveWithId(i internal.Invoice) (err error) {
	_, err = r.db.Exec("INSERT INTO invoices (`id`, `datetime`, `total`, `customer_id`) VALUES (?, ?, ?, ?) ON DUPLICATE KEY UPDATE `id`=VALUES(`id`), `datetime`=VALUES(`datetime`), `total`=VALUES(`total`), `customer_id`=VALUES(`customer_id`)", i.Id, i.Datetime, i.Total, i.CustomerId)
	return
}

// UpdateTotalById updates the total value by id
func (r *InvoicesMySQL) UpdateTotalById(id int, total float64) (err error) {
	_, err = r.db.Exec("UPDATE invoices SET total = ? WHERE id = ?", total, id)
	return
}

func (r *InvoicesMySQL) GetTotalAndCustomerCondition() ([]internal.TotalInvoiceCustomerCond, error) {
	var invoices []internal.TotalInvoiceCustomerCond

	rows, err := r.db.Query("SELECT TRUNCATE(SUM(i.total),2), c.condition FROM invoices i INNER JOIN customers c ON i.customer_id = c.id GROUP BY c.condition")
	if err != nil {
		return invoices, err
	}

	for rows.Next() {
		var invoice internal.TotalInvoiceCustomerCond
		err := rows.Scan(&invoice.Total, &invoice.CustomerCondition)
		if err != nil {
			return invoices, err
		}

		invoices = append(invoices, invoice)

		if rows.Err() != nil {
			return []internal.TotalInvoiceCustomerCond{}, err
		}
	}

	return invoices, nil
}
