package repository

import (
	"database/sql"

	"app/internal"
)

// NewProductsMySQL creates new mysql repository for product entity.
func NewProductsMySQL(db *sql.DB) *ProductsMySQL {
	return &ProductsMySQL{db}
}

// ProductsMySQL is the MySQL repository implementation for product entity.
type ProductsMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all products from the database.
func (r *ProductsMySQL) FindAll() (p []internal.Product, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `description`, `price` FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var pr internal.Product
		// scan the row into the product
		err := rows.Scan(&pr.Id, &pr.Description, &pr.Price)
		if err != nil {
			return nil, err
		}
		// append the product to the slice
		p = append(p, pr)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the product into the database.
func (r *ProductsMySQL) Save(p *internal.Product) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO products (`description`, `price`) VALUES (?, ?)",
		(*p).Description, (*p).Price,
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
	(*p).Id = int(id)

	return
}

// SaveWithId saves a product with a specific ID
func (r *ProductsMySQL) SaveWithId(p internal.Product) (err error) {
	_, err = r.db.Exec(
		"INSERT INTO products (`id`, `description`, `price`) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE `id`=VALUES(`id`), `description`=VALUES(`description`), `price`=VALUES(`price`)",
		p.Id, p.Description, p.Price,
	)

	return
}

func (r *ProductsMySQL) GetTopFiveBestSellers() ([]internal.BestSellers, error) {
	var products []internal.BestSellers
	rows, err := r.db.Query("SELECT p.description, SUM(s.quantity) AS total FROM products p INNER JOIN sales s ON p.id = s.product_id GROUP BY p.id ORDER BY total DESC LIMIT 5")
	if err != nil {
		return products, err
	}

	for rows.Next() {
		var product internal.BestSellers
		err := rows.Scan(&product.Description, &product.Total)
		if err != nil {
			return products, err
		}

		products = append(products, product)
		if rows.Err() != nil {
			return []internal.BestSellers{}, err
		}
	}

	return products, nil
}
