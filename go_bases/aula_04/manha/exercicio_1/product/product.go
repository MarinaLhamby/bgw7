package product

import "fmt"

type Product struct {
	ID          int64
	Name        string
	Price       float64
	Description string
	Category    string
}

type Checkout struct {
	Products []Product
}

func (c *Checkout) GetAll() {
	for _, product := range c.Products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f, Description: %s, Category: %s\n", product.ID, product.Name, product.Price, product.Description, product.Category)
	}
}

func (c *Checkout) GetByID(ID int64) Product {
	var product Product
	for _, p := range c.Products {
		if p.ID == ID {
			product = p
			return product
		}
	}

	return product
}

func (c *Checkout) Save(p Product) {
	c.Products = append(c.Products, p)
}
