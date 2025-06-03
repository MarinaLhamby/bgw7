package main

import (
	"exercicio_2/product"
	"fmt"
)

func main() {
	smallProduct := factoryProduct(product.Small, 10)
	mediumProduct := factoryProduct(product.Medium, 10)
	largeProduct := factoryProduct(product.Large, 10)

	fmt.Printf("Small product price: %.2f\n", smallProduct.Price())
	fmt.Printf("Medium product price: %.2f\n", mediumProduct.Price())
	fmt.Printf("Large product price: %.2f\n", largeProduct.Price())

}

func factoryProduct(productType product.ProductType, value float64) product.Product {
	return product.NewItem(productType, value)
}
