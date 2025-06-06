package json_data

import (
	"encoding/json"
	"os"

	"github.com/bgw7/exercicio_2_criar_servidor/product"
)

type ProductData struct {
	Products []product.Product
}

func GetProducts() []product.Product {
	f, err := os.Open("products.json")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var products []product.Product

	err = json.NewDecoder(f).Decode(&products)
	if err != nil {
		panic(err)
	}
	return products
}

func GetProductById(id int64) *product.Product {
	products := GetProducts()
	for _, product := range products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}

func FilterByPrice(minPrice float64) []product.Product {
	products := GetProducts()
	var filteredProducts []product.Product

	for _, product := range products {
		if product.Price > minPrice {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}
