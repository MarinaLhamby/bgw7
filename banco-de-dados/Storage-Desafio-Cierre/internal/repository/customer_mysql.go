package repository

import (
	"database/sql"

	"app/internal"
)

// NewCustomersMySQL creates new mysql repository for customer entity.
func NewCustomersMySQL(db *sql.DB) *CustomersMySQL {
	return &CustomersMySQL{db}
}

// CustomersMySQL is the MySQL repository implementation for customer entity.
type CustomersMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all customers from the database.
func (r *CustomersMySQL) FindAll() (c []internal.Customer, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `first_name`, `last_name`, `condition` FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var cs internal.Customer
		// scan the row into the customer
		err := rows.Scan(&cs.Id, &cs.FirstName, &cs.LastName, &cs.Condition)
		if err != nil {
			return nil, err
		}
		// append the customer to the slice
		c = append(c, cs)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the customer into the database.
func (r *CustomersMySQL) Save(c *internal.Customer) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO customers (`first_name`, `last_name`, `condition`) VALUES (?, ?, ?)",
		(*c).FirstName, (*c).LastName, (*c).Condition,
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
	(*c).Id = int(id)

	return
}

func (r *CustomersMySQL) SaveWithId(c internal.Customer) (err error) {
	_, err = r.db.Exec(
		"INSERT INTO customers (`id`, `first_name`, `last_name`, `condition`) VALUES (?, ?, ?, ?) ON DUPLICATE KEY UPDATE `first_name`=VALUES(`first_name`), `last_name`=VALUES(`last_name`), `condition`=VALUES(`condition`)",
		c.Id, c.FirstName, c.LastName, c.Condition,
	)
	return err
}

func (r *CustomersMySQL) GetTopBuyers() ([]internal.TopBuyer, error) {
	var topBuyers []internal.TopBuyer
	rows, err := r.db.Query("SELECT c.first_name, c.last_name, TRUNCATE(SUM(i.total),2) AS amount FROM customers c INNER JOIN invoices i ON c.id = i.customer_id GROUP BY c.id ORDER BY amount DESC LIMIT 5")
	if err != nil {
		return topBuyers, err
	}

	for rows.Next() {
		var buyer internal.TopBuyer
		err := rows.Scan(&buyer.FirstName, &buyer.LastName, &buyer.Amount)
		if err != nil {
			return topBuyers, err
		}

		topBuyers = append(topBuyers, buyer)
		if rows.Err() != nil {
			return []internal.TopBuyer{}, err
		}
	}

	return topBuyers, nil
}
