package main

import (
	"exercicio_1/product"
	"fmt"
)

func main() {
	checkout := product.Checkout{
		Products: []product.Product{
			{
				Name:        "Product 1",
				ID:          1,
				Price:       20.6,
				Description: "Description product 1",
				Category:    "Category product 1",
			},
		},
	}

	checkout.Save(product.Product{
		Name:        "Product 2",
		ID:          2,
		Price:       12.6,
		Description: "Description product 2",
		Category:    "Category product 2",
	})
	checkout.GetAll()

	foundProduct := checkout.GetByID(2)
	fmt.Printf("ID: %d, Name: %s, Price: %.2f, Description: %s, Category: %s\n", foundProduct.ID, foundProduct.Name, foundProduct.Price, foundProduct.Description, foundProduct.Category)

}
