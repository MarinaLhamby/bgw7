package main

import (
	"fmt"

	"github.com/bgw7/exercicio_1_iniciar_projeto/json_data"
)

func main() {
	products := json_data.GetJsonDataFromFile()
	for _, product := range products {
		fmt.Println("ID:", product.ID)
		fmt.Println("Name:", product.Name)
		fmt.Println("Quantity:", product.Quantity)
		fmt.Println("Code Value:", product.CodeValue)
		fmt.Println("Is Published:", product.IsPublished)
		fmt.Println("Expiration:", product.Expiration)
		fmt.Println("Price:", product.Price)
		fmt.Println("-------------------------")
	}
}
